package main

import "fmt"

func main() {

    /* 
LOOPS
- only Looping construct in go is the "for loop"
- for loop without initialization makes it behave like a while loop
- for loop without any conditions or components make it a while loop 
- 
BEST PRACTICES FOR LOOPS IN GO 
- Avoid infinite loops unless necessary 
- Use range for simplicity 
- Limit the use of goto 
*/


    fmt.Println("Standard for loop:")

    for i := 0; i < 3; i++ {
        fmt.Println(i)
    }
    // While-like for loop
    fmt.Println("\nWhile-like for loop:")
    count := 0 
    for count < 3 { 
        fmt.Println(count)
        count++
    }
    // infinite loop with break 
    fmt.Println("\nInfinite loop with break:")
    j := 0 
    for {
        fmt.Println(j)
         j++ 
        if j == 3 {
            break
        }
       
    }
    // Using range with an array
    fmt.Println("\n Using range with an array: ")
    nums := []int{10, 20, 30}
    for index, value := range nums { 
          fmt.Println("Index:", index, "Value:",value)
    }

}

/* 
Questions
1. What are the three main components of a for loop in Go, and what
purpose does each serve?

- Initialization : sets the condition of the loop 
- condition: continues as long as the condition is true
- Post:  Executes after each iteration, typically used to increment or decrement a counter

2. How does Go’s for loop behave when it is used without any
conditions? : It becomes and infinite Loop
3. Explain how the range keyword is used to iterate over arrays and maps
in Go.
Ans: syntax is 
for index, value := range collection {

}
4. What is the effect of using the break statement within a for loop?
Provide an example.
Ans: Break is used to stop an infinite loop based on a specified condition e.g if i >= 3{break}

5. Describe a scenario where using continue in a loop would be useful.
 -  As

 6. How can the for loop in Go mimic a while loop?
Ans: By not adding the initialization statements 

7. Why is the range keyword advantageous when working with
collections in Go?
- For each data type, the range privides a way to access each element and, if applicable, it's index or key 

8. Write a for loop that prints all numbers from 1 to 10 but skips the
number 5.
- Ans: 

9. When should you avoid using an infinite loop, and how can you
terminate one safely if needed?
- 
10. Explain the significance of the goto statement in loop control and why
it is generally discouraged.


*/