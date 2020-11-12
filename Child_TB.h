#ifndef CHILD_TB_H_INCLUDED
#define CHILD_TB_H_INCLUDED
#include <iostream>
#include "Parent_TB.h"

using namespace std;
#define first(L) L.first
#define next(P) P->next
#define info(P) P->info
#define parent(C) C->parent


struct Great_Data{
    string Barang;
    string nama_Per;
};
typedef struct elm_C *address_C;
struct elm_C{
    Great_Data info;
    address_C next;
    address_P parent;

};
struct List_C{
    address_C first;
};

void createList(List_C &L);
address_C alokasi(Great_Data y);
void insertFirstChild(List_C &L2, address_C P);
void insertLastChild(List_C &L2, address_C C);
void deleteFirstChild(List_C &L2, address_C &C);
void deleteLastChild(List_C &L2, address_C &C);
address_C FindElementChild(List_C L2, string b);





#endif // CHILD_TB_H_INCLUDED
