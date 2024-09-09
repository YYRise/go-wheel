#include <string.h>
#include <stdio.h>
#include <stdlib.h>

char* Add(char* src, int n)
{
    char str[20];
    sprintf(str, "%d", n);
    char *result = malloc(strlen(src)+strlen(str)+1);
    strcpy(result, src);
    strcat(result, str);
    return result;
}
