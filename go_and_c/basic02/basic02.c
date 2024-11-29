#include "basic02.h"

Person CreatePerson(const char *firstName, const char *lastName, int age) {
    Person ret = {0};
    ret.firstName = (char *)firstName;
    ret.lastName = (char *)lastName;
    ret.age = age;

    return ret;
}

void ChangeAge(Person *p, int age) {
    p->age = age;
}