// Chpter4 project main.go
package main

import "fmt"

func main() {
	//	var x [5]float64
	//	x[0] = 98
	//	x[1] = 93
	//	x[2] = 77
	//	x[3] = 82
	//	x[4] = 83

	//	var total float64 = 0
	//	for _, value := range x {
	//		total += value
	//	}
	//	fmt.Println(total / float64(len(x)))
	//	x := make(map[string]int)
	//	x["key"] = 10
	//	fmt.Println(x["key"])
	//	elements := map[string]string{
	//		"H":  "Hydrogen",
	//		"He": "Helium",
	//		"Li": "Lithium",
	//		"Be": "Beryllium",
	//		"B":  "Boron",
	//		"C":  "Carbon",
	//		"N":  "Nitrogen",
	//		"O":  "Oxygen",
	//		"F":  "Fluorine",
	//		"Ne": "Neon",
	//	}
	//	name, ok := elements["Li"]
	//	x := [6]string{"a", "b", "c", "d", "e", "f"}
	//	fmt.Println(x[2:5])
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	small := &x[0]
	for i, _ := range x {
		if *small > x[i] {
			small = &x[i]
		}
	}
	fmt.Println(*small)
}
