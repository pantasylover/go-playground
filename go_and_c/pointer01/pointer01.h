#include <stdlib.h>
#include <string.h>
#include <linux/types.h>

#define MAX_FIELD5_LEN 64

struct test {
    __u8 field1;
    __u16 field2;
    __u32 field3;
    __u64 field4;
    char field5[MAX_FIELD5_LEN];
} typedef Test;

Test* CreateTest(__u8 field1, __u16 field2, __u32 field3, __u64 field4, const char *field5);

__u8 GetField1(Test* t);
__u16 GetField2(Test* t);
__u32 GetField3(Test* t);
__u64 GetField4(Test* t);
char* GetField5(Test* t);

void SetField1(Test* t, __u8 val);
void SetField2(Test* t, __u16 val);
void SetField3(Test* t, __u32 val);
void SetField4(Test* t, __u64 val);
void SetField5(Test* t, const char *val);