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
        if j === 3 {
            break
        }
    }
    // Using range with an array
    fmt.Println("\n Using range with an array: ")
    nums := []int{10, 20, 30}
    for index, value := range nums { 
        fmt.Printf()
    }
}