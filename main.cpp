#include <iostream>
#include <stdio.h>
#include <stdlib.h>
#include "Parent_TB.h"
#include "Child_TB.h"
#include "multilist1-N_TB.h"
#include <time.h>

using namespace std;
List_P L;
List_C L2;
int main()

{
    time_t currentTime;
    address_C C;
    address_P P;
    Deskripsi x, z;
    string c, b, a;
    Great_Data y, s;
    int idxmax;


    cout<<"=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-==-=-=-=-="<<endl;
    cout<<"                                                                  "<<endl;
    cout<<"                                                                  "<<endl;
    cout<<"                       -MULTINATIONAL COMPANY-                    "<<endl;
    cout<<"                               [2020]                             "<<endl;
    cout<<"                                                                  "<<endl;
    cout<<"-=-=-=-=-=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-"<<endl;
    time(&currentTime);
    cout<<"[Data Ter-Update : ]"<<endl;
    printf("%s\n",ctime(&currentTime));
    cout<<"Data Masih Kosong"<<endl;
    cout<<endl;
    createList(L2);
    createList(L);
    time(&currentTime);
    cout<<"[Data Ter-Update : ]"<<endl;
    printf("%s\n",ctime(&currentTime));
    cout<<"------------------------------------------------------------------"<<endl;
    cout<<" 1. Ketik 1 Untuk memasukan data perusahaan"<<endl;
    cout<<" 2. Ketik 2 untuk menambahkan data produk dan jasa perusahaan"<<endl;
    cout<<" 3. Ketik 3 untuk menutup perusahaan "<<endl;
    cout<<" 4. Ketik 4 untuk mencari Barang "<<endl;
    cout<<" 5. Ketik 5 untuk mengganti / menghapus data produk Perusahaan"<<endl;
    cout<<" 6. Exit untuk melihat perusahaan top dunia"<<endl;
    cout<<"------------------------------------------------------------------"<<endl;
    time(&currentTime);
    cout<<"[Data Ter-Update : ]"<<endl;
    printf("%s\n",ctime(&currentTime));
    cout<<endl;
    cout<<"[iya] -> Jalankan Program / [apapun] -> matikan program"<<endl;
    cin>>c;
    while(c == "iya"){
        cout<<"Pilih Menu : " ;
        cin>>idxmax;
        if(idxmax == 1){
            cout<<"=================================================================="<<endl;
            cout<<"||          -Apakah anda ingin menambahkan data perusahaan ?-    ||"<<endl;
            cout<<"||              Data yang dimasukan adalah produk unggulan !     ||"<<endl;
            cout<<"=================================================================="<<endl;
            time(&currentTime);
            cout<<"[Data Ter-Update : ]"<<endl;
            printf("%s\n",ctime(&currentTime));
            cout<<endl;
            cout<<"Ketik [iya] jika ingin masukan data : "<<endl;
            cout<<"Ketik [apapun] jika selesai menginputkan data"<<endl;
            cin>>b;
            while(b == "iya"){
                cout<<"Masukan nama Perusahaan          : ";
                cin>>x.Nama;
                cout<<"Masukan Tahun Berdiri            : ";
                cin>>x.Tahun;
                cout<<"Masukan Jumlah Keuntungan Perusahaan : ";
                cin>>x.Laba_Tahun2020;
                insertFirst(L,alokasi(x));
                int i = 1;
                int jum;
                cout<<"Masukan Jumlah Barang Yang Ingin di Tambahkan :  ";
                cin>>jum;
                while(i <= jum){
                    cout<<"Barang ke- " <<i<<endl;
                    cin>>y.Barang;
                    i++;
                    insertFirstChild(L2,alokasi(y));
                    P = findElement(L,x.Nama);
                    C = FindElementChild(L2,y.Barang);
                    connect(P,C);
                }
                    cout<<"Ketik [iya] Jika Ingin Menambahkan Data Perusahaan Lagi "<< endl;
                    cout<<"Ketik [apapun] jika selesai menginputkan data"<<endl;
                    cin>>b;
            }
            cout<<"=================================================================="<<endl;
            cout<<"                                                                  "<<endl;
            cout<<"                                                                  "<<endl;
            cout<<"                 -NEW LIST MULTINATIONAL CORPORATION-             "<<endl;
            cout<<"                               [2020]                             "<<endl;
            cout<<"                                                                  "<<endl;
            cout<<"=================================================================="<<endl;
            time(&currentTime);
            cout<<"[Data Ter-Update : ]"<<endl;
            printf("%s\n",ctime(&currentTime));
            cout<<endl;
            printRelasi(L,L2);
            cout<<endl;
            time(&currentTime);
            cout<<"[Data Ter-Update : ]"<<endl;
            printf("%s\n",ctime(&currentTime));
        }else if(idxmax == 2){
            cout<<"==================================================================================="<<endl;
            cout<<"||          -Apakah anda ingin menambahkan data barang dan jasa perusahaan ?-    ||"<<endl;
            cout<<"==================================================================================="<<endl;
            cout<<endl;
            time(&currentTime);
            cout<<"[Data Ter-Update : ]"<<endl;
            printf("%s\n",ctime(&currentTime));
            cout<<endl;
            cout<<"Input [iya] / [apapun]:"<<endl;
            cin>>b;
            if(b == "iya"){
                cout<<"Ketik Nama Perusahaan yang ingin ditambah datanya :"<<endl;
                cin>>a;
                int i = 1;
                int jum;
                cout<<"Masukan Jumlah Barang Yang Ingin di Tambahkan :  ";
                cin>>jum;
                while(i <= jum){
                    cout<<"Barang ke- " <<i<<endl;
                    cin>>y.Barang;
                    i++;
                    insertFirstChild(L2,alokasi(y));
                    P = findElement(L,a);
                    C = FindElementChild(L2,y.Barang);
                    connect(P,C);
                }

            }
            cout<<"=================================================================="<<endl;
            cout<<"                                                                  "<<endl;
            cout<<"                                                                  "<<endl;
            cout<<"                 -NEW LIST MULTINATIONAL CORPORATION-             "<<endl;
            cout<<"                               [2020]                             "<<endl;
            cout<<"                                                                  "<<endl;
            cout<<"=================================================================="<<endl;
            time(&currentTime);
            cout<<"[Data Ter-Update : ]"<<endl;
            printf("%s\n",ctime(&currentTime));
            cout<<endl;
            printRelasi(L,L2);
            cout<<endl;
            time(&currentTime);
            cout<<"[Data Ter-Update : ]"<<endl;
            printf("%s\n",ctime(&currentTime));
            cout<<endl;
        }else if(idxmax == 3){
            cout<<"==================================================================================="<<endl;
            cout<<"||                             -HAPUS PERUSAHAAN-                                 ||"<<endl;
            cout<<"==================================================================================="<<endl;
            cout<<endl;
            time(&currentTime);
            cout<<"[Data Ter-Update : ]"<<endl;
            printf("%s\n",ctime(&currentTime));
            cout<<endl;
            cout<<"Masukan Nama perusahaan yang ingin dihapus :"<<endl;
            cin>>a;
            Hapus_Perusahaan(L,L2,a);
            cout<<"=================================================================="<<endl;
            cout<<"                                                                  "<<endl;
            cout<<"                                                                  "<<endl;
            cout<<"                 -NEW LIST MULTINATIONAL CORPORATION-             "<<endl;
            cout<<"                               [2020]                             "<<endl;
            cout<<"                                                                  "<<endl;
            cout<<"=================================================================="<<endl;
            time(&currentTime);
            cout<<"[Data Ter-Update : ]"<<endl;
            printf("%s\n",ctime(&currentTime));
            cout<<endl;
            printRelasi(L,L2);
            cout<<endl;
            time(&currentTime);
            cout<<"[Data Ter-Update : ]"<<endl;
            printf("%s\n",ctime(&currentTime));
            cout<<endl;
            cout<<"Apakah anda ingin menutup perusahaan ?"<<endl;
            cout<<" Input [iya]\ [apapun]"<<endl;
            cin>>b;
        }else if(idxmax == 4){
            cout<<"==================================================================================="<<endl;
            cout<<"||                             -CARI BARANG -                                    ||"<<endl;
            cout<<"==================================================================================="<<endl;
            cout<<endl;
            time(&currentTime);
            cout<<"[Data Ter-Update : ]"<<endl;
            printf("%s\n",ctime(&currentTime));
            cout<<endl;
            cout<<"Masukan Barang yang ingin dicari"<<endl;
            cin>>a;
            Relasi_Produk(L,L2,a);
        }else if(idxmax == 5){
            cout<<"==================================================================================="<<endl;
            cout<<"||                             -Hapus / Ganti Produk -                            ||"<<endl;
            cout<<"==================================================================================="<<endl;
            cout<<endl;
            cout<<"Input Nama Perusahaan yang ingin dihapus datanya :"<<endl;
            cin>>z.Nama;
            cout<<"Input Nama Barang yang mau dihapus               :"<<endl;
            cin>>s.Barang;
            Hapus_dan_Ganti_produk(L,L2,z,s);
            printRelasi(L,L2);
            time(&currentTime);
            cout<<"[Data Ter-Update : ]"<<endl;
            printf("%s\n",ctime(&currentTime));
        }
        cout<<"------------------------------------------------------------------"<<endl;
        cout<<" 1. Ketik 1 Untuk memasukan data perusahaan"<<endl;
        cout<<" 2. Ketik 2 untuk menambahkan data dan barang perusahaan"<<endl;
        cout<<" 3. Ketik 3 untuk menutup perusahaan "<<endl;
        cout<<" 4. Ketik 4 untuk mencari Barang "<<endl;
        cout<<" 5. Ketik 5 untuk mengganti / menghapus data produk Perusahaan"<<endl;
        cout<<" 6. Exit untuk melihat perusahaan top dunia"<<endl;
        cout<<"------------------------------------------------------------------"<<endl;
        time(&currentTime);
        cout<<"[Data Ter-Update : ]"<<endl;
        printf("%s\n",ctime(&currentTime));
        cout<<"[iya] untuk jalankan PROGRAM"<<endl;
        cin>>c;
        }
    Perusahaan_Profit(L,x);
    cout<<"[Program Selesai ]"<<endl;
    cout<<"[This Program made by Adrian and Indira]"<<endl;
    cout<<"Copyright(c)2020 Arman Indira all right reserved"<<endl;
    return 0;

}
