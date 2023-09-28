package main

import "fmt"

func main() {
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}
	
	fmt.Println()

	slice := []string{"Muhammad", "Ramdhani", "Khannedy", "Budi", "Joko"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	fmt.Println()

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
		// fmt.Println(value)
	}

	fmt.Println()

	person := make(map[string]string)
	person["name"] = "Eko"
	person["title"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
