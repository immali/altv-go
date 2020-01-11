package main

import (
	"fmt"
	"syscall"
)

type Runtime struct {
	runtime      uintptr
	resourceImpl *ResourceImpl
	module       *Module
}

func (r *Runtime) createImpl() uintptr {
	r.resourceImpl = NewResourceImpl(r.module)

	return TooPtr(true)
}

func (r *Runtime) onTick() uintptr {
	return 0
}

func (r *Runtime) destroyImpl() uintptr {
	return 0
}

func (r *Runtime) register() bool {
	proc := r.module.dll.MustFindProc("alt_ICore_RegisterScriptRuntime")
	resPtr := StrPtr("go")
	ret, _, _ := proc.Call(r.module.core, resPtr, r.runtime)

	if i := int(ret); i == 1 {
		fmt.Println("Registered")
		return true
	}

	fmt.Println("Not registered")
	return false
}

func NewRuntime(m *Module) *Runtime {
	r := &Runtime{module: m}

	r.resourceImpl = NewResourceImpl(m)

	CreateImpl := syscall.NewCallback(r.resourceImpl.createImpl)
	DestroyImpl := syscall.NewCallback(r.destroyImpl)
	OnTick := syscall.NewCallback(r.onTick)

	proc := m.dll.MustFindProc("alt_CAPIScriptRuntime_Create")
	ret, _, _ := proc.Call(CreateImpl, DestroyImpl, OnTick)

	r.runtime = ret
	r.register()

	return r
}
