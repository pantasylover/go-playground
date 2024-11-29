#include <stdlib.h>

struct person {
    char *firstName;
    char *lastName;
    int age;
} typedef Person;

Person CreatePerson(const char *firstName, const char *lastName, int age);
void ChangeAge(Person *p, int age);