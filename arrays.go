package main

import "fmt"

func mainTest(){
	var array [3]int = [3]int{10,20,30}
	var arrays = [3]int{40,50,60}
	fmt.Println(array)
	fmt.Println("Shorthand : ",arrays)
	fmt.Println("Length of array : ",len(array))

	names := [5]string{"Shraddha", "Gautam", "Lark", "Ram"}
	fmt.Println("Names : ",names)
	fmt.Println("Length of names array : ",len(names))

	//accessing array elements
	fmt.Println("First element of names array : ",names[0])
	fmt.Println("Second element of names array : ",names[1])

	//modifying array elements
	names[0] = "Shraddha Dongol"
	fmt.Println("Modified Names : ",names)


	//SLICES  - use arrays under the hood, but they are more flexible and powerful than arrays. Slices are dynamic in size, and they can grow and shrink as needed. Slices are more commonly used than arrays in Go.
	var scores = []int{100,200}
	fmt.Println("Scores : ",scores, " and length of scores : ",len(scores))
	scores =append(scores, 300) //append() function is used to add elements to the end of a slice. It returns a new slice with the added elements.
	fmt.Println("Scores after append : ",scores, " and length of scores : ",len(scores))

	//slice ranges
	rangeOne := names[1:3] //slice from index 1 to index 2 (3-1), includes index 1,2 not 3	
	fmt.Println("Range One : ",rangeOne, " and length of rangeOne : ",len(rangeOne))
	rangeTwo := names[2:] //slice from index 2 to the end of the array
	fmt.Println("Range Two : ",rangeTwo, " and length of rangeTwo : ",len(rangeTwo))
	rangeThree := names[:3] //slice from the beginning of the array to index 2 (3-1), includes index 0,1,2 not 3
	fmt.Println("Range Three : ",rangeThree, " and length of rangeThree : ",len(rangeThree))


	//Difference between arrays and slices
	//1. Arrays have a fixed size, while slices are dynamic in size.
	//2. Arrays are value types, while slices are reference types.
	//3. Arrays are less flexible and powerful than slices.
}