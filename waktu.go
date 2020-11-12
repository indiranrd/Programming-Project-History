package main

import (
		"fmt"
		"time"
		"bufio"
		"log"
		"os"
)

func getInput(input chan string) {
    for {
        in := bufio.NewReader(os.Stdin)
        result, err := in.ReadString('\n')
        if err != nil {
            log.Fatal(err)
        }

        input <- result
    }
}

func main() {
/*var Pilihan int
var	Desc string
	fmt.Println(" .----------------.        .----------------.        .----------------.        .----------------. ")
	fmt.Println("| .--------------. |      | .--------------. |      | .--------------. |      | .--------------. |")
	fmt.Println("| |     __       | |      | |    _____     | |      | |    ______    | |      | |   _    _     | |")
	fmt.Println("| |    /  |      | |      | |   / ___ `.   | |      | |   / ____ `.  | |      | |  | |  | |    | |")
	fmt.Println("| |    `| |      | |      | |  |_/___) |   | |      | |   `'  __) |  | |      | |  | |__| |_   | |")
	fmt.Println("| |     | |      | |      | |   .'____.'   | |      | |   _  |__ '.  | |      | |  |____   _|  | |")
	fmt.Println("| |    _| |_     | |      | |  / /____     | |      | |  | :____) |  | |      | |      _| |_   | |")
	fmt.Println("| |   |_____|    | |      | |  |_______|   | |      | |   :______.'  | |      | |     |_____|  | |")
	fmt.Println("| |              | |      | |              | |      | |              | |      | |              | |")
	fmt.Println("| '--------------' |      | '--------------' |      | '--------------' |      | '--------------' |")
	fmt.Println(" '----------------'        '----------------'        '----------------'        '----------------' ")
    fmt.Println("KETERANGAN :")
    fmt.Println("1. Mandar")
    fmt.Println("2. Nescafe")
    fmt.Println("3. Teh saha?")
    fmt.Printf("4. Mulai Voting")

    fmt.Printf("\n\nPilih Angka = ")
    timer1 := time.NewTimer(10 * time.Second)

    <-timer1.C
    fmt.Println("WAKTU HABIS")

    timer2 := time.NewTimer(3 * time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 expired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("SILAHKAN MEMULAI ULANG PROGRAM")

    fmt.Scanln(&Pilihan)
    if  Pilihan == 1 {
		Desc = "Mandar adalah seorang kapiteng ablakdsfljaosnciohovnkjsNViounskdjnvoiu nsjdvn oinonojKCNVoihdklj nsiojnvoi oiniondoijc oikajfoialkcn"
	}
	if  Pilihan == 2 {
		Desc = "A"
	}
	if  Pilihan == 3 {
		Desc = "A"
	}
	if  Pilihan == 4 {
		Desc = "A"
	}
	fmt.Println ("\n\nMANDAR\n",Desc)

	
    }*/
    var calon string
	input := make(chan string, 1)
    go getInput(input)

    
        fmt.Println("masukan nomor calon")

        select {
        case calon = <-input:
            fmt.Println("pemilihan berhasil")
            
        case <-time.After(5000 * time.Millisecond):
            fmt.Println("waktu pemilihan habis")
        }

        fmt.Print(calon)
    
}

/*
package main

import (
		"fmt"
		"time"
)
func main() {
var Pilihan int
var	Desc string
	fmt.Println(" .----------------.        .----------------.        .----------------.        .----------------. ")
	fmt.Println("| .--------------. |      | .--------------. |      | .--------------. |      | .--------------. |")
	fmt.Println("| |     __       | |      | |    _____     | |      | |    ______    | |      | |   _    _     | |")
	fmt.Println("| |    /  |      | |      | |   / ___ `.   | |      | |   / ____ `.  | |      | |  | |  | |    | |")
	fmt.Println("| |    `| |      | |      | |  |_/___) |   | |      | |   `'  __) |  | |      | |  | |__| |_   | |")
	fmt.Println("| |     | |      | |      | |   .'____.'   | |      | |   _  |__ '.  | |      | |  |____   _|  | |")
	fmt.Println("| |    _| |_     | |      | |  / /____     | |      | |  | :____) |  | |      | |      _| |_   | |")
	fmt.Println("| |   |_____|    | |      | |  |_______|   | |      | |   :______.'  | |      | |     |_____|  | |")
	fmt.Println("| |              | |      | |              | |      | |              | |      | |              | |")
	fmt.Println("| '--------------' |      | '--------------' |      | '--------------' |      | '--------------' |")
	fmt.Println(" '----------------'        '----------------'        '----------------'        '----------------' ")
    fmt.Println("KETERANGAN :")
    fmt.Println("1. Mandar")
    fmt.Println("2. Nescafe")
    fmt.Println("3. Teh saha?")
    fmt.Printf("4. Mulai Voting")

    fmt.Printf("\n\nPilih Angka = ")
    timer1 := time.NewTimer(10 * time.Second)

    <-timer1.C
    fmt.Println("WAKTU HABIS")

    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 expired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("SILAHKAN MEMULAI ULANG PROGRAM")

    fmt.Scanln(&Pilihan)
    if  Pilihan == 1 {
		Desc = "Mandar adalah seorang kapiteng ablakdsfljaosnciohovnkjsNViounskdjnvoiu nsjdvn oinonojKCNVoihdklj nsiojnvoi oiniondoijc oikajfoialkcn"
	}
	if  Pilihan == 2 {
		Desc = "A"
	}
	if  Pilihan == 3 {
		Desc = "A"
	}
	if  Pilihan == 4 {
		Desc = "A"
	}
	fmt.Println ("\n\nMANDAR\n",Desc)

	
    }
}
*/