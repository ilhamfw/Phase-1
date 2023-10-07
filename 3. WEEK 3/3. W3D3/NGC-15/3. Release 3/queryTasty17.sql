-- 1. Mengambil detail pesanan dan karyawan yang sesuai
SELECT
    Orders.order_id AS OrderID,  -- ID pesanan
    Orders.order_date AS OrderDate,  -- Tanggal pesanan
    Employees.first_name AS EmployeeFirstName,  -- Nama depan karyawan
    Employees.last_name AS EmployeeLastName  -- Nama belakang karyawan
FROM Orders
JOIN Employees ON Orders.employee_id = Employees.employee_id;

-- 2. Mengambil daftar pesanan, item, dan total biaya
SELECT
    Orders.order_id AS OrderID,  -- ID pesanan
    MenuItems.description AS ItemDescription,  -- Deskripsi item
    OrderItems.quantity AS Quantity,  -- Jumlah item dalam pesanan
    MenuItems.price AS ItemPrice,  -- Harga per item
    (OrderItems.quantity * MenuItems.price) AS TotalCost  -- Total biaya item dalam pesanan
FROM Orders
JOIN OrderItems ON Orders.order_id = OrderItems.order_id
JOIN MenuItems ON OrderItems.item_id = MenuItems.item_id;
