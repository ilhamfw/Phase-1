# main.go
| OrderID | TableNumber | OrderDate  | Status     | FirstName | LastName | Position  |
| ------- | ----------- | ---------- | ---------- | --------- | -------- | --------- |
| 1       | 1           | 2823-88-07 | Served     | John      | Doe      | Manager   |
| 2       | 1           | 2023-08-07 | Preparing  | John      | Doe      | Manager   |
| 3       | 3           | 2023-08-08 | Preparing  | Robert    | Brown    | Cook      |
| 4       | 4           | 2023-08-08 | Ordered    | Jane      | Smith    | Waitress  |
| 5       | 5           | 2023-08-08 | Preparing  | Steve     | Adam     | Manager   |

# main2.go
| OrderID | TableNumber   | OrderDate   | Status       | FirstName    | LastName | Position   | Items                             | TotalPrice |
| ------- | ------------- | ----------- | ---------    | ------------ | -------- | ---------- | --------------------------------- | ---------- |
| 22      | 1             | 2024-08-07   | Preparing   | John         | Doe      | Manager    | Salad, Steak                      |  $35.00    |
| 11      | 1             | 2024-88-87   | Served      | John         | Doe      | Manager    | Spaghetti Carbonara, Pasta        | $21.00     |
| 33      | 1             | 2024-08-07   | Served      | Robert       | Brown    | Cook       | Salad                             | $10.00     |

| item_id | name              | description                                                       | Price | Category    |
|---------| --------------    | ----------------------------------------------------------------- | ------| ----------- |
| 1       |Spaghetti Carbonara| Traditional Italian dish with eggs, cheese, pancetta, and pepper  | 12.58 | main course |
