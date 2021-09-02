package main

import (
	filecounter "filecounter/filecounter"
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		panic("Path must be specified as command line argument")
	}
	path := os.Args[1]
	fs := os.DirFS(path)

	start1 := time.Now()
	res1, _ := filecounter.FileCounterEasy(fs)
	elapsed1 := time.Since(start1)

	fmt.Printf("\nEasy\nCount: %d \n", res1)
	fmt.Printf("Ealapsed time: %s \n", elapsed1)

	start2 := time.Now()
	res2, _ := filecounter.FileCounterSync(fs)
	elapsed2 := time.Since(start2)

	fmt.Printf("\nSync\nCount: %d \n", res2)
	fmt.Printf("Ealapsed time: %s \n", elapsed2)

	start3 := time.Now()
	res3, _ := filecounter.FileCounterAsync(fs)
	elapsed3 := time.Since(start3)

	fmt.Printf("\nAsync\nCount: %d \n", res3)
	fmt.Printf("Ealapsed time: %s \n", elapsed3)

}
