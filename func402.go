package main

import "fmt"

func fun()int {
	a := 10
	b := 20
	msg :="welcome"
	fmt.Println(msg) // Function call 
	add:= a+b
	return add // return Value
}

func main() {
	x := fun()
	fmt.Print(x)
}