package main

import "fmt"

func main() {
	
	var a int
	var b int
	
	fmt.Println("Введите число a")
	fmt.Scanln(&a)

	fmt.Println("Введите число b")
	fmt.Scanln(&b)

	fmt.Println("Площадь прямоугольника равна", (a * b))
	
}