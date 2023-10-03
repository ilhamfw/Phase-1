Diberikan table employee management data sebagai berikut :
| Employee ID | Name  | Department | Supervisor ID | Supervisor Name | Salary   |
|-------------|-------|------------|---------------|-----------------|----------|
| 101         | John  | HR         | 201           | Mary            | 50000.0  |
| 102         | Jane  | Finance    | 202           | Sarah           | 55000.0  |
| 201         | Mary  | HR         | 203           | David           | 60000.0  |
| 202         | Sarah | Finance    | 203           | David           | 65000.0  |
| 203         | David | CEO        | NULL          | NULL            | 100000.0 |


Tabel Employee:
```
Employee ID | Name    | Department | Salary
------------------------------------------
101         | John    | HR         | 50000.0
102         | Jane    | Finance    | 55000.0
201         | Mary    | HR         | 60000.0
202         | Sarah   | Finance    | 65000.0
203         | David   | CEO        | 100000.0
```

Tabel Supervisor:
```
Supervisor ID | Supervisor Name
-------------------------------
201           | Mary
202           | Sarah
203           | David
```


```
Employee ID | Name    | Department | Salary
------------------------------------------
101         | John    | HR         | [Dilindungi]
102         | Jane    | Finance    | [Dilindungi]
201         | Mary    | HR         | [Dilindungi]
202         | Sarah   | Finance    | [Dilindungi]
203         | David   | CEO        | [Dilindungi]
```

