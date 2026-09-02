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

	var isAdult bool = age >= 18;
	if isAdult {
		fmt.Println("You are an adult") ;
	}


	var firstName string = "Matthew" 
	var lastName string = "Idungafa"
	fullName := firstName + " " + lastName 
	fmt.Println(fullName)

	// z := 3 + 4i // Creates a complex number with real part 3 and imaginary part 4
	z := complex(6, -8) //complex number 6 - 8i 
	r := real(z) // ri is 6
	i := imag(z) // i is -8

	fmt.Println("Real part:", r)
	fmt.Println("Imaginary Part:", i)

	//Addition: The sum of two complex numbers (a+bi)+(c+di) results in (a+c)+(b+d)i.

	z1 := 2 + 3i 
	z2 := 1 + 4i 
	sum := z1 + z2 //Result: 3 + 7i 
	fmt.Println("The sum is : ", sum )

}

// Variable Types and declaration Styles
//var x int     // Declares an integer variable 'x' with default value 0


//Multiple Declarations 
//var a, b, c int = 1, 2, 3
//var x, y = 10, "Hello"    // Different types in the same line


// Block declaration Style
var ( 
	uname string ="Shubham"
	age int = 30 
	height float64 = 5.9 
)

// constants : only certain types are allowed : Boolean, numeric and string

var greeting string = "Hello"
// len() - used to return the length of a string


// strings in go are immutable once declared 

/* 
floats: float64 is a more precise float and is the recommended type for calculations requiring high accuracy 

- Complex numbers are numbers that include both a real part and an imaginary part, they are commonly represented in the form a+bi  where : 
 a is the real part (a real number )
 b is the imaginary part (multiplied by i , where i ^2 =1)
 Direct Assignment: Complex numbers can also be directly assigned as a + bi, where a and b are floating-point literals

*/

var z1 complex64 = complex(5, 7) 
 //Create a complex number with real part3 and imaginary part4 