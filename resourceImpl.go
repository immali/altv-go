package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/immali/altv-go/events"
)

import "C"

type ResourceImpl struct {
	impl   uintptr
	module *Module
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

	fmt.Printf("Event (%d) for resource %s \n", eventID, name)

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
	fmt.Println("Create Resource Impl")
	proc := r.module.dll.MustFindProc("alt_CAPIResource_Impl_Create")

	// MakeClientCallback := syscall.NewCallback(func(info uintptr) uintptr {
	// 	crInfo := GetCreationInfo(info)
	// 	crInfo.SetType("js")

	// 	return BoolPtr(true)
	// })

	// StartCallback := syscall.NewCallback(func() uintptr {
	// 	fmt.Println("Start Resource")

	// 	return 1
	// })

	// StopCallback := syscall.NewCallback(func() uintptr {
	// 	fmt.Println("Stop Resource")
	// 	return 0
	// })

	// OnEventCallback := syscall.NewCallback(func(_, altEvent uintptr) uintptr {
	// 	proc := r.module.dll.MustFindProc("alt_CEvent_GetType")
	// 	ret, _, _ := proc.Call(altEvent)

	// 	switch int(ret) {
	// 	case events.NONE:
	// 		{
	// 			break
	// 		}

	// 	case events.PLAYER_CONNECT:
	// 		{
	// 			break
	// 		}
	// 	}

	// 	fmt.Printf("Event with ID: %d happened\n", int(ret))

	// 	return 0
	// })

	// OnCreateBaseObjectCallback := syscall.NewCallback(func() uintptr {
	// 	return 0
	// })

	// OnDeleteBaseObjectCallback := syscall.NewCallback(func() uintptr {
	// 	return 0
	// })
	MakeClient := syscall.NewCallback(r.makeClient)
	Start := syscall.NewCallback(r.start)
	Stop := syscall.NewCallback(r.stop)
	OnEvent := syscall.NewCallback(r.onEvent)
	OnCreateBaseObject := syscall.NewCallback(r.onCreateBaseObjectCallback)
	OnRemoveBaseObject := syscall.NewCallback(r.onRemoveBaseObjectCallback)

	ret, _, _ := proc.Call(resource, MakeClient, Start, Stop, OnEvent, OnCreateBaseObject, OnRemoveBaseObject)

	return ret
}

func NewResourceImpl(m *Module) *ResourceImpl {
	r := &ResourceImpl{}
	r.module = m

	return r
}
