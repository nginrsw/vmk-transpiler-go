package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func replaceInFile(filePath string) error {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Replace 'lock' with 'local'
	content = regexp.MustCompile(`\blck\b`).ReplaceAll(content, []byte("local"))
	// Replace 'fn' with 'function'
	content = regexp.MustCompile(`\bfn\b`).ReplaceAll(content, []byte("function"))
	// Replace 'str.' with 'string.'
	content = regexp.MustCompile(`\bstr.\b`).ReplaceAll(content, []byte("string."))

	err = ioutil.WriteFile(filePath, content, 0644)
	if err != nil {
		return err
	}
	return nil
}

func renameFile(oldPath string) error {
	// Replace VMK with Lua in file names
	newPath := strings.Replace(oldPath, "VMK", "LUA", -1)
	newPath = strings.Replace(newPath, "Vmk", "Lua", -1)
	newPath = strings.Replace(newPath, "vmk", "lua", -1)

	return os.Rename(oldPath, newPath)
}

func traverseDir(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Process files with .vmk extension
		if strings.HasSuffix(info.Name(), ".vmk") {
			// Replace content inside file
			if err := replaceInFile(path); err != nil {
				return err
			}

			// Rename the file
			if err := renameFile(path); err != nil {
				return err
			}
		}
		return nil
	})
}

func main() {
	dir := "." // Specify the root directory to start the traversal
	if err := traverseDir(dir); err != nil {
		fmt.Println("Error:", err)
	}
}
