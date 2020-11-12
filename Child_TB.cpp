#include "Child_TB.h"
#include <iostream>

void createList(List_C &L){
    first(L) = NULL;
}
address_C alokasi(Great_Data y){
    address_C P = new elm_C;
    info(P).Barang = y.Barang;
    next(P) = NULL;
    parent(P) = NULL;
    return P;
}
void insertFirstChild(List_C &L, address_C P){
    if(first(L) == NULL){
        first(L) = P;
    }else {
        next(P) = first(L);
        first(L) = P;
    }
}
void insertLastChild(List_C &L, address_C C){
    address_C Q = first(L);
    while(next(Q) != NULL){
        Q = next(Q);
    }
    next(Q) = C;

}
void deleteLastChild(List_C &L2, address_C &C){
    address_C Q = first(L2);
    while(next(next(Q)) != NULL){
        Q = next(Q);
    }
    C = next(Q);
    next(Q) = NULL;
}
void deleteFirstChild(List_C &L2, address_C &C){
    if(first(L2) == NULL){
        cout<<"empty"<<endl;
    }else{
        C = first(L2);
        first(L2) = next(C);
        next(C) = NULL;
    }
}
address_C FindElementChild(List_C L2, string b){
    address_C C = first(L2);
    while(C != NULL && info(C).Barang != b){
        C = next(C);
    }
    if(info(C).Barang == b){
        return C;
    }else {
        return NULL;
    }
}


