#include <stddef.h>
#include <stdlib.h>
#include "minOperations.h"

int main() {
    int res;

    printf("test1: ");
    int test1[] = {4,2,5,3};
    res = minOperations(test1, 4);
    if (res == 0) { printf("pass"); } else { printf("fail"); }
    printf("\n");


    printf("test2: ");
    int test2[] = {1,2,3,5,6};
    res = minOperations(test2, 5);
    if (res == 1) { printf("pass"); } else { printf("fail"); }
    printf("\n");


    printf("test3: ");
    int test3[] = {1,10,100,1000};
    res = minOperations(test3, 4);
    if (res == 3) { printf("pass"); } else { printf("fail"); }
    printf("\n");

    return 0;
}
