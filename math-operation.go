package main

import "fmt"

func main() {
	//Operasi Matematika
	fmt.Println("-----------------------------------")
	var result = 10 + 50
	fmt.Println(result)

	//Operasi Matematika
	var a = 50
	var b = 100
	var c = a * b
	var d = b - a
	var e = b / a
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	//Augmented Assigments
	fmt.Println("-----------------------------------")
	var i = 100
	i += 100
	fmt.Println(i)

	//Unary Operator
	fmt.Println("-----------------------------------")
	i++ // i = i + 1
	fmt.Println(i)

	//Positive & Negative
	fmt.Println("-----------------------------------")
	var negative = -100
	var positive = +500
	fmt.Println(negative)
	fmt.Println(positive)
}
