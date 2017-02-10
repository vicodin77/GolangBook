// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	var (
		countByte int64
		filename  *os.File
		err       error
	)
	if len(os.Args[2:]) == 0 {
		fmt.Println("Usage: fetch  <filename> <web-address> <web-address> ...")
	}
	if Exists(os.Args[1]) {
		filename, err = os.OpenFile(os.Args[1], os.O_APPEND, os.ModeAppend /*0644*/)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch1o: open %s: %v\n", os.Args[1], err)
			os.Exit(1)
		}
	} else {
		filename, err = os.Create(os.Args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch1c: create %s: %v\n", os.Args[1], err)
			os.Exit(1)
		}
	}
	defer filename.Close()
	startglob := time.Now()
	for _, url := range os.Args[2:] {
		if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch2: %v\n", err)
			os.Exit(1)
		}
		_, _ = filename.WriteString("--------- " + url + " ---------\n")
		//        if err != nil {
		//            fmt.Fprintf(os.Stderr, "fetch2.5: write sep %s: %v\n", url, err)
		//            os.Exit(1)
		//        }
		start := time.Now()
		b, err := io.Copy(filename, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch3: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		secs := time.Since(start).Seconds()
		_, _ = filename.WriteString("\n")
		countByte += b
		fmt.Println(url, resp.Status, secs)
	}
	secsglob := time.Since(startglob).Seconds()
	fmt.Printf("\nЗаписано %d байт за %f секунд\n", countByte, secsglob)
}

// Exists reports whether the named file or directory exists.
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

//!-
