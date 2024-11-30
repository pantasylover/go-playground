#include "array02.h"

Person CreatePerson(const char *firstName, const char *lastName, int age) {
    Person ret = {0};

    ret.age = age;

    strncpy(ret.firstName, firstName, MAX_NAME_LENGTH);
    strncpy(ret.lastName, lastName, MAX_NAME_LENGTH);

    return ret;
}

People* GetPeople(Person people[], int len, int max) {
    len = (len > MAX_PEOPLE) ? MAX_PEOPLE : len;

    People *ret = malloc(sizeof(People));

    if (!ret)
        return NULL;

    memset(ret, 0, sizeof(People));

    ret->entries = malloc(sizeof(Person) * max);

    if (!ret->entries) {
        free(ret);

        return NULL;
    }

    memset(ret->entries, 0, sizeof(Person) * max);

    int rLen = 0;

    for (int i = 0; i < len; i++) {        

        Person *p = &ret->entries[i];

        rLen++;

        memcpy(p, &people[i], sizeof(Person));
    }

    ret->length = rLen;

    return ret;
}