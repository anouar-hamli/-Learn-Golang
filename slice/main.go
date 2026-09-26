package main

import "fmt"

func main() {
	// var slice1 = []int{1, 2, 3, 4, 5}
	// slice2 := []int{6, 7, 8, 9, 10}
	// fmt.Println("Slice 1:", slice1)
	// fmt.Println("Slice 2:", slice2)


	/* 
	Go Exercise: Arrays & Slices

Beginner One exercise

Exercise: Student Grades Manager

Your task: Create a Go program that manages student grades using an Array and a Slice.

Requirements

Create an Array named grades containing these 5 grades: 12, 15, 18, 10, 14.

Print the first and last grades using their indexes.

Change the third grade from 18 to 20.

Create a Slice named passed containing the grades that are greater than or equal to 12.

Add two new grades, 16 and 19, to the passed Slice using append().

Print the length and capacity of passed.

Use range to print each grade in passed with its index.

Create a new Slice named topGrades containing only the first three elements of passed, using slicing.
	*/



	grades := [5]int{12, 15, 18, 10, 14}
	fmt.Println("First grade:", grades[0])
	fmt.Println("Last grade:", grades[len(grades)-1])
	grades[2] = 20
	passed := []int{}
	for _, grade := range grades {
		if grade >= 12 {
			passed = append(passed, grade)
		}
	}
	fmt.Println("Passed grades:", passed)
	passed = append(passed,16,19)
	fmt.Println("Passed grades with new elements:", passed)
	fmt.Println("Length of passed:", len(passed))
	fmt.Println("Capacity of passed:", cap(passed))
	for index, grade := range passed {
		fmt.Printf("Index: %d, Grade: %d\n", index, grade)
	}
	topGrades := passed[:3]
	fmt.Println("Top grades:", topGrades)
}
