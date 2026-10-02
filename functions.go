
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
- if all parameters ahve the same time, you can omit the type for all but the alst parameter 

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
}


