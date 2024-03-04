package main

import "fmt"

func main() {
	// declaring variables in go
	var name1 string = "shivansh"
	var num1 int = 32
	var num2 float32 = 32.3 //go requires type declation  here in this example float is assigned as

	// go also have type inference

	var lastname = "aggarwal"
	const num4 = 30 //cannot change a const value

	// num4=50

	// declaring an empty variable in go

	var name2 string

	name2 = "hello everyone"

	// note: an normal integer can be type interfered but we need to declare float value because float values are assigned with memory

	var gpa float32 = 9.2
	var gpa2 float64 = 9.5
	var marks uint32 = 10

	// there is a new syntax in go for variable declaration

	marks4 := 30

	// marks4 ="shivansh"

	fmt.Println(gpa, gpa2, marks)
	fmt.Println(marks4)

	fmt.Println(name1, num1, num2, lastname)
	fmt.Println(name2)
}
