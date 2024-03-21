package main

import "fmt"

func main() {
	type Noktp string
	type Married bool

	var noktpboy Noktp = "37232126110211119"
	var marriedstatus Married = true

	fmt.Println(noktpboy)
	fmt.Println(marriedstatus)
}
