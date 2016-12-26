package generator

import (
	"fmt"
	"io/ioutil"
	"os"
)

// FileGenerator to generate go source code files
type FileGenerator struct {
	Package   string
	Imports   []Generator
	Constants []Generator
	Functions []Generator
	Structs   []Generator
}

// Generate a string representation of the file
func (g *FileGenerator) Generate() string {
	return ""
}

// Export the file contents as a Go source code file
func (g *FileGenerator) Export(path string) (bytes int64, err error) {
	err = ioutil.WriteFile(path, []byte(g.Generate()), 0644)

	if err != nil {
		return 0, err
	}

	fileInfo, err := os.Stat(path)

	if err != nil {
		err = fmt.Errorf(`Could not read file at "%s"\n`, path)
	}

	return fileInfo.Size(), nil
}
