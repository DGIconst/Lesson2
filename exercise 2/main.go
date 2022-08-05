package main

import (
    "fmt"
    "math"
)
 
func main() {
	var a float64

	fmt.Println("Введите площадь круга")
	fmt.Scanln(&a)

    fmt.Println("Диаметр круга равен", 2*math.Sqrt(a/math.Pi))
	fmt.Println("Длинна окружности равна", 2*math.Pi*math.Sqrt(a/math.Pi)) /* вычисление из радиуса*/
	fmt.Println("Длинна окружности равна", math.Sqrt(4*math.Pi*a)) /* вычисление прямо из площади*/

//	так понятнее но неинтересно
//	var a float64
//	
//	fmt.Println("Введите площадь круга")
//	fmt.Scanln(&a)
//
//	var c float64 =  a/math.Pi
//  var r float64 =  math.Sqrt(c) /* радиус */
//
//  fmt.Println("Диаметр круга равен", 2*r)
//	fmt.Println("Длинна окружности равна", 2*math.Pi*r) /* вычисление из радиуса*/
//	fmt.Println("Длинна окружности равна", math.Sqrt(4*math.Pi*a)) /* вычисление прямо из площади*/
}

