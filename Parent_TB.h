#ifndef PARENT_TB_H_INCLUDED
#define PARENT_TB_H_INCLUDED
#include <iostream>
#include <time.h>
#include <stdio.h>
#include <stdlib.h>

using namespace std;
#define first(L) L.first
#define last(L) L.last
#define prev(P) P->prev
#define next(P) P->next
#define info(P) P->info

struct Deskripsi{
    string Nama;
    string Pendiri;
    string Deskription;
    string Tahun;
    float Laba_Tahun2019;
    float Laba_Tahun2020;
};
typedef struct elm_P *address_P;
struct elm_P{
    Deskripsi info;
    address_P next;
    address_P prev;
    address_P parent2;

};
struct List_P{
    address_P first;
    address_P last;
};
struct Cuan{
    string perusahaan;
    float cuan_lalu;
    float cuan_sekarang;
};

void createList(List_P &L);
address_P alokasi(Deskripsi x);
void insertFirst(List_P &L, address_P P);
void insertLast(List_P &L, address_P P);
void deleteFirst(List_P &L, address_P &P);
void deleteLast(List_P &L, address_P &P);
void printInfo(List_P L);
address_P findElement(List_P L, string a);
void Perusahaan_Profit(List_P L, Deskripsi x);





#endif // PARENT_TB_H_INCLUDED
