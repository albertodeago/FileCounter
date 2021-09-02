package filecounter

import (
	"fmt"
	"io/fs"
)

// Function to count the number of files in a folder (and in subfolders)
// This function uses the fs.WalkDir under the hood, so it's dead easy
func FileCounterEasy(fileSystem fs.FS) (int, error) {
	// fmt.Println("Easy FileCounter started")

	var numOfFiles int
	err := fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}

		if !d.IsDir() {
			numOfFiles++
		}

		return nil
	})

	if err != nil {
		return 0, nil
	}

	return numOfFiles, nil
}
