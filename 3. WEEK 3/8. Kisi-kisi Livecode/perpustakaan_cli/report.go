package main

import (
    "database/sql"
    "fmt"
    "log"
)



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
        var judulBuku string
        err := rows.Scan(&detail.ID, &detail.PeminjamanID, &detail.BukuID, &detail.TanggalKembali, &detail.Denda, &judulBuku)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("ID Detail: %d, ID Peminjaman: %d, ID Buku: %d, Judul Buku: %s, Tanggal Kembali: %s, Denda: Rp%.2f\n", detail.ID, detail.PeminjamanID, detail.BukuID, judulBuku, detail.TanggalKembali.Format("2006-01-02"), detail.Denda)
    }

    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }
}
