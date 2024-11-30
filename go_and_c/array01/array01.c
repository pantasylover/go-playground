#include "array01.h"

Person CreatePerson(const char *firstName, const char *lastName, int age) {
    Person ret = {0};

    ret.age = age;

    strncpy(ret.firstName, firstName, MAX_NAME_LENGTH);
    strncpy(ret.lastName, lastName, MAX_NAME_LENGTH);

    return ret;
}

People* GetPeople(Person people[], int len) {
    len = (len > MAX_PEOPLE) ? MAX_PEOPLE : len;

    People *ret = malloc(sizeof(People));

    if (!ret)
        return NULL;

    memset(ret, 0, sizeof(People));

    int rLen = 0;

    for (int i = 0; i < len; i++) {        

        ret->entries[i] = malloc(sizeof(Person));

        if (!ret->entries[i])
            continue;

        rLen++;

        memcpy(ret->entries[i], &people[i], sizeof(Person));
    }

    ret->length = rLen;

    return ret;
}

void FreePeople(People* p) {
    if (p) {
        for (int i = 0; i < p->length; i++) {
            Person *t = p->entries[i];

            if (t)
                free(t);
        }

        free(p);
    }
}