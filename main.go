package main

import "fmt"

func b(a int) {
	a = a + 1
	fmt.Println(a)
}
func main() {
	a := 0
	b(a)
	fmt.Println(a)
}
