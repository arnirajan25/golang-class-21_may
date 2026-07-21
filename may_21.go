package main

import "fmt"

var name string = "suhag"
var nirajan string
var age int
var temp float64
var choice bool

func modifybyvalue(val int) {
	val = 989
}

func modbyadd(val *int) {
	*val = 100
}

func main() {
	name1 := "Nirajan"
	age := 25
	ptr := &age
	*ptr = 20
	var ok *string = &name1

	fmt.Println(*ptr)
	fmt.Println(&ok)
	fmt.Println(*ok)
	fmt.Println(name1)
	fmt.Println(&ptr)
	fmt.Println(age)
	fmt.Println(&age)

	name := "suha"
	fmt.Println(name, nirajan, age, temp, choice)

	x := 250
	modifybyvalue(x)
	fmt.Println(x)

	modbyadd(&x)
	fmt.Println(x)
}
