package main

import (
	"fmt"
)

func Status() {
	fmt.Println("Panic Message [func invalid] :")
	fmt.Println("app stopped")
}

func invalid(val bool) (string, error) {
	defer Status()
	if val {
		panic("val is true")
	}

	fmt.Println("Apps Running")
	return "Valid", nil
}

func main() {
	_, _ = invalid(false)
}




// func main() {
// 	var number int
// 	var err error

// 	number, err = strconv.Atoi("123GH") // Perbaiki argumen strconv.Atoi
// 	if err != nil {
// 		fmt.Println("Not a number") // Perbaiki penggunaan errors.New
// 	} else {
// 		fmt.Println(number)
// 	}

// 	number, err = strconv.Atoi("123")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	} else {
// 		fmt.Println(number)
// 	}
// }
