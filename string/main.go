package main

import (
	"fmt"
	"strings"
)
func main() {
	var firstName string = ""
	var lastName string = ""
	var email string = ""
	fmt.Println("Enter your first name and last name:")
	fmt.Scan(&firstName, &lastName)
	fmt.Println("Enter your email address:")
	fmt.Scan(&email)
	fullName := firstName + " " + lastName
	fmt.Println("Your full name is:", fullName)
	fmt.Println("Your email address is:", email)
	fmt.Println("The length of your full name is:", len(fullName))
	fmt.Printf("The fifth character of your full name is: %c\n", fullName[4])
	fmt.Println("Your full name in uppercase is:", strings.ToUpper(fullName))
	fmt.Println("Your full name in lowercase is:", strings.ToLower(fullName))
	fmt.Println("Does your full name contain 'go'?", strings.Contains(fullName, "go"))
	fmt.Println("Does your email address contain '@'?", strings.Contains(email, "@"))
	var username = fullName
	fmt.Println("Your username is:", username)
	fmt.Println("word is:", firstName+"\n"+lastName)
	fmt.Println("Email ends with .com:", strings.HasSuffix(email, ".com"))
	

}