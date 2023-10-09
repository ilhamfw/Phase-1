-- a. List all books by 'Jane Smith':
SELECT book_title
FROM bookstore
WHERE author_name = 'Jane Smith';

-- b. Find the total sales (in terms of price) for each book type:
SELECT book_type, SUM(price) AS total_sales
FROM bookstore
GROUP BY book_type;


-- c. Identify customers who have ordered more than one book:
SELECT customer_name, customer_email
FROM bookstore
WHERE customer_name IN (
    SELECT customer_name
    FROM bookstore
    GROUP BY customer_name
    HAVING COUNT(DISTINCT book_title) > 1
);


-- d. Display the author who has the highest earnings from their bookstore:
SELECT author_name, SUM(price) AS total_earnings
FROM bookstore
GROUP BY author_name
ORDER BY total_earnings DESC
LIMIT 1;

