CREATE TABLE users (
    user_id INT PRIMARY KEY,
    username VARCHAR(50) UNIQUE,
    email VARCHAR(100) UNIQUE
);

CREATE TABLE employees (
    emp_id INT PRIMARY KEY,
    emp_name VARCHAR(100) UNIQUE,
    age INT CHECK (age >= 18 AND age <= 56), -- Perbaikan constraint
    salary DECIMAL(10,2)
);

-- Perbaikan query INSERT INTO users
INSERT INTO users (user_id, username, email)
VALUES
(1, 'John_Doe', 'john.doe@example.com'), -- Perbaikan tanda kurung
(2, 'jane smith', 'jane.smith@example.com'), -- Perbaikan tanda kurung
(3, 'alice', 'alice@example.com');

-- Perbaikan query INSERT INTO employees
INSERT INTO employees (emp_id, emp_name, age, salary)
VALUES
(101, 'john Smith', 35, 55000.00), -- Perbaikan nilai salary dan tanda kurung
(102, 'john doe', 35, 48000.00), -- Perbaikan nilai salary dan tanda kurung
(103, 'bob', 45, 62000.00); -- Perbaikan nilai salary dan tanda kurung

UPDATE users
SET is_active = 0
WHERE user_id in (2,4);

UPDATE users
SET is_active = 1
WHERE user_id NOT IN (2,4);

CREATE VIEW active_users AS
SELECT user_id, username, email
FROM users
WHERE is_active = 1;

CREATE OR REPLACE VIEW top_earners AS
SELECT emp_id, emp_name, salary
FROM employees
WHERE salary > 50000;

CREATE TRIGGER update_last_modified
BEFORE UPDATE ON employees
FOR EACH ROW
SEt NEW.last_modified = NOW();

DELIMITER //
CREATE TRIGGER check_age_before_insert
BEFORE INSERT ON employees
FOR EACH ROW
BEGIN 
	IF NEW.age < 18 OR NEW.age > 65 THEN
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'Age must be between 18 and 65';
    END IF;
END;
//
DELIMITER ;

DELIMITER //
CREATE TRIGGER check_salary_before_insert
BEFORE INSERT ON employees
FOR EACH ROW
BEGIN 
	IF NEW.salary < 0 THEN
    	SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'Salary cannot be negative';
    END IF;
END;
//
DELIMITER ;

