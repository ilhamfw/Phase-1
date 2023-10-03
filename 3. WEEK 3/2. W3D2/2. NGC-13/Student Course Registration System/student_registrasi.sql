-- Tabel siswa (students)
CREATE TABLE students (
    student_id INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE,
    major VARCHAR(50),
    year_of_study INT
);

-- Tabel kursus (courses)
CREATE TABLE courses (
    course_id INT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    instructor VARCHAR(255),
    schedule VARCHAR(100),
    credit_hours INT
);

-- Tabel kursus_siswa (student_courses)
CREATE TABLE student_courses (
    student_id INT,
    course_id INT,
    preferred_schedule VARCHAR(100),
    PRIMARY KEY (student_id, course_id),
    FOREIGN KEY (student_id) REFERENCES students(student_id),
    FOREIGN KEY (course_id) REFERENCES courses(course_id)
);

-- Masukkan data siswa
INSERT INTO students (student_id, name, email, major, year_of_study)
VALUES
    (1, 'Siswa 1', 'siswa1@example.com', 'Teknik Informatika', 2),
    (2, 'Siswa 2', 'siswa2@example.com', 'Ekonomi', 3),
    (3, 'Siswa 3', 'siswa3@example.com', 'Kimia', 1),
    (4, 'Siswa 4', 'siswa4@example.com', 'Fisika', 4),
    (5, 'Siswa 5', 'siswa5@example.com', 'Biologi', 2);

-- Masukkan data mata kuliah
INSERT INTO courses (course_id, title, instructor, schedule, credit_hours)
VALUES
    (101, 'Matematika Dasar', 'Prof. A', 'Senin 09:00-11:00', 3),
    (102, 'Ekonomi Mikro', 'Prof. B', 'Rabu 14:00-16:00', 3),
    (103, 'Kimia Organik', 'Prof. C', 'Selasa 10:00-12:00', 4),
    (104, 'Fisika Modern', 'Prof. D', 'Kamis 13:00-15:00', 4),
    (105, 'Biologi Sel', 'Prof. E', 'Jumat 08:00-10:00', 3);

-- Daftarkan siswa untuk beberapa mata kuliah dengan jadwal pilihan
INSERT INTO student_courses (student_id, course_id, preferred_schedule)
VALUES
    (1, 101, 'Senin 09:00-11:00'),
    (1, 102, 'Rabu 14:00-16:00'),
    (2, 103, 'Selasa 10:00-12:00'),
    (3, 104, 'Kamis 13:00-15:00'),
    (4, 105, 'Jumat 08:00-10:00');

-- Ambil daftar semua siswa yang terdaftar dalam kursus tertentu (contoh kursus dengan course_id = 101)
SELECT s.name
FROM students s
JOIN student_courses sc ON s.student_id = sc.student_id
WHERE sc.course_id = 101;

-- Temukan semua kursus yang telah didaftarkan oleh siswa tertentu (contoh siswa dengan student_id = 1)
SELECT c.title
FROM courses c
JOIN student_courses sc ON c.course_id = sc.course_id
WHERE sc.student_id = 1;

-- Dapatkan jadwal pilihan siswa untuk kursus tertentu (contoh siswa dengan student_id = 1 dan course_id = 101)
SELECT sc.preferred_schedule
FROM student_courses sc
WHERE sc.student_id = 1 AND sc.course_id = 101;
