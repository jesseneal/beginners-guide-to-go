package main

import "fmt"

func HelloWorld(name string, age int) {
	fmt.Println("Hello", name)
	fmt.Println("Age", age)
}

func AddTotal(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	HelloWorld("Jesse", 39)
	total, negativeTotal := AddTotal(1, 2)
	fmt.Println(total, negativeTotal)
}
