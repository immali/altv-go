package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/immali/altv-go/events"
)

import "C"

// ResourceImpl in go
type ResourceImpl struct {
	altImpl uintptr
	module  *Module
}

func (r *ResourceImpl) makeClient(_, altInfo uintptr) uintptr {
	creationInfo := GetCreationInfo(altInfo)
	creationInfo.SetType("js")

	return BoolPtr(false)
}

func (r *ResourceImpl) start() uintptr {
	return BoolPtr(true)
}

func (r *ResourceImpl) stop() uintptr {
	return BoolPtr(true)
}

func (r *ResourceImpl) onEvent(res, altEvent uintptr) uintptr {
	proc := r.module.dll.MustFindProc("alt_IResource_GetName")
	namePtr, _, _ := proc.Call(res)

	proc = r.module.dll.MustFindProc("alt_StringView_GetData")
	namePtr, _, _ = proc.Call(namePtr)

	cname := (*C.char)(unsafe.Pointer(namePtr))
	name := C.GoString(cname)

	eventID := events.GetEventID(r.module.dll, altEvent)

	_module.logInfo(fmt.Sprintf("Event (%d) for resource %s", eventID, name))

	// switch eventID {
	// case events.PLAYER_CONNECT:
	// 	{
	// 		fmt.Printf("Player connected")
	// 		break
	// 	}
	// }
	return BoolPtr(true)
}

func (r *ResourceImpl) onCreateBaseObjectCallback() uintptr {
	return BoolPtr(true)
}

func (r *ResourceImpl) onRemoveBaseObjectCallback() uintptr {
	return BoolPtr(true)
}

func (r *ResourceImpl) createImpl(_, resource uintptr) uintptr {
	proc := r.module.dll.MustFindProc("alt_CAPIResource_Impl_Create")

	MakeClient := syscall.NewCallback(r.makeClient)
	Start := syscall.NewCallback(r.start)
	Stop := syscall.NewCallback(r.stop)
	OnEvent := syscall.NewCallback(r.onEvent)
	OnCreateBaseObject := syscall.NewCallback(r.onCreateBaseObjectCallback)
	OnRemoveBaseObject := syscall.NewCallback(r.onRemoveBaseObjectCallback)

	ret, _, _ := proc.Call(resource, MakeClient, Start, Stop, OnEvent, OnCreateBaseObject, OnRemoveBaseObject)

	r.altImpl = ret
	return ret
}

// NewResourceImpl creates a new ResourceImpl
func NewResourceImpl(r *Runtime) func(altRuntime, altResource uintptr) uintptr {
	impl := &ResourceImpl{}
	impl.module = r.module

	r.resourceImpl = append(r.resourceImpl, impl)

	return impl.createImpl
}
