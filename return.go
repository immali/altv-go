package main

import "unsafe"

// TooPtr returns pointer to given value
func TooPtr(val interface{}) uintptr {
	return (uintptr)(unsafe.Pointer(&val))
}

// StrPtr returns pointer to given string
func StrPtr(val string) uintptr {
	return (uintptr)(unsafe.Pointer(&val))
}

// BoolPtr returns pointer to given bool
func BoolPtr(val bool) uintptr {
	return (uintptr)(unsafe.Pointer(&val))
}
