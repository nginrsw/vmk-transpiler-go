package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func reverseInFile(filePath string) error {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Replace 'local' with 'lock'
	content = regexp.MustCompile(`\blocal\b`).ReplaceAll(content, []byte("lock"))
	// Replace 'function' with 'fn'
	content = regexp.MustCompile(`\bfunction\b`).ReplaceAll(content, []byte("fn"))
	// Replace 'string.' with 'str.'
	content = regexp.MustCompile(`\bstring.\b`).ReplaceAll(content, []byte("str."))

	err = ioutil.WriteFile(filePath, content, 0644)
	if err != nil {
		return err
	}
	return nil
}

func reverseRenameFile(oldPath string) error {
	// Replace Lua with VMK in file names
	newPath := strings.Replace(oldPath, "LUA", "VMK", -1)
	newPath = strings.Replace(newPath, "Lua", "Vmk", -1)
	newPath = strings.Replace(newPath, "lua", "vmk", -1)

	return os.Rename(oldPath, newPath)
}

func traverseDirReverse(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Process files with .lua extension
		if strings.HasSuffix(info.Name(), ".lua") {
			// Reverse content inside file
			if err := reverseInFile(path); err != nil {
				return err
			}

			// Rename the file
			if err := reverseRenameFile(path); err != nil {
				return err
			}
		}
		return nil
	})
}

func main() {
	dir := "." // Specify the root directory to start the traversal
	if err := traverseDirReverse(dir); err != nil {
		fmt.Println("Error:", err)
	}
}
