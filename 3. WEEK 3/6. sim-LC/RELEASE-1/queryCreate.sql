-- Create Fields table
CREATE TABLE Fields (
    FieldID INT PRIMARY KEY AUTO_INCREMENT,
    FieldName VARCHAR(255),
    FieldSize VARCHAR(10),
    Location VARCHAR(255),
    HourlyRate DECIMAL(10, 2)
);

-- Create Customers table
CREATE TABLE Customers (
    CustomerID INT PRIMARY KEY AUTO_INCREMENT,
    FirstName VARCHAR(255),
    LastName VARCHAR(255),
    Email VARCHAR(255),
    PhoneNumber VARCHAR(20)
);

-- Create Bookings table
CREATE TABLE Bookings (
    BookingID INT PRIMARY KEY AUTO_INCREMENT,
    CustomerID INT,
    FieldID INT,
    BookingDate DATE,
    StartTime TIME,
    EndTime TIME,
    TotalAmount DECIMAL(10, 2),
    FOREIGN KEY (CustomerID) REFERENCES Customers(CustomerID),
    FOREIGN KEY (FieldID) REFERENCES Fields(FieldID)
);

-- Create Payments table
CREATE TABLE Payments (
    PaymentID INT PRIMARY KEY AUTO_INCREMENT,
    BookingID INT,
    PaymentDate DATE,
    PaymentAmount DECIMAL(10, 2),
    PaymentMethod VARCHAR(255),
    FOREIGN KEY (BookingID) REFERENCES Bookings(BookingID)
);
