typedef struct alt_IResource_CreationInfo {
    char* type;
    char* name;
    char* main;
} alt_IResource_CreationInfo;

void alt_IResource_CreationInfo_SetType(alt_IResource_CreationInfo* instance, char *tp);
char* alt_IResource_CreationInfo_GetType(alt_IResource_CreationInfo* instance);