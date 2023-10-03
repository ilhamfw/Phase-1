-- Tabel Student
CREATE TABLE Student (
    student_id INT PRIMARY KEY,
    student_name VARCHAR(100)
);

-- Tabel Professor
CREATE TABLE Professor (
    professor_id INT PRIMARY KEY,
    professor_name VARCHAR(100)
);

-- Tabel Course
CREATE TABLE Course (
    course_id INT PRIMARY KEY,
    course_title VARCHAR(200),
    max_capacity INT
);

-- Tabel Enrollment (Ternary Relationship)
CREATE TABLE Enrollment (
    enrollment_id INT PRIMARY KEY,
    student_id INT,
    course_id INT,
    enrollment_date DATE,
    FOREIGN KEY (student_id) REFERENCES Student(student_id),
    FOREIGN KEY (course_id) REFERENCES Course(course_id)
);

-- Tabel TeachingAssignment (Ternary Relationship)
CREATE TABLE TeachingAssignment (
    assignment_id INT PRIMARY KEY,
    professor_id INT,
    course_id INT,
    start_date DATE,
    FOREIGN KEY (professor_id) REFERENCES Professor(professor_id),
    FOREIGN KEY (course_id) REFERENCES Course(course_id)
);

-- Masukkan data mahasiswa
INSERT INTO Student (student_id, student_name)
VALUES
    (1, 'Mahasiswa 1'),
    (2, 'Mahasiswa 2'),
    (3, 'Mahasiswa 3');

-- Masukkan data dosen
INSERT INTO Professor (professor_id, professor_name)
VALUES
    (1, 'Dosen 1'),
    (2, 'Dosen 2'),
    (3, 'Dosen 3'),
    (4, 'Dosen 4'),
    (5, 'Dosen 5');

-- Masukkan data kursus
INSERT INTO Course (course_id, course_title, max_capacity)
VALUES
    (101, 'Kursus 1', 30),
    (102, 'Kursus 2', 25),
    (103, 'Kursus 3', 40),
    (104, 'Kursus 4', 20),
    (105, 'Kursus 5', 35);

-- Daftar mahasiswa yang terdaftar dalam kursus tertentu (misalnya, kursus dengan course_id = 101)
SELECT s.student_name
FROM Student s
JOIN Enrollment e ON s.student_id = e.student_id
WHERE e.course_id = 101;

-- Daftar kursus yang diajar oleh seorang profesor tertentu (misalnya, profesor dengan professor_id = 1)
SELECT c.course_title
FROM Course c
JOIN TeachingAssignment ta ON c.course_id = ta.course_id
WHERE ta.professor_id = 1;

-- Daftar profesor yang mengajar kursus tertentu (misalnya, kursus dengan course_id = 101)
SELECT p.professor_name
FROM Professor p
JOIN TeachingAssignment ta ON p.professor_id = ta.professor_id
WHERE ta.course_id = 101;
