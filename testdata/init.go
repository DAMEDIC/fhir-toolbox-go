package testdata

import (
	"os"
	"path"
	"runtime"
)

// init changes the working directory to the root of the project
// https://stackoverflow.com/a/70050794
func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}
