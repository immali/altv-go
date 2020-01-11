package events

import (
	"syscall"
)

var (
	NONE              = 0
	PLAYER_CONNECT    = 1
	PLAYER_DISCONNECT = 3
	RESOURCE_START    = 4
	RESOURCE_STOP     = 5
)

func GetEventID(dll *syscall.DLL, altEvent uintptr) int {
	proc := dll.MustFindProc("alt_CEvent_GetType")
	ret, _, _ := proc.Call(altEvent)

	return int(ret)
}
