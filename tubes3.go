package main
import (
		"fmt"
		"os"
		"os/exec"
		"runtime"
		//"time"
		//"bufio"
		//"log"
		//"os"
)
var clear map[string]func()
func init(){
	clear = make(map[string]func()) // proses instalasi
	clear ["linux"] = func() {
		cmd := exec.Command("clear") // contoh di linux, sudah ditest bung
		cmd.Stdout = os.Stdout
		cmd.Run()
		}
	clear ["windows"] = func () {
		cmd := exec.Command ("cmd", "/c", "cls") // contoh di Windows, sudah juga di ditest
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
func CallClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is unsupported!")
	}
}

type calon struct{
	Namacalon string
	Partai string
	Nomor int
	Threshold int
}

type pemilih struct{
	Nama string
	Gender string
	Umur int
	Pilihan int
	TPS int
}
const N int = 100
const PT int = 3

type arcalon [N] calon
type arpemilih [N] pemilih

var con int

func main (){
	var calonpemimpin arcalon
	var pemilihnya arpemilih
	var manusia, npemilih int
	var esc bool=false
	calonpemimpin[0]=calon{Namacalon:"Prabowo",Partai:"Gerindra",Nomor:1,Threshold:6}
	calonpemimpin[1]=calon{Namacalon:"Jokowi",Partai:"PDIP",Nomor:2,Threshold:     5}
	calonpemimpin[2]=calon{Namacalon:"Bahrul",Partai:"Bem Tel U",Nomor:3,Threshold:4}
	ncalon:=3
	pemilihnya[0]=pemilih{Nama:"Indira",Gender:"Wanita",Umur:19,Pilihan:6,TPS:111}
	pemilihnya[1]=pemilih{Nama:"Atikah",Gender:"Wanita",Umur:19,Pilihan:2,TPS:111}
	pemilihnya[2]=pemilih{Nama:"Naufal",Gender:"Pria",Umur:19,Pilihan:2,TPS:111}

	pemilihnya[3]=pemilih{Nama:"Mandar",Gender:"Pria",Umur:19,Pilihan:1,TPS:222}
	pemilihnya[4]=pemilih{Nama:"Nuriz",Gender:"Wanita",Umur:19,Pilihan:1,TPS:222}
	pemilihnya[5]=pemilih{Nama:"Dewa",Gender:"Pria",Umur:19,Pilihan:1,TPS:222}

	pemilihnya[6]=pemilih{Nama:"Reikiko",Gender:"Pria",Umur:19,Pilihan:7,TPS:333}
	pemilihnya[7]=pemilih{Nama:"Nanda",Gender:"Wanita",Umur:19,Pilihan:3,TPS:333}
	pemilihnya[8]=pemilih{Nama:"Blamma",Gender:"Pria",Umur:19,Pilihan:0,TPS:333}

	pemilihnya[9]=pemilih{Nama:"Hawa",Gender:"Wanita",Umur:19,Pilihan:2,TPS:111}
	pemilihnya[10]=pemilih{Nama:"Sitti",Gender:"Wanita",Umur:19,Pilihan:2,TPS:111}
	pemilihnya[11]=pemilih{Nama:"Yais",Gender:"Pria",Umur:19,Pilihan:2,TPS:111}

	pemilihnya[12]=pemilih{Nama:"Maulana",Gender:"Pria",Umur:19,Pilihan:1,TPS:222}
	pemilihnya[13]=pemilih{Nama:"Annisa",Gender:"Wanita",Umur:19,Pilihan:1,TPS:222}
	pemilihnya[14]=pemilih{Nama:"Dika",Gender:"Pria",Umur:19,Pilihan:1,TPS:222}

	pemilihnya[15]=pemilih{Nama:"Raditya",Gender:"Pria",Umur:19,Pilihan:3,TPS:333}
	pemilihnya[16]=pemilih{Nama:"Nadin",Gender:"Wanita",Umur:19,Pilihan:3,TPS:333}
	pemilihnya[17]=pemilih{Nama:"Ceb",Gender:"Pria",Umur:19,Pilihan:3,TPS:333}
	npemilih=18
	for !esc{
	CallClear()
	fmt.Println("\n\n\n\n\n\n\n")
	fmt.Println("		███████╗███████╗██╗      █████╗ ███╗   ███╗ █████╗ ████████╗    ██████╗  █████╗ ████████╗ █████╗ ███╗   ██╗ ██████╗     ██████╗ ██╗  ")                                                                                                                
	fmt.Println("		██╔════╝██╔════╝██║     ██╔══██╗████╗ ████║██╔══██╗╚══██╔══╝    ██╔══██╗██╔══██╗╚══██╔══╝██╔══██╗████╗  ██║██╔════╝     ██╔══██╗██║  ")                                                                                                                
	fmt.Println("		███████╗█████╗  ██║     ███████║██╔████╔██║███████║   ██║       ██║  ██║███████║   ██║   ███████║██╔██╗ ██║██║  ███╗    ██║  ██║██║  ")                                                                                                               
	fmt.Println("		╚════██║██╔══╝  ██║     ██╔══██║██║╚██╔╝██║██╔══██║   ██║       ██║  ██║██╔══██║   ██║   ██╔══██║██║╚██╗██║██║   ██║    ██║  ██║██║  ")                                                                                                                
	fmt.Println("		███████║███████╗███████╗██║  ██║██║ ╚═╝ ██║██║  ██║   ██║       ██████╔╝██║  ██║   ██║   ██║  ██║██║ ╚████║╚██████╔╝    ██████╔╝██║  ")                                                                                                                
	fmt.Println("		╚══════╝╚══════╝╚══════╝╚═╝  ╚═╝╚═╝     ╚═╝╚═╝  ╚═╝   ╚═╝       ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═══╝ ╚═════╝     ╚═════╝ ╚═╝  ")
	fmt.Println("\n")
	fmt.Println("		 █████╗ ██████╗ ██╗     ██╗██╗  ██╗ █████╗ ███████╗██╗    ██████╗ ███████╗███╗   ███╗██╗██╗     ██╗██╗  ██╗ █████╗ ███╗   ██╗   ")        
	fmt.Println("		██╔══██╗██╔══██╗██║     ██║██║ ██╔╝██╔══██╗██╔════╝██║    ██╔══██╗██╔════╝████╗ ████║██║██║     ██║██║  ██║██╔══██╗████╗  ██║   ")        
	fmt.Println("		███████║██████╔╝██║     ██║█████╔╝ ███████║███████╗██║    ██████╔╝█████╗  ██╔████╔██║██║██║     ██║███████║███████║██╔██╗ ██║   ")        
	fmt.Println("		██╔══██║██╔═══╝ ██║     ██║██╔═██╗ ██╔══██║╚════██║██║    ██╔═══╝ ██╔══╝  ██║╚██╔╝██║██║██║     ██║██╔══██║██╔══██║██║╚██╗██║   ")        
	fmt.Println("		██║  ██║██║     ███████╗██║██║  ██╗██║  ██║███████║██║    ██║     ███████╗██║ ╚═╝ ██║██║███████╗██║██║  ██║██║  ██║██║ ╚████║   ")        
	fmt.Println("		╚═╝  ╚═╝╚═╝     ╚══════╝╚═╝╚═╝  ╚═╝╚═╝  ╚═╝╚══════╝╚═╝    ╚═╝     ╚══════╝╚═╝     ╚═╝╚═╝╚══════╝╚═╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝   ") 
	fmt.Println("\n")
	fmt.Println("		██╗   ██╗███╗   ███╗██╗   ██╗███╗   ███╗    ██╗     ███████╗ ██████╗ ██╗███████╗██╗      █████╗ ████████╗██╗███████╗  ")
	fmt.Println("		██║   ██║████╗ ████║██║   ██║████╗ ████║    ██║     ██╔════╝██╔════╝ ██║██╔════╝██║     ██╔══██╗╚══██╔══╝██║██╔════╝  ")
	fmt.Println("		██║   ██║██╔████╔██║██║   ██║██╔████╔██║    ██║     █████╗  ██║  ███╗██║███████╗██║     ███████║   ██║   ██║█████╗    ")
	fmt.Println("		██║   ██║██║╚██╔╝██║██║   ██║██║╚██╔╝██║    ██║     ██╔══╝  ██║   ██║██║╚════██║██║     ██╔══██║   ██║   ██║██╔══╝    ")
	fmt.Println("		╚██████╔╝██║ ╚═╝ ██║╚██████╔╝██║ ╚═╝ ██║    ███████╗███████╗╚██████╔╝██║███████║███████╗██║  ██║   ██║   ██║██║       ")
	fmt.Println(" 		 ╚═════╝ ╚═╝     ╚═╝ ╚═════╝ ╚═╝     ╚═╝    ╚══════╝╚══════╝ ╚═════╝ ╚═╝╚══════╝╚══════╝╚═╝  ╚═╝   ╚═╝   ╚═╝╚═╝       ")
	fmt.Println("\n")
	fmt.Print(" ")
	fmt.Scanln(&con)
	CallClear()
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Println ("								Siapa anda? Apa urusan anda?")
	fmt.Println("\n")
	fmt.Println ("								1. Saya Pemilih, Saya ingin Memilih!")
	fmt.Println("\n")
	fmt.Println ("								2. Saya Petugas, Apa urusan anda menanyakan itu?")
	fmt.Println("\n")
	fmt.Println ("								3. Keluar")
	fmt.Println("\n")
	fmt.Print("								Masukan Pilihan : ")
	fmt.Scanln (&manusia)
	switch manusia {
	case 1:
		CallClear()
		pemilihh(&calonpemimpin,&pemilihnya,&ncalon,&npemilih)
	case 2:
		CallClear()
		petugas(&calonpemimpin,&pemilihnya,&ncalon,&npemilih)
	case 3:
		esc=true
	}}
}

func pemilihh(ahh *arcalon, nawn *arpemilih, ncalon *int, npemilih *int){
	var menupemilih int
	var keluar bool = false
	var n int = *npemilih
	fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
	fmt.Print("								Nama Pemilih  : ")
	fmt.Scan(&nawn[n].Nama)
	fmt.Println("\n")
	fmt.Print("								Jenis Kelamin   : ")
	fmt.Scan(&nawn[n].Gender)
	fmt.Println("\n")
	fmt.Print("								Umur      : ")
	fmt.Scan(&nawn[n].Umur)
	for !keluar {
		CallClear()
		fmt.Println("\n\n\n\n\n\n\n\n\n\n")
		fmt.Println("								******************************")
		fmt.Println("								******************************")
		fmt.Println("								***------------------------***")
		fmt.Println("								*** M E N U  P E M I L I H ***")
		fmt.Println("								***------------------------***")
        fmt.Println("								***                        ***")
		fmt.Println("								***(1) LIHAT CALON         ***")
		fmt.Println("								***(2) MULAI MEMILIH       ***")
		fmt.Println("								***(3) CARI DATA           ***")
		fmt.Println("								***(4) KELUAR              ***")
		fmt.Println("								***                        ***")
		fmt.Println("								******************************")
		fmt.Println("								******************************")
		fmt.Print("								Pilihan 	: ")
		fmt.Scanln(&menupemilih)
		if menupemilih == 1{
			lihatData(ahh,ncalon)
		} else if menupemilih==4{
				keluar=true
		} else if menupemilih==2{
				fmt.Print("								Masukkan TPS : ")
				fmt.Scanln(&nawn[n].TPS)
				fmt.Print("								Masukkan Pilihan : ")
				fmt.Scanln(&nawn[n].Pilihan)

				for j:=0;j<n;j++{
					if nawn[n].Pilihan==ahh[j].Nomor{
						ahh[j].Threshold++
					}
				}
		} else if menupemilih==3{
			find(ahh,ncalon)
		}
	}
	*npemilih++
}
//}

func petugas(sahh *arcalon, nawn *arpemilih, n*int, npemilih *int){
	var menupetugas int
	var esc bool = false
	for !esc {
		CallClear()
		fmt.Println("\n\n\n\n\n\n\n\n\n\n")
		fmt.Println("								***--------------------------***")
		fmt.Println("								***----------M E N U---------***")
		fmt.Println("								***    (1) MASUKAN CALON     ***")
		fmt.Println("								***    (2) HAPUS CALON       ***")
		fmt.Println("								***    (3) UBAH CALON        ***")
		fmt.Println("								***    (4) LIHAT CALON       ***")
		fmt.Println("								***    (5) CARI CALON        ***")
		fmt.Println("								***    (6) URUTKAN CALON     ***")
		fmt.Println("								***    (7) HASIL TPS         ***")
		fmt.Println("								***    (8) CEK GOLPUT        ***")
		fmt.Println("								***    (9) KELUAR	     ***")
		fmt.Println("								***--------------------------***")
		fmt.Println("								***--------------------------***")
		fmt.Print("								Pilihan 	: ")
		fmt.Scanln(&menupetugas)
		switch menupetugas {
		case 1:insert(sahh,n)
		case 2:delete(sahh,n)
		case 3:edit(sahh,n)
		case 4:lihatData(sahh,n)
		case 5:find(sahh,n)
	    case 6:urut(sahh, n)
		case 7:hasil(sahh,n)
		case 8:golput(sahh,nawn,n,npemilih)
		case 9:esc=true
		}
	}
}

func edit(editcalon *arcalon, ncalon *int){
	var pilihedit, editnomor int
	fmt.Print("								Masukkan Nomor Calon yang mau di edit : ")
	fmt.Scanln(&editnomor)
	for i:=0;i<*ncalon;i++{
		if editnomor==editcalon[i].Nomor{
			fmt.Println("								Pilih Data yang maw di edit : ")
			fmt.Println("								1. Nama Calon")
			fmt.Println("								2. Partai Calon")
			fmt.Println("								3. Nomor Calon")
			fmt.Print("								Pilihan : ")
			fmt.Scanln(&pilihedit)
			if pilihedit == 1{
				fmt.Print("								Masukkan Nama Calon yang baru : ")
				fmt.Scanln(&editcalon[i].Namacalon)
			} else if pilihedit == 2{
				fmt.Print("								Masukkan Partai Calon yang baru : ")
				fmt.Scanln(&editcalon[i].Partai)
			} else if pilihedit == 3{
				fmt.Print("								Masukkan Nomor Calon yang baru : ")
				fmt.Scanln(&editcalon[i].Nomor)
			}
		}
	}
}

func insert(masukcalon *arcalon, ncalon *int){
	fmt.Print("								Nama Calon Baru		: ")
	fmt.Scanln(&masukcalon[*ncalon].Namacalon)
	fmt.Print("								Partai Calon Baru	: ")
	fmt.Scanln(&masukcalon[*ncalon].Partai)
	fmt.Print("								Nomor Calon Baru	: ")
	fmt.Scanln(&masukcalon[*ncalon].Nomor)
	*ncalon++
	fmt.Println("								\tOK")
	fmt.Print("								Press any key to continue ")
	fmt.Scanln(&con)
}

func lihatData(lihatcalon *arcalon, ncalon *int){
	fmt.Println("								Nomor\t\tNama Calon\t\tPartai\n")
	fmt.Print("								----------------------------------------------------------\n")
	for i:=0;i<*ncalon;i++{
		fmt.Println("\t\t\t\t\t\t\t\t",lihatcalon[i].Nomor,"\t\t",lihatcalon[i].Namacalon,"\t\t",lihatcalon[i].Partai,"\n")
	}
	fmt.Print("								Press any key to continue ")
	fmt.Scanln(&con)
}

func delete(hpscalon *arcalon, ncalon *int){
	var hapus int
	fmt.Print("								Masukkan Nomor yang di hapus	: ")
	fmt.Scanln(&hapus)
	for j:=0;j<*ncalon;j++{
		if hapus == hpscalon[j].Nomor{
			hpscalon[j]=hpscalon[j+1]
		}
	}
	*ncalon--
	fmt.Println("\n\t\t\t\t\t\t\t\tData sudah dihapus! Silahkan cek menu Lihat Calon")
	fmt.Print("\n")
	fmt.Print("								Press any key to continue ")
	fmt.Scanln(&con)
}

func find(cari *arcalon, ncalon *int){
	var pilihcari,i int
	var found bool
	fmt.Println("								Pilih data yang mau dicari : ")
	fmt.Println("								1. Nama Calon")
	fmt.Println("								2. Partai Calon")
	fmt.Println("								3. Nomor Calon")
	fmt.Println("								!!!CASE SENSITIVE!!!")
	fmt.Print("								Pilihan : ")
	fmt.Scanln(&pilihcari)
	if pilihcari==1{
		var carinama string
		fmt.Print("								Cari Nama : ")
		fmt.Scanln(&carinama)
		i=0
		found=false
		for i<*ncalon && !found{
			if carinama==cari[i].Namacalon{
				fmt.Print("								\t\tNomor\t\tNama Calon\tPartai\n")
				fmt.Print("								\t\t--------------------------------------------------\n")
				fmt.Print("								\t\t",cari[i].Nomor,"\t\t",cari[i].Namacalon,"\t\t",cari[i].Partai,"\n")
				found=true
			}
			i++
		}
		if found==false {
			fmt.Println("								DATA TIDAK ADA")
		}
		fmt.Print("								Press any key to continue")
		fmt.Scanln(&con)
	} else if pilihcari==2{
		var caripartai string
		fmt.Print("								Cari Partai : ")
		fmt.Scanln(&caripartai)
		i=0
		found=false
		for i<*ncalon && !found{
			if caripartai==cari[i].Partai{
				fmt.Print("								Nomor\t\tNama Calon\tPartai\n")
				fmt.Print("								--------------------------------------------------\n")
				fmt.Print("								",cari[i].Nomor,"\t\t",cari[i].Namacalon,"\t\t",cari[i].Partai,"\n")
				found=true
			}
			i++
		}
		if found==false {
			fmt.Println("								DATA TIDAK ADA")
		}
		fmt.Print("								Press any key to continue")
		fmt.Scanln(&con)
	} else if pilihcari==3{
		var carinomor int
		fmt.Print("								Cari Nomor : ")
		fmt.Scanln(&carinomor)
		i=0
		found=false
		for i<*ncalon && !found{
			if carinomor==cari[i].Nomor{
				fmt.Print("								Nomor\t\tNama Calon\tPartai\n")
				fmt.Print("								-----------------------------------------------\n")
				fmt.Print("								",cari[i].Nomor,"\t\t",cari[i].Namacalon,"\t\t",cari[i].Partai,"\n")
				found=true
			}
			i++
		}
		if found==false {
			fmt.Println("								DATA TIDAK ADA")
		}
		fmt.Print("								Press any key to continue")
		fmt.Scanln(&con)
	}

}

func hasil(hasilcalon *arcalon, ncalon *int){
	var max calon
	for i:=0;i<*ncalon;i++{
		if hasilcalon[i].Threshold>PT{
			max=hasilcalon[i]
			for j:=0;j<*ncalon;j++{
				if hasilcalon[j].Threshold>max.Threshold{
					max=hasilcalon[j]
				}
			}
		}
	}
	fmt.Println("								Pemenang Pemilu adalah ",max.Namacalon," dengan perolehan suara",max.Threshold,"\n")
	fmt.Print("								Press any key to continue ")
	fmt.Scanln(&con)
}

func golput(ahh *arcalon, nawn *arpemilih, ncalon *int, npemilih *int){
	var tps111 arpemilih
	var tps222 arpemilih
	var tps333 arpemilih
	var nopilihan[N] int
	var j int
	for j=0;j<*ncalon;j++{
		nopilihan[j]=ahh[j].Nomor
	}

	var tpsberapa int
	var golput int = 0
	var k int = 0
	fmt.Println("								Mau liat golput di TPS berapa?")
	fmt.Println("								1. TPS 111")
	fmt.Println("								2. TPS 222")
	fmt.Println("								3. TPS 333")
	fmt.Print("								Pilihannya : ")
	fmt.Scanln(&tpsberapa)
	if tpsberapa==1 {
		for i:=0;i<*npemilih;i++{
			if nawn[i].TPS==111{
				tps111[k]=nawn[i]
				k++
			}
		}
		var b,c int
		for b=0;b<j;b++{
			for c=0;c<=k;c++{
				if nopilihan[b]==tps111[c].Pilihan{

				}else if nopilihan[b]!=tps111[c].Pilihan {
					golput++
				}
			}
		}
		fmt.Println("\t\t\t\t\t\t\t\t",b,k,golput)
		golput=golput-((k*b)-3)
		fmt.Println("								Jumlah Golput di TPS 111 adalah ",golput," dari ",k," orang")
	} else if tpsberapa==2{
		for i:=0;i<*npemilih;i++{
			if nawn[i].TPS==222{
				tps222[k]=nawn[i]
				k++
			}
		}
		var b,c int
		for b=0;b<j;b++{
			for c=0;c<=k;c++{
				if nopilihan[b]==tps222[c].Pilihan{

				}else if nopilihan[b]!=tps222[c].Pilihan {
					golput++
				}
			}
		}
		fmt.Println("\t\t\t\t\t\t\t\t",b,k,golput)
		golput=golput-((k*b)-3)
		fmt.Println("								Jumlah Golput di TPS 222 adalah ",golput," dari ",k," orang")
	}	else if tpsberapa==3 {
		for i:=0;i<*npemilih;i++{
			if nawn[i].TPS==333{
				tps333[k]=nawn[i]
				k++
			}
		}
		var b,c int
		for b=0;b<j;b++{
			for c=0;c<=k;c++{
				if nopilihan[b]==tps333[c].Pilihan{

				}else if nopilihan[b]!=tps333[c].Pilihan {
					golput++
				}
			}
		}
		fmt.Println("\t\t\t\t\t\t\t\t",b,k,golput)
		golput=golput-((k*b)-3)
		fmt.Println("								Jumlah Golput di TPS 333 adalah ",golput," dari ",k," orang")
	}
	fmt.Print("								Press any key to continue ")
	fmt.Scanln(&con)
}

func urut(hasilcalon *arcalon, ncalon *int){
	var max calon
	for i:=0;i<*ncalon;i++{
		if hasilcalon[i].Threshold>PT{
			max=hasilcalon[i]
			for j:=0;j < *ncalon;j++{
				if hasilcalon[j].Threshold > max.Threshold{
					max.Threshold = max.Threshold + 1
					max=hasilcalon[j]
				}
			}
			max.Threshold++
		}
	}
	fmt.Println("								Nomor\t\tNama Calon\t\tPartai\t\tPerolehan\n")
	fmt.Print("								------------------------------------------------------------------\n")
	for i:=0;i<*ncalon;i++{
		fmt.Println("\t\t\t\t\t\t\t\t",hasilcalon[i].Nomor,"\t\t",hasilcalon[i].Namacalon,"\t\t",hasilcalon[i].Partai,"\t\t",hasilcalon[i].Threshold,"\n")
	}
	fmt.Print("								Press any key to continue ")
	fmt.Scanln(&con)
}

/*func selection_sort(arr *array, n int){
	var min_index int

	for i := 0; i < n; i++ {
		min_index = i
		for j := i + 1; j < n; j++ {
			//urutan besar ke kecil atau kecil ke besar (perhatikan tanda)
			if arr[j] < arr [min_index] {
				min_index = j

		    }
	    }
	    temp := arr[min_index]
	    arr[min_index] = arr[i]
	    arr[i] = temp
  }

}

func UrutkanData(arr *array, n int){

	var index array
	var m int

	fmt.Scanln(&n)
	i := 0
	for (i < n){
		j := 0
		fmt.Scan(&m)
		for ( j < m){
			fmt.Scan(&index[j])
			j++
	    }
		selection_sort(&index, m)
		for k := 0; k < m; k++ {
			fmt.Print(index[k], " ")

	    }

	}

}
*/
