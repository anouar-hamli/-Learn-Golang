package main

import (
	"fmt"
)
func main() {
	// \b= backspace
	fmt.Println("Hello, \bWorld!\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b\b")
	//\n= new line
	fmt.Println("Hello,\n \nWorld!")
	//\r= carriage return
	fmt.Println("Hello, \rWorld!")

	//\t= tab
	fmt.Println("Hello, \tWorld!")
	// \\= backslash
	fmt.Println("Hello, \\World!")
	fmt.Println(`hello \ word`)

}