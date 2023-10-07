-- A. Retrieve all orders with their applied discounts:
SELECT
    O.order_id,
    O.table_number,
    O.order_date,
    O.status,
    E.first_name AS employee_first_name,
    E.last_name AS employee_last_name,
    E.position AS employee_position,
    GROUP_CONCAT(M.name SEPARATOR ', ') AS items,
    (SUM(OI.subtotal) - IFNULL(D.discount_amount, 0)) AS total_price_with_discount
FROM
    Orders AS O
INNER JOIN
    Employees AS E ON O.employee_id = E.employee_id
INNER JOIN
    OrderItems AS OI ON O.order_id = OI.order_id
INNER JOIN
    MenuItems AS M ON OI.item_id = M.item_id
LEFT JOIN
    Discounts AS D ON O.order_id = D.order_id
GROUP BY
    O.order_id
ORDER BY
    O.order_id;

-- B. Calculate the total revenue (including discounts) for a specific day

SELECT
    SUM(OI.subtotal) - IFNULL(SUM(D.discount_amount), 0) AS total_revenue_with_discount
FROM
    Orders AS O
INNER JOIN
    OrderItems AS OI ON O.order_id = OI.order_id
LEFT JOIN
    Discounts AS D ON O.order_id = D.order_id
WHERE
    O.order_date = '2023-08-04';
