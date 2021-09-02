package filecounter

import (
	"io/fs"
	"testing"
	"testing/fstest"
)

var (
	fakeFS = fstest.MapFS{
		"root-folder":                                  {Mode: fs.ModeDir},
		"root-folder/file-1.md":                        {Data: []byte("I'm a file in the root")},
		"root-folder/sub-folder-1":                     {Mode: fs.ModeDir},
		"root-folder/sub-folder-2":                     {Mode: fs.ModeDir},
		"root-folder/sub-folder-2/file-1.md":           {Data: []byte("I'm a file in folder2")},
		"root-folder/sub-folder-2/file-2.md":           {Data: []byte("I'm another file in folder 2 ")},
		"root-folder/sub-folder-3":                     {Mode: fs.ModeDir},
		"root-folder/sub-folder-3/sub-sub-1":           {Mode: fs.ModeDir},
		"root-folder/sub-folder-3/sub-sub-1/file-1.md": {Data: []byte("file")},
	}
)

func TestFileCounter(t *testing.T) {
	t.Run("should read the number of files in a fileSystem (excluding folders) using WalkDir function", func(t *testing.T) {
		got, err := FileCounterSync(fakeFS)
		want := 4

		if err != nil {
			t.Errorf("Didnt expected an error, but got one %s", err)
		}

		if got != want {
			t.Errorf("got %d wanted %d", got, want)
		}
	})
	t.Run("should read the number of files in a fileSystem (excluding folders) syncronously", func(t *testing.T) {
		got, err := FileCounterEasy(fakeFS)
		want := 4

		if err != nil {
			t.Errorf("Didnt expected an error, but got one %s", err)
		}

		if got != want {
			t.Errorf("got %d wanted %d", got, want)
		}
	})
	t.Run("should read the number of files in a fileSystem (excluding folders) asyncronously", func(t *testing.T) {
		got, err := FileCounterAsync(fakeFS)
		want := 4

		if err != nil {
			t.Errorf("Didnt expected an error, but got one %s", err)
		}

		if got != want {
			t.Errorf("got %d wanted %d", got, want)
		}
	})
}

func BenchmarkFileCounterSync(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		FileCounterSync(fakeFS)
	}
}
func BenchmarkFileCounterAsync(b *testing.B) {
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		FileCounterAsync(fakeFS)
	}
}
