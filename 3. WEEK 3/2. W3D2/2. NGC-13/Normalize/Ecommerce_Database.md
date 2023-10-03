Diberikan table belanja online sebagai berikut :

| Order ID | Customer Name | Customer Address | Product ID | Product Name | Category     | Price  |
|----------|---------------|------------------|------------|--------------|--------------|--------|
| 1001     | John Smith    | 123 Main St      | P001       | T-shirts     | Clothing     | 20.99  |
| 1002     | Mary Johnson  | 456 Elm St       | P002       | Smartphone   | Electronics  | 499.99 |
| 1003     | John Smith    | 123 Main St      | P003       | Headphone    | Electronics  | 49.99  |
| 1004     | Sarah Brown   | 789 Oak Ave      | P004       | Book         | Books        | 12.49  |
| 1005     | Mary Johnson  | 456 Elm St       | P005       | Camera       | Electronics  | 299.39 |



Tabel Customer:
```
Customer Name | Customer Address
-------------------------------
John Smith    | 123 Main St
Mary Johnson  | 456 Elm St
Sarah Brown   | 789 Oak Ave
```

Tabel Product:
```
Product ID | Product Name | Category     | Price
----------------------------------------------
P001       | T-shirts     | Clothing     | 20.99
P002       | Smartphone   | Electronics  | 499.99
P003       | Headphone    | Electronics  | 49.99
P004       | Book         | Books        | 12.49
P005       | Camera       | Electronics  | 299.39
```

Tabel Order:
```
Order ID | Customer Name | Product ID
-------------------------------------
1001     | John Smith    | P001
1002     | Mary Johnson  | P002
1003     | John Smith    | P003
1004     | Sarah Brown   | P004
1005     | Mary Johnson  | P005
```

**Langkah 3: Normalisasi ke 3NF (Bentuk Normal Ketiga)**
Dalam bentuk normal ketiga (3NF), tidak ada ketergantungan transitif antara atribut-atribut non-kunci. Dalam tabel "Product," "Category" sepenuhnya bergantung pada "Product ID." Tidak ada ketergantungan transitif lain dalam tabel ini.

Dengan melakukan normalisasi ke 3NF, kita telah membagi tabel awal menjadi tiga tabel terpisah dengan dependensi yang jelas antara atribut. Ini meningkatkan integritas dan efisiensi data dalam basis data.

Selanjutnya, Anda dapat membuat file Markdown di Visual Studio Code dengan nama "belanja_online.md" dan menambahkan tabel yang sudah dinormalisasi ke dalamnya sesuai dengan format Markdown. Berikut contoh format tabel dalam file Markdown:

```markdown
| Order ID | Customer Name | Product ID |
|----------|---------------|------------|
| 1001     | John Smith    | P001       |
| 1002     | Mary Johnson  | P002       |
| 1003     | John Smith    | P003       |
| 1004     | Sarah Brown   | P004       |
| 1005     | Mary Johnson  | P005       |

**Tabel Customer:**

| Customer Name | Customer Address |
|---------------|------------------|
| John Smith    | 123 Main St      |
| Mary Johnson  | 456 Elm St       |
| Sarah Brown   | 789 Oak Ave      |

**Tabel Product:**

| Product ID | Product Name | Category     | Price  |
|------------|--------------|--------------|--------|
| P001       | T-shirts     | Clothing     | 20.99  |
| P002       | Smartphone   | Electronics  | 499.99 |
| P003       | Headphone    | Electronics  | 49.99  |
| P004       | Book         | Books        | 12.49  |
| P005       | Camera       | Electronics  | 299.39 |
```

Dengan cara ini, Anda dapat menyimpan data yang sudah dinormalisasi dalam format Markdown di Visual Studio Code.