#include <stdlib.h>

int compare(const void * a, const void * b) {
    return (*(int*)a - *(int*)b);
}

void sort(int* arr, int arrSize) {
    qsort(arr, arrSize, sizeof(int), compare);
}