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

// Function to count the number of files in a folder (and in subfolders)
// This function "manually" cycles through the filesystem synchronously
func FileCounterSync(fileSystem fs.FS) (int, error) {
	// fmt.Println("Sync FileCounter started")

	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return 0, err
	}

	var numOfFiles int
	for _, f := range dir {
		if !f.IsDir() {
			if err != nil {
				return 0, err
			}
		} else {
			dirs, err := fs.ReadDir(fileSystem, f.Name())
			if err != nil {
				return 0, err
			}

			for _, file := range dirs {
				if !file.IsDir() {
					numOfFiles++
				} else {
					n, err := countFilesRecursively(fileSystem, f.Name(), file)
					if err != nil {
						return 0, err
					}
					numOfFiles += n
				}
			}
		}
	}

	return numOfFiles, nil
}

// Helper function for "FileCounterSync"
func countFilesRecursively(fileSystem fs.FS, prevPath string, dir fs.DirEntry) (int, error) {
	var n int
	newPath := prevPath + "/" + dir.Name()
	dirs, err := fs.ReadDir(fileSystem, newPath)

	if err != nil {
		return 0, err
	}

	for _, file := range dirs {
		if !file.IsDir() {
			n++
		} else {
			num, err := countFilesRecursively(fileSystem, newPath, file)
			if err != nil {
				return 0, err
			}

			n += num
		}
	}

	return n, nil
}

func FileCounterAsync(fileSystem fs.FS) (int, error) {
	return 0, nil
}
