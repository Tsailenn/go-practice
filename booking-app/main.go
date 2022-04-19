package main

import "fmt"

func main() {

	confName := "Go conference"
	const tickets = 50
	var remainingTicket uint = tickets

	fmt.Println(confName)
	fmt.Printf("monkaw %v pog\n", confName)
	fmt.Printf("%v left\n", remainingTicket)

	var userName string
	var userTicket uint

	fmt.Scan(&userName)
	fmt.Scan(&userTicket)

	fmt.Printf("name: %v\nordered: %v", userName, userTicket)

	//fmt.Println("bye")
}
