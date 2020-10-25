package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	m := make(map[string][]string, 0)
	err := filepath.Walk("./src", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Skip if is directory.
		if info.IsDir() {
			return nil
		}
		fullDir := filepath.Dir(path)
		paths := strings.Split(fullDir, "/")
		if len(paths) != 2 {
			return errors.New("invalid directory structure")
		}
		dir := paths[len(paths)-1]
		fileName := filepath.Base(path)
		ext := filepath.Ext(path)
		name := fileName[0 : len(fileName)-len(ext)]
		if _, ok := m[dir]; !ok {
			m[dir] = make([]string, 0)
		}
		m[dir] = append(m[dir], name)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(m)
}
