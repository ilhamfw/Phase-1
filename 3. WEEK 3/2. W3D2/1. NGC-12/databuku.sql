-- Create the authors table
CREATE TABLE authors (
    author_id INT PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    country VARCHAR(50),
    CONSTRAINT unique_name UNIQUE (name) -- Constraint untuk memastikan kolom name pada tabel authors unik
);

-- Create the books table
CREATE TABLE books (
    book_id INT PRIMARY KEY,
    title VARCHAR(200) UNIQUE NOT NULL,
    author_id INT,
    publication_year INT,
    available_quantity INT CHECK (available_quantity >= 0), -- Constraint untuk memastikan kuantitas yang tersedia (available_quantity) tidak negatif
    FOREIGN KEY (author_id) REFERENCES authors(author_id)
);

-- Create the book_loans table
CREATE TABLE book_loans (
    loan_id INT PRIMARY KEY,
    book_id INT,
    borrower_name VARCHAR(100),
    loan_date DATE,
    return_date DATE CHECK (return_date >= loan_date), -- Constraint untuk memastikan bahwa kolom return_date pada tabel book_loans tidak boleh sebelum loan_date
    FOREIGN KEY (book_id) REFERENCES books(book_id)
);


-- Insert sample data into authors table
INSERT INTO authors (author_id, name, country) VALUES
(1, 'J.K. Rowling', 'United Kingdom'),
(2, 'George Orwell', 'United Kingdom'),
(3, 'Harper Lee', 'United States'),
(4, 'Jane Austen', 'United Kingdom');

-- Insert sample data into books table
INSERT INTO books (book_id, title, author_id, publication_year, available_quantity) VALUES
(101, 'Harry Potter and the Philosopher''s Stone', 1, 1997, 10),
(102, '1984', 2, 1949, 8),
(103, 'To Kill a Mockingbird', 3, 1960, 12),
(104, 'Pride and Prejudice', 4, 1813, 15);

-- Insert sample data into book_loans table
INSERT INTO book_loans (loan_id, book_id, borrower_name, loan_date, return_date) VALUES
(1001, 101, 'John Smith', '2023-09-15', '2023-09-30'),
(1002, 102, 'Jane Johnson', '2023-09-10', '2023-09-25'),
(1003, 103, 'Mary Davis', '2023-09-12', '2023-09-28'),
(1004, 101, 'Sarah Brown', '2023-09-20', '2023-10-05');


-- Add constraints to ensure data integrity


-- create indexes for faster retrilevel of books by 'author_id' and finding overdue book loans
CREATE INDEX idx_return_date ON book_loans(return_date, book_id);
CREATE INDEX idx_year_publish ON books(author_id, publication_year)

-- Menggunakan SUBSTRING untuk mengambil sebagian dari string
 SELECT name, SUBSTRING(name, 1, 5) AS short_name FROM authors

-- Menggunakan CONCAT untuk menggabungkan beberapa kolom menjadi 1 string
SELECT CONCAT(name, ' from ', country) AS Author_Info FROM authors;

-- Menggunakan DATE_ADD untuk menambahkan hari ke tanggal (Interval peminjaman)
SELECT book_id, DATE_ADD(loan_date, INTERVAL 7 DAY) AS due_date FROM book_loans;

-- Menggunakan DATEDIFF untuk menghitung selisih hari antara dua tanggal
SELECT book_id, DATEDIFF(return_date, loan_date) AS loan_duration FROM book_loans;

-- Menggunakan CASE untuk memberi label status buku berdasarkan ketersediaan
SELECT title,
	CASE
    	WHEN available_quantity > 0 THEN 'Available'
        ELSE 'Out Of Stock'
    END AS Book_Status
FROM books;



-- Menggunakan CASE untuk mengelompokkan buku berdasarkan tahun publikasi >= 2020 THEN 'Publication_Category = RECENT', < 2020 && >= 2010 THEN 'MODERATE', di luar tahun itu OLD
-- Menggunakan CASE untuk mengelompokkan buku berdasarkan tahun publikasi
SELECT
    book_id,
    title,
    publication_year,
    CASE
        WHEN publication_year >= 2020 THEN 'RECENT'
        WHEN publication_year >= 2010 AND publication_year < 2020 THEN 'MODERATE'
        ELSE 'OLD'
    END AS Publication_Category
FROM
    books;


