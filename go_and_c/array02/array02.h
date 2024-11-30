#include <stdlib.h>
#include <string.h>

#define MAX_NAME_LENGTH 256
#define MAX_PEOPLE 128000

struct person {
    char firstName[MAX_NAME_LENGTH];
    char lastName[MAX_NAME_LENGTH];
    int age;
} typedef Person;

struct people {
    int length;
    Person* entries;
} typedef People;

Person CreatePerson(const char *firstName, const char *lastName, int age);
People* GetPeople(Person people[], int len, int max);
void FreePeople(People* p);