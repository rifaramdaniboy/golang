package main

import "fmt"

func main() {
	var ujian = 80
	var absensi = 70

	var lulusUjian = ujian >= 80
	var lulusAbsensi = absensi <= 70
	fmt.Println(lulusAbsensi)
	fmt.Println(lulusUjian)

	var lulus = lulusAbsensi && lulusUjian
	fmt.Println(lulus)
}
