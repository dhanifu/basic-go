package main

import "fmt"

func main() {
	var nilais [3]byte
	var nilai = [3]byte{
		80, 90, 100,
	}

	nilais[0] = 70
	nilais[1] = 80
	nilais[2] = 90

	fmt.Println(nilai)
	fmt.Println(nilais)
}