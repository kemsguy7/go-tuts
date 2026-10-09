package main

import "fmt"

/*
function syntax in GO
func functionName(parameterList)(returnType) {
	// Function Body
}

- void functions don't return anything and are declared without specifying a retun type
-  Functions as first-class citizens means:
1. They can be passed as arguments to other functions
2. Can be returned from other functions
3. Can be assigned to variables
- if all parameters have the same time, you can omit the type for all but the alst parameter

*/
func main() { 

	sayHello := func() {
		fmt.Println("Hello!")
	}
	
	sayHello()

	divide := func(dividend, divisor int) (int, error) { //func with multiple return values
		if divisor == 0 { 
			return 0, fmt.Errorf("cannot divide by zero")
		}
		return dividend / divisor , nil
	}
	result, err := divide(10, 2) 
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}


	//named return values 
	rectangleProperties := func(length, width int) (area, perimeter int)  {
		area = length * width 
		perimeter = 2 * (length + width)
		return //Automaticaly returns area and perimeter 
	}
	//calling the funciton 
	area, perimiter := rectangleProperties(10, 20)
	fmt.Println("Area:", area, "Perimeter.",perimiter)
	
	//using blank identifiers with return values 
	performCalculation := func(x, y int) (int, int) {
		sum := x + y 
		product := x * y 
		return sum, product 
	}
	sum, _:= performCalculation(3, 4) 

	fmt.Println("Sum:", sum, "Product:" ) 
	
	
	
	multiplier := func(factor int) func(int) int {
		return func(x int) int {
			return x * factor 
		}
	}
	double := multiplier(2)
	fmt.Println((double(5)))

	//variadic function take an arbitraty no of arguments 
	
	fmt.Println(sum(50, 9, 8 , 9))

}

func sum(numbers ... int) int {
	total := 0 
	for _, number := range numbers {
		total += number 
	}
	return total 
}
