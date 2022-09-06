//Task1: Swap 2 numbers. In this task, a user is asked to enter two numbers 
//and the program will swap two numbers without using the third variable.

package main

import (
	"fmt"
)

func main(){
	var a, b int = 3,7

	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	fmt.Printf("a = %d, b = %d\n", a,b)

	var c int = a
	a = b
	b = c
	fmt.Printf("a = %d, b = %d\n", a, b)
}