package main

import  "fmt"


func main() {

	var a int

	fmt.Println("Введите трехзначное число")
	fmt.Scanln(&a)

	hundreds := a/100%10 
	dozens := a/10%10
	units := a%10

	fmt.Println("В данном числе сотен", hundreds, "десятков", dozens, "единиц", units)

//	fmt.Println("В данном числе сотен", (a/100%10), "десятков", (a/10%10), "единиц", (a%10))
//	или сразу вот так	
}