package main

import "fmt"

/*
Customer Struct
*/
type Customer struct {
	name string
	age  int
}

func (cus *Customer) yell() {
	fmt.Println("Can i buy a banana")
}
