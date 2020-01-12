package main

import "C"

var _module *Module

// GetSDKVersion returns the SDK version
//export GetSDKVersion
func GetSDKVersion() int {
	initModule()

	return _module.getSDKVersion()
}

//export altMain
func altMain(altCore uintptr) bool {
	initModule()

	_module.altCore = altCore

	_module.logInfo("=== Module Initialized ===")
	_module.createRuntime()

	return true
}

func initModule() {
	if _module == nil {
		_module = NewModule()
	}
}

func main() {}
