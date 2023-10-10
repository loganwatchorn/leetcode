#include <stdio.h>
#include <stdlib.h>
#include "Set.h"
#include "sort.h"

int minOperations(int* nums, int numsSize) {
    Set* set = arrToSet(nums, numsSize);
    int uniqueSize;
    int* unique = setToArr(set, &uniqueSize);
    destroySet(set);

    sort(unique, uniqueSize);

    int l = -1;
    int r = 0;
    int maxCount = 0;

    int maxDiff = numsSize-1;

    while (r<uniqueSize) {
        l++;
        while (r<uniqueSize && unique[r]-unique[l]<=maxDiff) {
            if (r-l >= maxCount) { maxCount++; }
            r++;
        }
    }

    free(unique);
    return numsSize - maxCount;
}