package goUtils

import (
	"os"
	"fmt"
)

func GetFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil
	}
	return file
}