//Task2: isOdd and isEven. Write go functions to check whether a number is even and is odd.

package main

import (
	"fmt"
)

func main(){
	var a, b int = 3,7

	fmt.Printf("a = %d, b = %d\n", a, b)

	fmt.Println(a, "is odd: ", isOdd(a) )
	fmt.Println(b, "is even: ", isEven(b))
}
	func isOdd(num int) bool  {
		 return (num % 2 != 0)
		}
	

	func isEven(num int) bool {
		return (num % 2 == 0)
	}
	

