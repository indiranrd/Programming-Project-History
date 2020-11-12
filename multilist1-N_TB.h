#ifndef MULTILIST1-N_TB_H_INCLUDED
#define MULTILIST1-N_TB_H_INCLUDED
#include "Parent_TB.h"
#include "Child_TB.h"

void connect(address_P &P, address_C &C);
void disconnect(address_C &C);
void printRelasi(List_P L, List_C L2);
/*void Tambah_Data_Barang_Jasa(List_P L, List_C L2, string a);
*/void Relasi_Produk(List_P L, List_C L2, string a);
void Hapus_Perusahaan(List_P &L, List_C &L2, string a);
void Hapus_dan_Ganti_produk(List_P &L, List_C &L2, Deskripsi z, Great_Data s);
/*void Tambah_Data_Barang_Jasa2(List_P L, List_C L2, string a);
void Pindah_Data(List_P L, List_C &L2, Deskripsi x);*/

#endif // MULTILIST1-N_TB_H_INCLUDED
