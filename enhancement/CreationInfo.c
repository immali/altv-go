#include "CreationInfo.h"

void alt_IResource_CreationInfo_SetType(alt_IResource_CreationInfo* instance, char *tp) {
	instance->type = tp;
}

char* alt_IResource_CreationInfo_GetType(alt_IResource_CreationInfo* instance) {
	return instance->type;
}