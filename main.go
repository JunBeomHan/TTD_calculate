package main

import "fmt"

func main() {
	Test()
}

func Test() {
	o := Calculate("+", 3, 2)
	if o != 5 {
		fmt.Printf("Test Faild! expected:%d output:%d\n", 5, o)
		return
	}
}
