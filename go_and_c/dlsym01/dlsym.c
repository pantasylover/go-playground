#include <stdio.h>
#include <string.h>

int Func1() {
    return 1;
}

void Func2(int x, int y) {
    printf("x => %d. y => %d\n", x, y);
}

int Func3(const char *x, int y, int z) {
    printf("x => %s. y => %d. z => %d\n", x, y, z);

    if (strlen(x) > y)
        return 1;

    return 0;
}