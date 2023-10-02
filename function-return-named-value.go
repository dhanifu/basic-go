package main

import "fmt"

func getFullName2() (firstName string, middleName string, lastName string) {
	firstName = "Muhammad"
	middleName = "Ramdhani"
	lastName = "Gaada"

	return
}

func getIdentity() (name string, age int, married bool) {
	name = "Muhammad Ramdhani"
	age = 20
	married = false

	return 
}

func main() {
	a, b, c := getFullName2()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	name, age, married := getIdentity()

	fmt.Println("name    :",name)
	fmt.Println("age     :",age)
	fmt.Println("married :",married)
}
