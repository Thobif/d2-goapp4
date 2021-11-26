package main

import "fmt"

func add() {
	a:=10
	b:=5
	add:= a+b
	fmt.Printf("add : %v = %v + %v\n",add,a,b)
}

func sub() {
	a:=10
	b:=5
	sub:= a-b
	fmt.Printf("sub : %v = %v - %v\n",sub,a,b)
}

func mul() {
	a:=10
	b:=5
	mul:= a*b
	fmt.Printf("mul : %v = %v * %v\n",mul,a,b)
}
func div() {
	a:=10
	b:=5
	div:= a/b
	fmt.Printf("div : %v = %v / %v",div,a,b)

}




func main() {

	add()
	sub()
	mul()
	div()	
}

