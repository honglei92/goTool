package main

import "fmt"

func hello1() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello groutine")
	}
}
func goroutine() {
	go hello1()
	fmt.Println("stop goroutine")
}
func main() {
	// fmt.Println("hello ")
	goroutine()
	fmt.Println("hello ")
}
