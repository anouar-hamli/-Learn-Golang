package main
import (
	"fmt"
)
func main() {
	var myVar int = 42
	fmt.Println("myVar:", myVar)
	myVar = 100
	fmt.Println("myVar:", myVar)
	var myVar2 int
	myVar2 = 200
	fmt.Println("myVar2:", myVar2)
	myVar3 := ""
	myVar3 = "Hello, World!"
	fmt.Println("myVar3:", myVar3)

	tisro,smana := 10, 20
	fmt.Println("tisro:", tisro)
	fmt.Println("smana:", smana)

	var a, b = 10, 20
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	var myfloat float64 = 3.14
	fmt.Println("myfloat:", myfloat)
	myslice :=[]int{1,2,3,4}
	fmt.Println("myslice:", myslice)
}