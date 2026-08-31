package main 
import "fmt"



// Basic Data Types includes integers, floats, Booleans,  strings, and complex numbers. 
var a int = 10 // Declaring an interger variable 'a' with the value 10 

var name = "Matthew"

func main() {
	fmt.Println("Hello, World!");


 	age := 20 // Shorthand declaration where GO infers 'b' as an integer (can only be used in functions not package level)
	greeting :="Hello, " + name
	println(greeting, "You are", age, "years old.") 
}

// Variable Types and declaration Styles
var x int     // Declares an integer variable 'x' with default value 0


//Multiple Declarations 
var a, b, c int = 1, 2, 3
var x, y = 10, "Hello"    // Different types in the same line


// Block declaration Style
var ( 
	name string ="Shubham"
	age int = 30 
	height float64 = 5.9 
)

// constants : only certain types are allowed : Boolean, numeric and string

var isAdult bool = age >= 18;
func checkAdult
if isAdult {
	fmt.PrintIn("You are an adult") ;
}

// strings in go are immutable once declared 