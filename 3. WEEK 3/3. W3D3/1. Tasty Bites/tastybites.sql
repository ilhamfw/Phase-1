-- Tabel untuk menyimpan informasi karyawan
CREATE TABLE Employees (
    employee_id INT PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    position VARCHAR(255)
);

-- Tabel untuk menyimpan daftar menu makanan
CREATE TABLE MenuItems (
    item_id INT PRIMARY KEY,
    name VARCHAR(255),
    description TEXT,
    price DECIMAL(10, 2),
    category VARCHAR(255)
);

-- Tabel untuk melacak pesanan yang ditempatkan
CREATE TABLE Orders (
    order_id INT PRIMARY KEY,
    table_number INT,
    employee_id INT,
    order_date DATE,
    status VARCHAR(255),
    FOREIGN KEY (employee_id) REFERENCES Employees(employee_id)
);

-- Tabel untuk melacak item dalam setiap pesanan
CREATE TABLE OrderItems (
    order_item_id INT PRIMARY KEY,
    order_id INT,
    item_id INT,
    quantity INT,
    subtotal DECIMAL(10, 2),
    FOREIGN KEY (order_id) REFERENCES Orders(order_id),
    FOREIGN KEY (item_id) REFERENCES MenuItems(item_id)
);

-- Tabel untuk melacak pembayaran untuk pesanan
CREATE TABLE Payments (
    payment_id INT PRIMARY KEY,
    order_id INT,
    payment_date DATE,
    payment_method VARCHAR(255),
    total_amount DECIMAL(10, 2),
    FOREIGN KEY (order_id) REFERENCES Orders(order_id)
);

-- Tabel untuk Discounts
CREATE TABLE Discounts (
    discount_id INT AUTO_INCREMENT PRIMARY KEY,
    order_id INT,
    discount_amount DECIMAL(10, 2),
    FOREIGN KEY (order_id) REFERENCES Orders(order_id)
);



-- Memasukkan catatan karyawan
INSERT INTO Employees (employee_id, first_name, last_name, position) VALUES (1, 'John', 'Doe', 'Waiter');

-- Memasukkan catatan menu item
INSERT INTO MenuItems (item_id, name, description, price, category) VALUES (1, 'Steak', 'Grilled sirloin steak', 25.99, 'Main Course');

-- Memasukkan catatan pesanan
INSERT INTO Orders (order_id, table_number, employee_id, order_date, status) VALUES (1, 5, 1, '2023-08-04', 'Pending');

-- Memasukkan catatan item pesanan
INSERT INTO OrderItems (order_item_id, order_id, item_id, quantity, subtotal) VALUES (1, 1, 1, 2, 51.98);

-- Memasukkan catatan pembayaran
INSERT INTO Payments (payment_id, order_id, payment_date, payment_method, total_amount) VALUES (1, 1, '2023-08-04', 'Credit Card', 51.98);

