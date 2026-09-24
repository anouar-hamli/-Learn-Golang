package main
 import (
 	"fmt"
 )
 func main() {
	fmt.Println("Enter your name:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("Hello, " + name + "!")
 }