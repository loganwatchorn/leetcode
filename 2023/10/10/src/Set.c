#include <stdlib.h>
#include "Set.h"

#define TABLE_SIZE 10000

typedef struct Node {
    int value;
    struct Node* next;
} Node;

struct Set {
    int size;
    int tableSize;
    Node* table[];
};

static int chooseTableSize(int n) {
    // https://planetmath.org/goodhashtableprimes
    if (n >= (1 << 5) && n < (1 << 6)) return 53;
    if (n >= (1 << 6) && n < (1 << 7)) return 97;
    if (n >= (1 << 7) && n < (1 << 8)) return 193;
    if (n >= (1 << 8) && n < (1 << 9)) return 389;
    if (n >= (1 << 9) && n < (1 << 10)) return 769;
    if (n >= (1 << 10) && n < (1 << 11)) return 1543;
    if (n >= (1 << 11) && n < (1 << 12)) return 3079;
    if (n >= (1 << 12) && n < (1 << 13)) return 6151;
    if (n >= (1 << 13) && n < (1 << 14)) return 12289;
    if (n >= (1 << 14) && n < (1 << 15)) return 24593;
    if (n >= (1 << 15) && n < (1 << 16)) return 49157;
    if (n >= (1 << 16) && n < (1 << 17)) return 98317;
    return 196613;
}

static unsigned int hash(int value, int tableSize) {
    return value % tableSize;
}

Set* createSet(int n) {
    int tableSize = chooseTableSize(n);

    Set* set = (Set*) malloc(2*sizeof(int) + tableSize*sizeof(Node*));
    set->size = 0;
    set->tableSize = tableSize;
    memset(set->table, 0, tableSize*sizeof(Node*));
    
    return set;
}

void destroySet(Set* set) {
    if (set) {
        for (int i=0 ; i<set->tableSize ; i++) {
            Node* curr = set->table[i];
            while (curr) {
                Node* temp = curr;
                curr = curr->next;
                free(temp);
            }
        }
        free(set);
    }
}

void addToSet(Set* set, int value) {
    unsigned int key = hash(value, set->tableSize);
    Node* curr = set->table[key];

    while (curr) {
        if (curr->value == value) { return; }
        if (!curr->next) { break; }
        curr = curr->next;
    }

    Node* node = (Node*) malloc(sizeof(Node));
    node->value = value;
    node->next = NULL;
    
    if (curr) { curr->next = node; }
    else { set->table[key] = node; }
    set->size++;
}

Set* arrToSet(int* arr, int arrSize) {
    Set* set = createSet(arrSize);
    for (int i=0 ; i<arrSize ; i++) {
        addToSet(set, arr[i]);
    }
    return set;
}

int* setToArr(Set* set, int* arrSize) {
    *arrSize = set->size;

    int* arr = (int*) malloc(set->size * sizeof(int));

    int i=0;
    Node* curr;

    for (int key=0 ; key<set->tableSize ; key++) {
        curr = set->table[key];
        while (curr) {
            arr[i] = curr->value;
            i++;
            curr = curr->next;
        }
    }

    return arr;
}

