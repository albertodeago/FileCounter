package filecounter

import (
	"fmt"
	"io/fs"
	"sync"
)

// Function to count the number of files in a folder (and in subfolders)
// This function "manually" cycles through the filesysem, spawning a gorouting for each folder
func FileCounterAsync(fileSystem fs.FS) (int, error) {
	// 1. read the root dir
	// for each subdir call a goroutine that:
	// for each subdir call the same goroutin
	// for each file send a +1 in the chnnel
	// wait for all the grouting to end
	// return the counter

	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return 0, err
	}

	var numOfFiles int
	var wg sync.WaitGroup
	c := make(chan int)

	for _, f := range dir {
		if !f.IsDir() {
			numOfFiles++
		} else {
			// fmt.Println("Opening group for " + f.Name())
			wg.Add(1)
			go countFiles(fileSystem, &wg, f.Name(), c)
		}
	}

	go func() {
		wg.Wait()
		// fmt.Println("Closing the channel")
		close(c)
	}()

	for v := range c {
		// fmt.Printf("chan value %d \n",v)
		numOfFiles += v
	}

	return numOfFiles, nil
}

func countFiles(fileSystem fs.FS, wg *sync.WaitGroup, path string, c chan int) {
	defer func() {
		// fmt.Println("closing group for " + path)
		wg.Done()
	}()

	dirs, err := fs.ReadDir(fileSystem, path)
	if err != nil {
		fmt.Printf("Error while reading %s", path)
		fmt.Printf("Error: %W", err)
	} else {
		for _, f := range dirs {
			if !f.IsDir() {
				// fmt.Println("counting file " + f.Name())
				c <- 1
			} else {
				// fmt.Println("Opening group for: " + f.Name())
				wg.Add(1)
				go countFiles(fileSystem, wg, path+"/"+f.Name(), c)
			}
		}
	}
}
