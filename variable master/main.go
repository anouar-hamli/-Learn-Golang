package main

import "fmt"

func main(){
	name:="yasser"
	age:=23
	height:=1.67
	isStudent:=true
	
	fmt.Println("his name is "+name)
	fmt.Println("his age is ",age)
	fmt.Println("his height is ",height)
	fmt.Println("is he student ",isStudent)

//THE SECOUND EXERSICE
	productName:="LAPTOP"
	price:=800.2
	quantity:=9
	total:=price* float64(quantity)
	discount := 10
	fmt.Println("product name is "+productName)
	fmt.Println("price is ",price)
	fmt.Println("quantity is",quantity)
	fmt.Println("the total is ",total)
	withDiscount:=total*float64(discount) / 100
	fmt.Println("the total is ",withDiscount)
	finalPrice:=total-withDiscount
	fmt.Println("this is the final price",finalPrice)
// THE THIRD EXERCISE
	math:= 16.9
	programming:=18.9
	english:=17.5
	average := (math+programming+english)/3
	fmt.Println("your math note is ",math)
	fmt.Println("your programming note is ",programming)
	fmt.Println("your english note is ",english)
	fmt.Println("the average note is ",average)
	
}