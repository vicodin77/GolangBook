// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	var (
		countByte int
		filename  *os.File
		err       error
	)
	ch := make(chan string)
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
		go fetch(url, ch) // start a goroutine
		//		countByte += b
		//		fmt.Println(url, resp.Status, secs)
	}
	for _, url := range os.Args[2:] {
		//		fmt.Println(<-ch) // receive from channel ch
		b, err := io.WriteString(filename, <-ch)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch3: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		countByte += b
		secsglob := time.Since(startglob).Seconds()
		fmt.Printf("\nЗаписано из %s %d байт за %f секунд\n", url, b, secsglob)
	}
	secsglob := time.Since(startglob).Seconds()
	fmt.Printf("\nВсего записано %d байт за %f секунд\n", countByte, secsglob)
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

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	//	nbytes, err := io.Copy(ch <-, resp.Body)
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs %s", secs, url)
	ch <- "--------- " + url + " ---------\n" + string(b) + "\n"

}

//!-
