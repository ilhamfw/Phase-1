//Dengan Konkurensi
package main

import (
    "fmt"
    "sync"
)

func factorial(n int, wg *sync.WaitGroup, resultChan chan<- int) {
    defer wg.Done()
    
    result := 1
    for i := 2; i <= n; i++ {
        result *= i
    }

    resultChan <- result
    
    fmt.Printf("Faktorial dari %d adalah %d\n", n, result)
}

func main() {
    numbers := []int{5, 7, 10, 3, 8}
    
    var wg sync.WaitGroup
    resultChan := make(chan int, len(numbers))

    for _, n := range numbers {
        wg.Add(1)
        go factorial(n, &wg, resultChan)
    }

    go func(){
        wg.Wait()
        close(resultChan)
    }()

    for r :=range resultChan{
        fmt.Println(r)
    }
    
    
}

//Tanpa Konkurensi
// package main

// import (
// 	"fmt"
// )

// func factorial(n int) int {
// 	if n == 0 || n == 1 {
// 		return 1
// 	}
// 	result := 1
// 	for i := 2; i <= n; i++ {
// 		result *= i
// 	}
// 	return result
// }

// func main() {
// 	numbers := []int{5, 7, 10}

// 	for _, num := range numbers {
// 		fact := factorial(num)
// 		fmt.Printf("Faktorial dari %d adalah %d\n", num, fact)
// 	}
// }

