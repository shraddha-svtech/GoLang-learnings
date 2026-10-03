package main

import (
	"fmt"
	"strings"
	"sort"

	//different standard library packages are provided by GO for different purposes, and we can import them as per our needs. For example, "fmt" package is used for formatting strings and printing messages to console, and "strings" package is used for manipulating strings.
	// Eg: strings.Contains(), strings.ReplaceAll(), strings.ToUpper(), strings.ToLower(), strings.TrimSpace(), strings.Split(), strings.Join() etc.
)

func main(){

	//1. strings package:
	var greetings = "Hello, World!"
	fmt.Println(greetings)

	// fmt.Println(strings.Contains(greetings, "hello!"))
	// fmt.Println(strings.Contains(greetings, "hello"))
	// fmt.Println(strings.Contains(greetings, "Hello"))

	fmt.Println("Replaced with strings package's method : " , strings.ReplaceAll(greetings, "Hello", "Namaste"));
	fmt.Println("Original Value still same : ", greetings)

	fmt.Println("Finding index : ", strings.Index(greetings,"e"))
	fmt.Println("Finding last index : ", strings.LastIndex(greetings,"l"))

	fmt.Println("Splitting string : ", strings.Split(greetings, " "))
	fmt.Println("Spliting strings with l : ", strings.Split(greetings,"l"))


	//2. sort package:
	var age = []int{24, 30, 18, 22, 28}
	fmt.Println("Age before sorting : ", age)
	sort.Ints(age)
	fmt.Println("Age after sorting : ", age)


	index := sort.SearchInts(age, 18)
	fmt.Println("Index of 18 in sorted slice : ", index)


	fruits := []string{"Banana", "Apple", "Mango", "Grapes"}
	fmt.Println("Fruits before sorting : ", fruits)
	sort.Strings(fruits)
	fmt.Println("Fruits after sorting : ", fruits)

	fmt.Println("Index of Banana in sorted slice : ", sort.SearchStrings(fruits, "Banana"))


}