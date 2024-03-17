package gen

import (
	"os"
)

func CleanGeneratedFiles(genTarget string) {
	os.RemoveAll(genTarget)
}
