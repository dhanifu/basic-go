package main

import "fmt"

func main() {
	var months = [...]string {
		"Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember",
	}

	var slice1 = months[4:7]

	fmt.Println(months)
	fmt.Println(slice1)
	fmt.Println(len(slice1)) // length
	fmt.Println(cap(slice1)) // capacity

	months[5] = "June"
	fmt.Println(slice1)

	slice1[0] = "Mei DIubah"
	fmt.Println(slice1)

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Dhani")
	fmt.Println(slice3)

	slice3[1] = "Anjay"
	fmt.Println(slice3)
}