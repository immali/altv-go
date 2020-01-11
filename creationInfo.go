package main

import "unsafe"

// #include "enhancement/CreationInfo.c"
import "C"

// CreationInfo alt_IResource_CreationInfo mirror
type CreationInfo struct {
	altInfoPtr *C.alt_IResource_CreationInfo
	Type       string
	Name       string
	Main       string
}

// SetType sets the type of the CreationInfo
func (i *CreationInfo) SetType(str string) {
	C.alt_IResource_CreationInfo_SetType(i.altInfoPtr, C.CString(str))
	i.Type = str
}

// GetCreationInfo fetches information for the info
func GetCreationInfo(ptr uintptr) *CreationInfo {
	info := &CreationInfo{}
	altInfoPtr := (*C.alt_IResource_CreationInfo)(unsafe.Pointer(&ptr))

	infoType := C.alt_IResource_CreationInfo_GetType(altInfoPtr)

	info.altInfoPtr = altInfoPtr
	info.Type = C.GoString(infoType)

	return info
}
