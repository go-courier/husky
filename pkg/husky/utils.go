package husky

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func ResolveGitRoot() string {
	cwd, _ := os.Getwd()
	return resolveGitRoot(cwd)
}

func resolveGitRoot(path string) string {
	f, err := os.Lstat(filepath.Join(path, ".git"))
	if err != nil {
		if os.IsNotExist(err) {
			return resolveGitRoot(filepath.Join(path, ".."))
		}
		panic(err)
	}
	if !f.IsDir() {
		panic(fmt.Errorf(".git must be a directory"))
	}
	return path
}

func ListGithookName(root string) ([]string, error) {
	githooks := make([]string, 0)

	files, err := ioutil.ReadDir(path.Join(root, ".git/hooks"))
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		if filepath.Ext(f.Name()) == ".sample" {
			githooks = append(githooks, strings.Split(f.Name(), ".")[0])
		}
	}

	return githooks, nil
}

func WriteFile(filename string, data []byte) error {
	dir := filepath.Dir(filename)

	if dir != "" {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}

	return ioutil.WriteFile(filename, data, os.ModePerm)
}
