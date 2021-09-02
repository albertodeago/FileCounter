package filecounter

import (
	"io/fs"
)

// Function to count the number of files in a folder (and in subfolders)
// This function "manually" cycles through the filesystem synchronously
func FileCounterSync(fileSystem fs.FS) (int, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return 0, err
	}

	var numOfFiles int
	for _, f := range dir {
		if !f.IsDir() {
			// fmt.Printf("\nFound a file %s, before we had %d", f.Name(), numOfFiles)
			numOfFiles++
		} else {
			dirs, err := fs.ReadDir(fileSystem, f.Name())
			if err != nil {
				return 0, err
			}

			for _, file := range dirs {
				if !file.IsDir() {
					// fmt.Printf("\nFound a file %s, before we had %d", file.Name(), numOfFiles)
					numOfFiles++
				} else {
					n, err := countFilesRecursively(fileSystem, f.Name(), file)
					// fmt.Printf("\ncounted recursively %d, before we had %d", n, numOfFiles)
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
			// fmt.Printf("\nFound a file: %s", file.Name())
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
