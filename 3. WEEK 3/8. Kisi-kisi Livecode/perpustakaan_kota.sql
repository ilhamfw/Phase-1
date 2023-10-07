-- 1A. Membuat database perpustakaan_kota
CREATE DATABASE perpustakaan_kota;

-- 1B. Membuat tabel buku
USE perpustakaan_kota;
CREATE TABLE buku (
    buku_id INT AUTO_INCREMENT PRIMARY KEY,
    judul VARCHAR(255) NOT NULL,
    penulis VARCHAR(255) NOT NULL
);

-- Membuat tabel anggota
CREATE TABLE anggota (
    anggota_id INT AUTO_INCREMENT PRIMARY KEY,
    nama VARCHAR(255) NOT NULL
);

-- Membuat tabel peminjaman
CREATE TABLE peminjaman (
    peminjaman_id INT AUTO_INCREMENT PRIMARY KEY,
    anggota_id INT,
    tanggal_pinjam DATE,
    FOREIGN KEY (anggota_id) REFERENCES anggota(anggota_id)
);

-- Membuat tabel detail_peminjaman
CREATE TABLE detail_peminjaman (
    detail_id INT AUTO_INCREMENT PRIMARY KEY,
    peminjaman_id INT,
    buku_id INT,
    tanggal_kembali DATE,
    denda DECIMAL(10, 2),
    FOREIGN KEY (peminjaman_id) REFERENCES peminjaman(peminjaman_id),
    FOREIGN KEY (buku_id) REFERENCES buku(buku_id)
);

UPDATE detail_peminjaman
SET tanggal_kembali = CURRENT_DATE()
WHERE peminjaman_id = :peminjaman_id;


-- 1C. Memasukkan data sampel ke dalam tabel-tabel
-- Tabel buku
INSERT INTO buku (judul, penulis) VALUES
    ('Ayat-ayat Cinta', 'Habiburrahman El-Shirazy'),
    ('Hujan', 'Tere Liye'),
    ('Surga Yang Tak Dirindukan', 'Asma Nadia');

-- Tabel anggota
INSERT INTO anggota (nama) VALUES
    ('Ahmad'),
    ('Peter Parker'),
    ('John Wick');

-- Tabel peminjaman
INSERT INTO peminjaman (anggota_id, tanggal_pinjam) VALUES
    (1, '2023-10-01'),
    (2, '2023-10-02'),
    (3, '2023-10-03');

-- Tabel detail_peminjaman
INSERT INTO detail_peminjaman (peminjaman_id, buku_id, tanggal_kembali, denda) VALUES
    (1, 1, '2023-10-10', 0.00),
    (2, 2, '2023-10-11', 5.00),
    (3, 3, '2023-10-12', 2.50);
