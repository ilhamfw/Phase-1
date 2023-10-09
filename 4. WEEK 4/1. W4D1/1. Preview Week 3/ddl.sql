
CREATE TABLE bookstore (
order_id INT,
customer_name VARCHAR(255),
customer_email VARCHAR(255),
book_title VARCHAR(255),
book_type VARCHAR(255), /* E-Book or Physical */
author_name VARCHAR(255),
author_email VARCHAR(255),
order_date DATE,
price DECIMAL(10, 2)
);


INSERT INTO bookstore VALUES
(1, 'John Doe', 'john.doe@email.com', 'Database Design', 'Physical', 'Jane Smith', 'jane.smith@email.com', '2023-08-10', 25.99), 
(2, 'John Doe', 'john.doe@email.com', 'Web Development', 'E-Book', 'Tom Brown', 'tom.brown@email.com', '2023-08-11', 19.99), 
(3, 'Alice Bob', 'alice.bob@email.com', 'Database Design', 'E-Book', 'Jane Smith', 'jane.smith@email.com', '2023-08-12', 28.99);

-- Tabel Customers
CREATE TABLE Customers (
    customer_id INT AUTO_INCREMENT PRIMARY KEY,
    customer_name VARCHAR(255) NOT NULL,
    customer_email VARCHAR(255) NOT NULL UNIQUE
);

-- Tabel Books
CREATE TABLE Books (
    book_id INT AUTO_INCREMENT PRIMARY KEY,
    book_title VARCHAR(255) NOT NULL,
    book_type VARCHAR(50) NOT NULL,
    author_name VARCHAR(255) NOT NULL,
    author_email VARCHAR(255) NOT NULL,
    FOREIGN KEY (author_email) REFERENCES Authors(author_email)
);

-- Tabel Authors
CREATE TABLE Authors (
    author_id INT AUTO_INCREMENT PRIMARY KEY,
    author_name VARCHAR(255) NOT NULL,
    author_email VARCHAR(255) NOT NULL UNIQUE
);

-- Tabel Orders
CREATE TABLE Orders (
    order_id INT AUTO_INCREMENT PRIMARY KEY,
    customer_id INT NOT NULL,
    order_date DATE NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    FOREIGN KEY (customer_id) REFERENCES Customers(customer_id),
    FOREIGN KEY (book_title) REFERENCES Books(book_title)
);

-- Insert data into Authors table
INSERT INTO Authors (author_name, author_email)
VALUES
    ('Jane Smith', 'jane.smith@email.com'),
    ('Tom Brown', 'tom.brown@email.com');

-- Insert data into Customers table
INSERT INTO Customers (customer_name, customer_email)
VALUES
    ('John Doe', 'john.doe@email.com'),
    ('Alice Bob', 'alice.bob@email.com');

-- Insert data into Books table
INSERT INTO Books (book_title, book_type, author_name, author_email)
VALUES
    ('Database Design', 'Physical', 'Jane Smith', 'jane.smith@email.com'),
    ('Web Development', 'E-Book', 'Tom Brown', 'tom.brown@email.com');

-- Insert data into Orders table
INSERT INTO Orders (customer_id, order_date, price)
VALUES
    (1, '2023-08-10', 25.99),
    (1, '2023-08-11', 19.99),
    (2, '2023-08-12', 28.99);
