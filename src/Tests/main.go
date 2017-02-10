// Tests project main.go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	//	xs := []float64{98, 93, 77, 82, 83}
	//	fmt.Println(sum(1, 2, 3, 4, 5))
	str1 := "Привет всем!"
	//	strlen := len(str1)
	//	runeStr := []rune(str1)
	fmt.Print("Строка \"", str1, "\" имеет длинну ", strconv.Itoa(len([]rune(str1))), "\n")
}

//func sum(x1 []float64) int {
//	var sum int
//	for _, a := range x1 {
//		if maximum < a {
//			maximum = a
//		}
//	}
//	return maximum
//}
