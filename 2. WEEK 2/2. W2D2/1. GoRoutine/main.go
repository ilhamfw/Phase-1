package main

import (
	"fmt"
	"time"
)

func sayHello(){
	for i := 0; i < 10; i++{
		time.Sleep(100* time.Millisecond)
		fmt.Println("Hello")
	}
}

func main(){
	go sayHello()
	fmt.Println("Goroutine telah dijalankan!")
	time.Sleep(1 * time.Second)
}