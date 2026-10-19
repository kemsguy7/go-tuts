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

//variadic function take an arbitrary no of arguments
func sumFunc(numbers ...int) int {
	total := 0 
	for _, number := range numbers {
		total += number 
	}
	return total 
}

func printNames(names ...string)  {
	for i, name := range names {
		fmt.Printf("Name %d: %s\n", i + 1, name) 
	}
}

//variadic functions with multiple return values
func average(numbers ...float64)(float64, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("no numbers provided")
	}
	total := 0.0 
	for _, number := range numbers {
		total += number
	}
	return total / float64(len(numbers)), nil 

}



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


	//printing sunFunc
	fmt.Println(sumFunc(50, 9, 8 , 9))

	// printing printNames
	printNames("Matthew", "Udo", "Zoe")



	Avgresult, Avgerr := average(2.0, 4.0, 6.0)
	if err != nil {
		fmt.Println("Error: ", Avgerr)
	} else {
		fmt.Println("Average", Avgresult) 
	}

}


