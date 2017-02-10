// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	t0 := time.Now()
	s, sep := "", ""
	for i, arg := range os.Args[:] {
		s += sep + arg
		sep = " "
		fmt.Println("Индекс аргумента:", i, "Значение аргумента:", arg)
	}
	fmt.Println(s)
	t1 := time.Nanosecond.(Now().Sub(t0))
	fmt.Println("The call took to run.\n", t1)
}

//!-
