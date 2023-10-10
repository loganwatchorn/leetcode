#ifndef SET_H
#define SET_H

typedef struct Set Set;

Set* createSet(int n);
void destroySet(Set* set);
void addToSet(Set* set, int value);
Set* arrToSet(int* arr, int arrSize);
int* setToArr(Set* set, int* arrSize);

#endif