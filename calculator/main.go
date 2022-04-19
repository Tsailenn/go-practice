package main

import (
	"fmt"
)

func main() {

}

func observe(i struct{}) {

	// using the format specifier
	// %T to check type in interface
	fmt.Printf("The type passed is: %T\n", i)

	// using the format specifier %#v
	// to check value in interface
	fmt.Printf("The value passed is: %#v \n", i)
	fmt.Println("-------------------------------------")
}
