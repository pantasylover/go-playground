#include <stdlib.h>
#include <string.h>

#define MAX_NAME_LENGTH 256
#define MAX_PEOPLE 10

struct person {
    char firstName[MAX_NAME_LENGTH];
    char lastName[MAX_NAME_LENGTH];
    int age;
} typedef Person;

struct people {
    int length;
    Person* entries[MAX_PEOPLE];
} typedef People;

Person CreatePerson(const char *firstName, const char *lastName, int age);
People* GetPeople(Person people[MAX_PEOPLE], int len);
void FreePeople(People* p);