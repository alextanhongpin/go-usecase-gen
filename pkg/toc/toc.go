package toc

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func getWd() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

type Toc struct {
	FileName string
	Dir      string
	FullPath string
}

func relativePath(path ...string) string {
	wd := getWd()
	result := append([]string{wd}, path...)
	out := filepath.Join(result...)

	fmt.Println(out)
	return out
}

func GetFilesInDir(targetDir string) (map[string][]Toc, error) {
	m := make(map[string][]Toc, 0)
	err := filepath.Walk(relativePath(targetDir), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error walking directory: %w", err)
		}
		// Skip if is directory.
		if info.IsDir() {
			return nil
		}
		fullDir := filepath.Dir(path)
		paths := strings.Split(fullDir, "/")
		if len(paths) < 2 {
			return errors.New("invalid directory structure")
		}
		dir := paths[len(paths)-1]
		fileName := filepath.Base(path)
		ext := filepath.Ext(path)
		name := fileName[0 : len(fileName)-len(ext)]
		if _, ok := m[dir]; !ok {
			m[dir] = make([]Toc, 0)
		}
		m[dir] = append(m[dir], Toc{
			FileName: name,
			FullPath: fullDir,
			Dir:      dir,
		})
		return nil
	})
	if err != nil {
		return nil, err
	}

	return m, nil
}
