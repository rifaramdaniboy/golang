package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Budi"
	names[1] = "Rifa Ramdani"
	names[2] = "Dimas Saputra"

	fmt.Println(names[2])
	fmt.Println(names[0])
	fmt.Println(names[1])

	var values = [3]int{
		95,
		85,
		75,
	}

	fmt.Println(values)
	fmt.Println(values[2])
	fmt.Println(values[0])
	fmt.Println(values[1])
}
