package main

import "fmt"

var(
	area = func(length int,width int)int {
		return length * width
	}
)


func main() {
	func(a int,b int) {
		fmt.Println(a*b)
	}(20,30)

		fmt.Printf(
			"%.2f c",
			func (f float64) float64 {
				return (f-32.0) * (5.0/9.0)
			}(100),
		)
}