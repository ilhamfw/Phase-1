ERD Title: Mini Soccer Report System

Entities and their Attributes:

`Entity: Field`
FieldID (INT) (PK, AI)
FieldName (VARCHAR)
FieldSize (VARCHAR)
Location (VARCHAR)
HourlyRate (FLOAT)

`Entity: CustomerTable`
CustomerID (INT) (PK, AI)
FirstName (VARCHAR)
LastName (VARCHAR)
Email (VARCHAR)
PhoneNumber (VARCHAR)

`Entity: BookingTable`
BookingID (INT) (PK, AI)
CustomerID (INT) (FK)
FieldID (INT) (FK)
BookingDate (DATE)
StartTime (TIMESTAMP)
EndTime (TIMESTAMP)
TotalAmount (FLOAT)

`Entity: PaymentTable`

PaymentID (INT) (PK, AI)
BookingID (INT) (FK)
PaymentDate (DATE)
PaymentAmount (FLOAT)
PaymentMethod (VARCHAR)

2. Relationship:

Berikut adalah relasi antar tabel berdasarkan tabel yang telah diberikan:

`Field to BookingTable (Field to Booking)`

Type: One to Many

Description: Satu lapangan (Field) dapat memiliki banyak pemesanan (Booking), tetapi setiap pemesanan hanya terkait dengan satu lapangan.

`CustomerTable to BookingTable (CustomerTable to Booking)`

Type: One to Many

Description: Satu pelanggan (Customer) dapat melakukan banyak pemesanan (Booking), tetapi setiap pemesanan hanya terkait dengan satu pelanggan.

`BookingTable to PaymentTable (Booking to Payment)`

Type: One to Many

Description: Satu pemesanan (Booking) dapat memiliki banyak pembayaran (Payment), tetapi setiap pembayaran hanya terkait dengan satu pemesanan. 