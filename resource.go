package main

type Resource struct {
	altResource *uintptr
}

func (r *Resource) MakeClient(_, altInfo uintptr) uintptr {
	creationInfo := GetCreationInfo(altInfo)
	creationInfo.SetType("js")

	return TooPtr(true)
}
