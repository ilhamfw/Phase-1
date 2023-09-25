1. for loop
2. for range loop

// for loop - 1. Ketika kita sudah mengetahu iterasi secara spesifik


for i := 0; i < 5; i++{
    fmt.Println(i)
}

// for loop - 2. Ketika kita mau ada increment atau decrement
for i := 0; i < 100; i += 10{
    fmt.Println(i)
}

// for range loop 1. - Iterasi atas sbuah collection / struktur Data

arr := []int{1,2,3,4,5}
for index, value := range arr {
	fmt.Printf("Index: %d, Value: %d\n, index, value")
}

// for range loop 2. - Bekerja denga strings 

str := "Hello, World"
for index, char := range str {
	fmt.Printf("Index: %d, Char: %c\n, index, char")
}

// for range loop dengan library bufio
scanner := bufio.NewScanner(file)
for scanner.Scan(){
	fmt.Println(Scanner.Text())
}

