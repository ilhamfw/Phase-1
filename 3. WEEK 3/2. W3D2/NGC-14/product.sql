CREATE TABLE Products (
    ProductID INT PRIMARY KEY,
    ProductName VARCHAR(100),
    ProductCode VARCHAR(20),
    Price DECIMAL(10, 2)
);

INSERT INTO Products (ProductID, ProductName, ProductCode, Price)
VALUES
   Laptop', 'LT001', 799.99),
    (2, 'Smartphone', (1, ' 'SP002', 499.99),
    (3, 'Tablet', 'TB003', 299.99),
    (4, 'Headphones', 'HD004', 149.99),
    (5, 'Monitor', 'MN005', 399.99);

CREATE TABLE ProductSales (
    SaleID INT PRIMARY KEY,
    ProductID INT,
    SaleDate DATE,
    QuantitySold INT,
    FOREIGN KEY (ProductID) REFERENCES Products(ProductID)
);

INSERT INTO ProductSales (SaleID, ProductID, SaleDate, QuantitySold)
VALUES 
    (1, 1, '2023-08-01', 10),
    (2, 1, '2023-08-02', 5),
    (3, 2, '2023-08-02', 8),
    (4, 3, '2023-08-03', 12),
    (5, 3, '2023-08-03', 6),
    (6, 4, '2023-08-04', 20),
    (7, 5, '2023-08-04', 15);

-- Challenge 1: Product Count by Price Category
SELECT
    CASE
        WHEN Price < 300 THEN 'Low Price'
        WHEN Price >= 300 AND Price <= 600 THEN 'Medium Price'
        WHEN Price > 600 THEN 'High Price'
    END AS Price_Category,
    COUNT(*) AS Product_Count
FROM Products
GROUP BY Price_Category;

-- Challenge 2: Top N Products by Sales
SELECT
    p.ProductName,
    SUM(ps.QuantitySold) AS TotalQuantitySold
FROM Products p
INNER JOIN ProductSales ps ON p.ProductID = ps.ProductID
GROUP BY p.ProductName
ORDER BY TotalQuantitySold DESC
LIMIT 3; -- Ganti nilai 3 dengan N sesuai kebutuhan

-- Challenge 3: Product Sales Growth
SELECT
    p.ProductName,
    SUM(CASE WHEN ps.SaleDate >= DATE_SUB(CURDATE(), INTERVAL 1 MONTH) THEN ps.QuantitySold ELSE 0 END) AS CurrentPeriodSales,
    SUM(CASE WHEN ps.SaleDate < DATE_SUB(CURDATE(), INTERVAL 1 MONTH) AND ps.SaleDate >= DATE_SUB(CURDATE(), INTERVAL 2 MONTH) THEN ps.QuantitySold ELSE 0 END) AS PreviousPeriodSales,
    CASE
        WHEN SUM(CASE WHEN ps.SaleDate < DATE_SUB(CURDATE(), INTERVAL 1 MONTH) AND ps.SaleDate >= DATE_SUB(CURDATE(), INTERVAL 2 MONTH) THEN ps.QuantitySold ELSE 0 END) = 0 THEN 100.00 -- Handle division by zero
        ELSE ROUND(((SUM(CASE WHEN ps.SaleDate >= DATE_SUB(CURDATE(), INTERVAL 1 MONTH) THEN ps.QuantitySold ELSE 0 END) - SUM(CASE WHEN ps.SaleDate < DATE_SUB(CURDATE(), INTERVAL 1 MONTH) AND ps.SaleDate >= DATE_SUB(CURDATE(), INTERVAL 2 MONTH) THEN ps.QuantitySold ELSE 0 END)) / SUM(CASE WHEN ps.SaleDate < DATE_SUB(CURDATE(), INTERVAL 1 MONTH) AND ps.SaleDate >= DATE_SUB(CURDATE(), INTERVAL 2 MONTH) THEN ps.QuantitySold ELSE 0 END)) * 100, 2)
    END AS SalesGrowthPercentage
FROM Products p
LEFT JOIN ProductSales ps ON p.ProductID = ps.ProductID
GROUP BY p.ProductName;

-- Challenge 4: Average Price Difference
SELECT
    p1.ProductID,
    p1.ProductName,
    p1.Price,
    AVG(p1.Price - COALESCE(p2.Price, p1.Price)) AS AvgPriceDifference
FROM
    Products p1
LEFT JOIN
    Products p2
ON
    p1.ProductID = p2.ProductID + 1
GROUP BY
    p1.ProductID, p1.ProductName, p1.Price
HAVING
    p1.ProductID > 1;


-- Challenge 5: Products without Price
SELECT
    ProductName,
    ProductCode
FROM
    Products
WHERE
    Price IS NOT NULL;


    