// Chapter3 project main.go
package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 100; i++ {

		//		switch i % 3 {
		//		case 0:
		//			fmt.Println("Zero")
		//		case 1:
		//			fmt.Println("One")
		//		case 2:
		//			fmt.Println("Two")
		//		case 3:
		//			fmt.Println("Three")
		//		case 4:
		//			fmt.Println("Four")
		//		case 5:
		//			fmt.Println("Five")
		//		default:
		//			fmt.Println("Unknown Number")
		//		}

		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println(i, " - FizzBuzz")
			} else {
				fmt.Println(i, " - Fizz")
			}
		} else if i%5 == 0 {
			fmt.Println(i, " - Buzz")
		} /*else {
			fmt.Println(i)
		}*/

	}
}
