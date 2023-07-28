package main

import "fmt"
func main(){
	var fullName string
	fmt.Print("Enter name: ")
	fmt.Scan(&fullName)
	fmt.Printf("%v is a student",fullName)
}