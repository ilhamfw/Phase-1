package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
	_ "github.com/go-sql-driver/mysql"// Driver MySQL
)

// Definisikan struktur data untuk buku
type Buku struct {
	ID      int
	Judul   string
	Penulis string
}

// Definisikan struktur data untuk anggota
type Anggota struct {
	ID   int
	Nama string
}

// Definisikan struktur data untuk peminjaman
type Peminjaman struct {
	ID           int
	AnggotaID    int
	TanggalPinjam time.Time
}

// Definisikan struktur data untuk detail peminjaman
type DetailPeminjaman struct {
	ID             int
	PeminjamanID   int
	BukuID         int
	TanggalKembali time.Time
	Denda          float64
}

func main() {
	// Membuka koneksi ke database MySQL
	db, err := sql.Open("mysql", "root@tcp(localhost:3307)/perpustakaan_kota")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Memanggil fungsi utama untuk menampilkan menu dan interaksi dengan staf perpustakaan
	menuUtama(db)
}

func menuUtama(db *sql.DB) {
    for {
        // Menampilkan menu utama
        fmt.Println("===== Perpustakaan Kota - Aplikasi CLI =====")
        fmt.Println("1. Tampilkan Daftar Buku")
        fmt.Println("2. Tampilkan Daftar Peminjaman")
        fmt.Println("3. Pinjam Buku")
        fmt.Println("4. Kembali Buku")
        fmt.Println("5. TampilkanAnggotaTertinggi")
        fmt.Println("6. tampilkanDetailPeminjaman")
        fmt.Println("7. Keluar")
        fmt.Print("Pilih menu (1/2/3/4/5): ")

        var pilihan int
        fmt.Scanln(&pilihan)

        switch pilihan {
        case 1:
            tampilkanDaftarBuku(db)
        case 2:
            tampilkanDaftarPeminjaman(db)
        case 3:
            pinjamBuku(db)
        case 4:
            kembaliBuku(db)
		case 5:
			tampilkanAnggotaTertinggi(db)
		case 6 :
			tampilkanDetailPeminjaman(db)
        case 7:
            os.Exit(0)
        default:
            fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
        }
    }
}

