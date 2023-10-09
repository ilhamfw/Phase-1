package main

import (
    "database/sql"
    "fmt"
    "log"
	"time"
)

func tampilkanDaftarBuku(db *sql.DB) {
    // Query untuk mengambil daftar buku dari database
    rows, err := db.Query("SELECT buku_id, judul, penulis FROM buku")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // Menampilkan daftar buku
    fmt.Println("===== Daftar Buku =====")
    for rows.Next() {
        var buku Buku
        err := rows.Scan(&buku.ID, &buku.Judul, &buku.Penulis)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("ID: %d, Judul: %s, Penulis: %s\n", buku.ID, buku.Judul, buku.Penulis)
    }

    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }
}


func tampilkanDaftarPeminjaman(db *sql.DB) {
    // Query untuk mengambil daftar peminjaman dari database
    rows, err := db.Query("SELECT peminjaman_id, anggota_id, tanggal_pinjam FROM peminjaman")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // Menampilkan daftar peminjaman
    fmt.Println("===== Daftar Peminjaman =====")
    for rows.Next() {
        var peminjaman Peminjaman
        var tanggalPinjamStr string // Temporary variable to store the date string
        err := rows.Scan(&peminjaman.ID, &peminjaman.AnggotaID, &tanggalPinjamStr)
        if err != nil {
            log.Fatal(err)
        }
        
        // Parse the date string into a time.Time value
        peminjaman.TanggalPinjam, err = time.Parse("2006-01-02", tanggalPinjamStr)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Printf("ID: %d, Anggota ID: %d, Tanggal Pinjam: %s\n", peminjaman.ID, peminjaman.AnggotaID, peminjaman.TanggalPinjam.Format("2006-01-02"))
    }

    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }
}


func pinjamBuku(db *sql.DB) {
    // Memasukkan data anggota dan buku yang ingin dipinjam
    var anggotaID int
    var bukuID int

    fmt.Print("Masukkan ID Anggota: ")
    fmt.Scanln(&anggotaID)
    
    fmt.Print("Masukkan ID Buku yang ingin dipinjam: ")
    fmt.Scanln(&bukuID)

    // Mendapatkan tanggal saat ini
    tanggalPinjam := time.Now()

    // Memasukkan data peminjaman ke dalam tabel peminjaman
    _, err := db.Exec("INSERT INTO peminjaman (anggota_id, tanggal_pinjam) VALUES (?, ?)", anggotaID, tanggalPinjam)
    if err != nil {
        log.Fatal(err)
    }

    // Mendapatkan ID peminjaman yang baru saja dimasukkan
    var peminjamanID int
    err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&peminjamanID)
    if err != nil {
        log.Fatal(err)
    }

    // Memasukkan data detail peminjaman ke dalam tabel detail_peminjaman
    _, err = db.Exec("INSERT INTO detail_peminjaman (peminjaman_id, buku_id, tanggal_kembali) VALUES (?, ?, ?)", peminjamanID, bukuID, tanggalPinjam.AddDate(0, 0, 7)) // Mengatur tanggal kembali 7 hari dari tanggal pinjam
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Buku berhasil dipinjam!")
}


func kembaliBuku(db *sql.DB) {
    // Meminta input ID peminjaman dari pengguna
    var peminjamanID int
    fmt.Print("Masukkan ID Peminjaman: ")
    fmt.Scanln(&peminjamanID)

    // Prepare the SQL statement with a placeholder to retrieve "tanggal_pinjam" as a string
    stmt, err := db.Prepare("SELECT tanggal_pinjam FROM peminjaman WHERE peminjaman_id = ?")
    if err != nil {
        log.Fatal(err)
    }
    defer stmt.Close()

    var tanggalPinjamStr string
    // Execute the prepared statement with the actual value
    err = stmt.QueryRow(peminjamanID).Scan(&tanggalPinjamStr)
    if err == sql.ErrNoRows {
        fmt.Println("Peminjaman dengan ID tersebut tidak ditemukan.")
        return
    } else if err != nil {
        log.Fatal(err)
    }

    // Parse "tanggal_pinjamStr" into a time.Time value
    tanggalPinjam, err := time.Parse("2006-01-02", tanggalPinjamStr)
    if err != nil {
        log.Fatal(err)
    }

    // Mendapatkan tanggal saat ini
    tanggalSekarang := time.Now()

    // Menghitung selisih hari antara tanggal kembali yang diharapkan dan tanggal sekarang
    selisihHari := tanggalSekarang.Sub(tanggalPinjam).Hours() / 24

    // Menghitung denda berdasarkan aturan yang Anda tentukan
    denda := 0.0
    if selisihHari > 7 {
        denda = (selisihHari - 7) * 5000 // Denda Rp5.000 per hari
    }

    if denda > 0 {
        fmt.Printf("Buku dikembalikan terlambat! Anda memiliki denda sebesar Rp%.2f\n", denda)
    } else {
        fmt.Println("Buku telah dikembalikan tepat waktu.")
    }

    // Prepare the SQL statement to update tanggal_kembali and denda in detail_peminjaman
    updateStmt, err := db.Prepare("UPDATE detail_peminjaman SET tanggal_kembali = ?, denda = ? WHERE peminjaman_id = ?")
    if err != nil {
        log.Fatal(err)
    }
    defer updateStmt.Close()

    // Execute the update statement with the actual values
    _, err = updateStmt.Exec(tanggalSekarang, denda, peminjamanID)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Pengembalian buku berhasil dengan denda sebesar Rp%.2f\n", denda)
}


func tampilkanAnggotaTertinggi(db *sql.DB) {
    // Query untuk mengambil anggota dengan peminjaman tertinggi
    rows, err := db.Query("SELECT anggota_id, nama FROM anggota ORDER BY anggota_id DESC LIMIT 1")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // Menampilkan anggota dengan peminjaman tertinggi
    fmt.Println("===== Anggota dengan Peminjaman Tertinggi =====")
    for rows.Next() {
        var anggota Anggota
        err := rows.Scan(&anggota.ID, &anggota.Nama)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("ID Anggota: %d, Nama: %s\n", anggota.ID, anggota.Nama)
    }

    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }
}

func tampilkanDetailPeminjaman(db *sql.DB) {
    // Query untuk mengambil detail peminjaman dari database
    rows, err := db.Query("SELECT dp.detail_id, dp.peminjaman_id, dp.buku_id, dp.tanggal_kembali, dp.denda, b.judul FROM detail_peminjaman dp JOIN buku b ON dp.buku_id = b.buku_id")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // Menampilkan detail peminjaman
    fmt.Println("===== Detail Peminjaman =====")
    for rows.Next() {
        var detail DetailPeminjaman
        var tanggalKembaliStr string
        err := rows.Scan(&detail.ID, &detail.PeminjamanID, &detail.BukuID, &tanggalKembaliStr, &detail.Denda, &detail.Judul)
        if err != nil {
            log.Fatal(err)
        }
        // Konversi tanggalKembaliStr ke time.Time
        detail.TanggalKembali, err = time.Parse("2006-01-02", tanggalKembaliStr)
        if err != nil {
            log.Fatal(err)
        }

        fmt.Printf("ID Detail: %d, ID Peminjaman: %d, ID Buku: %d, Judul Buku: %s, Tanggal Kembali: %s, Denda: Rp%.2f\n", detail.ID, detail.PeminjamanID, detail.BukuID, detail.Judul, detail.TanggalKembali.Format("2006-01-02"), detail.Denda)
    }

    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }
}
