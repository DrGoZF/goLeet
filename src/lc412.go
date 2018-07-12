package src

import (
    "strconv"
    "fmt"
)

func main() {
	fmt.Println(fizzBuzz(1))
}

func fizzBuzz(n int) []string {
    var result []string
    for i := 0; i < n; i ++ {
	    if i % 15 == 0 {
		    result[i] = "FizzBuzz"
	    } else if i % 3 == 0 {
		    result[i] = "Fizz"
	    } else if i % 5 == 0 {
		    result[i] = "Buzz"
	    } else {
		    result[i] = strconv.Itoa(i)
	    }
    }
    return result
}

