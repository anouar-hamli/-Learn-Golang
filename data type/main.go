package main
import (
	"fmt"
	
)
func main() {
	//boolean
	//true or false
	// 1-0
	mybool := true
	fmt.Println("mybool:", mybool)
	fmt.Printf("%v %T\n\n ",mybool,mybool)


	//integer
	myint := 42
	fmt.Println("myint:", myint)
	fmt.Printf("%v %T\n\n ",myint,myint)
	myInt8 := 127*22
	fmt.Println("myInt8:", myInt8)
	fmt.Printf("%v %T\n\n ",myInt8,myInt8)

	//string
	mystr:= "Hello, World!"
	fmt.Println("mystr:", mystr)
	fmt.Printf("%v %T\n\n ",mystr,mystr)
	//float=floating
	//440.3
	myfloat := 440.3
	fmt.Println("myfloat:", myfloat)
	fmt.Printf("%v %T\n\n ",myfloat,myfloat)

	//[]string array of slices
	var mySlice []string
	mySlice = append(mySlice, "Hello")
	mySlice = append(mySlice, "World")
	fmt.Println("mySlice:", mySlice)
	fmt.Printf("%v %T\n\n ",mySlice,mySlice)

} 
	
