#include "multilist1-N_TB.h"
#include <iostream>


void connect(address_P &P, address_C &C){
    if(P != NULL && C != NULL){
        parent(C) = P;
    }
}
void disconnect(address_C &C){
    if(C !=NULL){
        parent(C) = NULL;
    }
}
void printRelasi(List_P L, List_C L2){
    address_P P;
    address_C C;
    P = first(L);
    while(next(P) != first(L)){
        cout<<endl;
        cout<<"Nama Perusahaan   : "<<" ";
        cout<<info(P).Nama   <<" ";
        cout<<endl;
        cout<<"Laba Perusahaan   : "<<" ";
        cout<<info(P).Laba_Tahun2020<<" ";
        C = first(L2);
        while(C != NULL){
            if(parent(C) == P){
                cout<<endl;
                cout<<"Daftar Barang :"<<endl;
                cout<<info(C).Barang<<endl;
            }
            C = next(C);
        }
        P = next(P);
        cout<<endl;
    }
    cout<<endl;
    cout<<"Nama Perusahaan   : "<<" ";
    cout<<info(P).Nama   <<" ";
    cout<<endl;
    cout<<"Laba Perusahaan   : "<<" ";
    cout<<info(P).Laba_Tahun2020<<" ";
    cout<<endl;
    C = first(L2);
    while(C != NULL){
        if(parent(C) == P){
          cout<<"Daftar Barang :"<<endl;
          cout<<info(C).Barang<<endl;
        }
        C = next(C);
    }
}

void Relasi_Produk(List_P L, List_C L2, string a){
    address_P P;
    address_C C;
    int i;
    cout<<"Barang : "<<"["<<a<<"]"<<endl;
    cout<<"Terdapat diperusahaan :"<<endl;
    P = first(L);
    while(next(P) != first(L)){
        C = first(L2);
        while(C != NULL){
            if(parent(C) == P && info(C).Barang == a){
                cout<<"["<<info(P).Nama<<"]"<<endl;
            }
            C = next(C);
        }
        P = next(P);
    }
    C = first(L2);
        while(C != NULL){
            if(parent(C) == P && info(C).Barang == a){
                cout<<"["<<info(P).Nama<<"]"<<endl;
            }
            C = next(C);
        }
}
void Hapus_Perusahaan(List_P &L,List_C &L2, string a){
    address_P P;
    address_P Q;
    address_C C;
    P = findElement(L,a);
    if(P != NULL){
        if(P == first(L)){
            C = first(L2);
            while(C != NULL){
                if(parent(C) == P){
                    disconnect(C);
                }
                C = next(C);
            }
            deleteFirst(L,P);
        }else{
            C = first(L2);
            while(C != NULL){
                if(parent(C) == P){
                    disconnect(C);

                }
                C = next(C);
            }
            Q = prev(P);
            next(Q) = next(P);
            prev(next(P)) = Q;
            next(P) = NULL;
            prev(P) = NULL;
        }
    }else{
        cout<<"Maaf Data Tidak ditemukan"<<endl;
    }
}
void Hapus_dan_Ganti_produk(List_P &L, List_C &L2, Deskripsi z, Great_Data s){
    address_C C;
    address_P P;
    P = first(L);
    while(info(P).Nama != z.Nama){
        P = next(P);
    }
    if(info(P).Nama == z.Nama){
        C = first(L2);
        while(C != NULL){
            if(parent(C) == P && info(C).Barang == s.Barang){
                disconnect(C);
            }
            C = next(C);
        }
    }else{
        cout<<"Undefinded"<<endl;
    }

}
/*void Tambah_Data_Barang_Jasa2(List_P L, List_C L2, string a){
    address_P P;
    address_C C;
    string con;
    int i;
    P = findElement(L,a);
    if(P != NULL){
        C = first(L2);
        while(C != NULL){
            if(parent(C) == P){
                i = 0;
                con = "Nice";
                while(con != "stop"){
                    cout<<"Masukan Barang ke "<<i<<":";
                    cin>>info(C).Barang[i];
                    cout<<endl;
                    cout<<"Apakah masih ingin menambahkan daftar barang ?"<<endl;
                    cout<<"Input [stop] jika berhenti, input[apapun] jika lanjut"<<endl;
                    cin>>con;
                    i = i+1;
                }
                cout<<endl;
                cout<<"[~Input Data barang selesai~]"<<endl;
                i = 0;
                con = "Nice";
                cout<<endl;
                while(con != "stop"){
                    cout<<"Masukan Daftar Jasa ke "<<i<<":";
                    cin>>info(C).jasa[i];
                    cout<<endl;
                    cout<<"Apakah masih ingin menambahkan daftar jasa ?"<<endl;
                    cout<<"Input [stop] jika berhenti, input[apapun] jika lanjut"<<endl;
                    cin>>con;
                    i = i+1;
                }
                cout<<endl;
                cout<<"[~Input Data Jasa selesai~]"<<endl;

            }
            C = next(C);
        }
    }
}
void Pindah_Data(List_P L, List_C &L2, Deskripsi x){
    address_P P;
    address_C C;
    P = first(L);
    while(next(P) != first(L)){
        C = first(L2);
        while(C != NULL){
            if(parent(C) == P){
                info(C).nama_Per = info(P).Nama;
            }
            C = next(C);
        }
        P = next(P);
    }
    C = first(L2);
    while(C != NULL){
        if(parent(C) == P){
            info(C).nama_Per = info(P).Nama;
        }
        C = next(C);
    }
}*/




