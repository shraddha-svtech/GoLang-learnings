package main //package(collection of files) "main" tells our code must be compiled as an executable program
import "fmt" //package from GO standard library for formatting strings and printing messages to console

func main(){ //only 1 main function --> an entry point of the program
	 fmt.Println("Hello, Shraddha!")  //fmt.PackageName()


	 //LESSON 2: VARIABLES (string, numbers, boolean, float, constants)
	 var name1 string = "Shraddha" //we use "" double quotes for string values, and we use single quotes for character values
	 var name2 = "Dongol"
	 var name3 string //we can declare a variable without assigning a value to it, and we can assign a value later
	 fmt.Println("Hello," , name1, name2) //print the values of name1 and name2 variables

	 name1 = "Shraddhaaaaaaa"
	 name2 = "Dongolllllll"
	 name2 = "Shraddha Dongol"
	 fmt.Println(name1, name2, name3)

	 //shorthand for declaring and assigning a variable
	 name4 := "Short" //same as :  var name4 string = "shraddha Dongol"
	 fmt.Println("Shorthand for declaration: var name4 string = \"Short\" ----> name4 := ", name4)

	 var number1 int = 10
	 var number2 =  20
	 number3 := 30
	 fmt.Println("Numbers: ", number1, number2, number3)

	 //bits and memory
	 var num int8 = 127 //int8 can store values from -128 to 127
	 var num2 int16 = 32767 //int16 can store values from -32768 to 32767
	 var num3 uint8 = 25 //uint can store values from 0 to 255 ( no negative values)
	fmt.Println("Bits : ", num, num2, num3)
	 var price float32 = 12.23//float32 can store decimal values, but it is less precise than float64
	 var price2 float64 = 12.983424324324 //float64 is more precise than float32
	price3 := 12.983424324324 //shorthand for float64
	price4 := 12.11
	fmt.Println("Float: ", price, price2, price3, price4)

	var isActive bool = true //bool can store true or false values
	isDarkMode := false //shorthand for bool
	fmt.Println("Boolean: ", isActive, isDarkMode)

	const pi float64 = 3.14 //constants are values that cannot be changed after they are declared
	const pi2 = 3.14159 //shorthand for constants
	const pi3 = 3.141592653589793238462643383279502884197169399375105820974944592307816406286208998628034825342117067982148086513282306647093844609550582231725359408128481117450284102701938521105559644622948954930381964428810975665933446128475648233786783165271201909145648566923460348610454326648213393607260249141273724587006606315588174881520920962829254091715364367892590360011330530548820466521384146951941511


}