package main

import (
	"fmt"
	"os"
)

var (
	pilih int
	Histori []string
)

type buku struct {
	NamaBuku string
	JumlahBuku int
	}

var (
		DaftarBuku = []buku{
			buku{
				NamaBuku : "Pemrograman",
				JumlahBuku : 10,
			},
			buku{
				NamaBuku : "Film",
				JumlahBuku : 5,
			},
			buku{
				NamaBuku : "Printing",
				JumlahBuku : 20,
				},
			} 
			)

func main() {
	var (
		InputUsername , InputPassword string
		username = "Mutia"
		password = "2406359430"
		pilih int
	)

fmt.Println("===============================================")
fmt.Println("==== Selamat Datang di Perpustakaan Vokasi ====")
fmt.Println("===============================================")

// Input Username dan Password
fmt.Print("Silakan Input Username : ")
fmt.Scanln(&InputUsername)

fmt.Print("Silakan Input Password : " )
fmt.Scanln(&InputPassword) 
fmt.Println("===============================================")

if InputUsername != username && InputPassword != password {
	fmt.Println("Username atau Password Salah !")
	return
}
fmt.Println("===============================================")
for {
		fmt.Println("=== MENU ===")
		fmt.Println("Silakan Pilih Menu : ")
		fmt.Println("1. Lihat Informasi Pengguna Program")
		fmt.Println("2. Lihat Daftar Buku")
		fmt.Println("3. Tambah Daftar Buku")
		fmt.Println("4. Tambah Peminjaman Buku")
		fmt.Println("5. Histori Peminjaman Buku")
		fmt.Println("6. Keluar dari Program")
		fmt.Print("Input Menu yang Anda Inginkan : ") 
		fmt.Scan(&pilih)
fmt.Println("===============================================")

switch pilih {
case 1:
	LihatInformasiPenggunaProgram()
case 2:
	LihatDaftarBuku()
case 3:
	TambahDaftarBuku()
case 4:
	TambahPeminjamanBuku()
case 5:
	HistoriPeminjamanBuku()
case 6:
	KeluarDariProgram()
	default: 
	fmt.Println("PILIHAN TIDAK VALID! SILAKAN PILIH NOMOR SESUAI MENU")
}
}
}

func LihatInformasiPenggunaProgram() {
	
	fmt.Println("=== INFORMASI AKUN ===")
		fmt.Println("Nama          		: Mutia Zahra")
		fmt.Println("NPM           		: 2406359430")
		fmt.Println("Username      		: Mutia")
		fmt.Println("Jenis Kelamin 		: Perempuan")
		fmt.Println("Makanan Favorit    : Soto Banjar")
		fmt.Println("Minuman Favorit    : Air Putih")
}


func LihatDaftarBuku() {
	fmt.Println("=== Daftar Buku ===")
	for _, buku := range DaftarBuku {
		fmt.Printf("Nama Buku : %s, Jumlah %d\n", buku.NamaBuku , buku.JumlahBuku)
	}
}


func TambahDaftarBuku() {
	fmt.Println("=== Tambah Daftar Buku ===")
	var namaBuku string
	var jumlahBuku int
	
	fmt.Print("Masukkan nama buku yang ingin ditambah : ")
	fmt.Scan(&namaBuku)
	fmt.Print("Masukkan jumlah buku yang ingin ditambah : ")
	fmt.Scan(&jumlahBuku)

	if jumlahBuku <= 0 {
	fmt.Println("Jumlah Buku Harus lebih dari 0 !") 
	} else {
	for i, buku := range DaftarBuku {
	if buku.NamaBuku == namaBuku {
			DaftarBuku[i].JumlahBuku += jumlahBuku
			Histori = append(Histori, fmt.Sprintf("Ditambah : %s, Jumlah %d", buku.NamaBuku, jumlahBuku))
			fmt.Println("Buku berhasil ditambahkan")
			break
			}
		}
	}
}

func TambahPeminjamanBuku() {
	fmt.Println("=== Tambah Peminjaman Buku ===")
	var namaBuku string
	var JumlahPinjam int
	fmt.Print("Masukkan nama buku yang ingin dipinjam : ")
	fmt.Scan(&namaBuku)
	fmt.Print("Masukkan jumlah buku yang ingin dipinjam : ")
	fmt.Scan(&JumlahPinjam)

	for i, buku := range DaftarBuku {
		if buku.NamaBuku == namaBuku {
			if JumlahPinjam <= 0 {
				fmt.Println("Jumlah peminjaman harus lebih besar dari 0.")
				break
		} else if buku.JumlahBuku < JumlahPinjam {
			fmt.Println("JUMLAH BUKU TIDAK MENCUKUPI !")
			break
		} else {
			DaftarBuku[i].JumlahBuku -= JumlahPinjam
			Histori = append(Histori, fmt.Sprintf("Dipinjam : %s, Jumlah : %d", buku.NamaBuku, JumlahPinjam))
			fmt.Println("BUKU BERHASIL DIPINJAM")
			break
			}
		}
	}
}

func HistoriPeminjamanBuku() {
	fmt.Println("Histori Peminjaman Buku")
		for _, pinjaman := range Histori {
			fmt.Println(pinjaman)
	}
}

func KeluarDariProgram() {
	fmt.Println("Terima kasih telah mengunjungi Perpustakaan Vokasi!")
	os.Exit(1) 
}