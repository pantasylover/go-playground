#include "pointer01.h"

Test* CreateTest(__u8 field1, __u16 field2, __u32 field3, __u64 field4, const char *field5) {
    // Manually allocate test structure.
    Test* ret = malloc(sizeof(Test));

    // Set memory allocated to 0's.
    memset(ret, 0, sizeof(Test));

    // Set fields.
    ret->field1 = field1;
    ret->field2 = field2;
    ret->field3 = field3;
    ret->field4 = field4;

    strcpy(ret->field5, field5);

    return ret;
}

__u8 GetField1(Test* t) {
    return t->field1;
}

__u16 GetField2(Test* t) {
    return t->field2;
}

__u32 GetField3(Test* t) {
    return t->field3;
}

__u64 GetField4(Test* t) {
    return t->field4;
}

char* GetField5(Test* t) {
    return t->field5;
}

void SetField1(Test* t, __u8 val) {
    t->field1 = val;
}

void SetField2(Test* t, __u16 val) {
    t->field2 = val;
}

void SetField3(Test* t, __u32 val) {
    t->field3 = val;
}

void SetField4(Test* t, __u64 val) {
    t->field4 = val;
}

void SetField5(Test* t, const char *val) {
    strcpy(t->field5, val);
}
