package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"golang.org/x/sys/unix"
)

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var (
		path   string
		bins   map[string]string
		output []byte
		err    error
	)

	bins = make(map[string]string)

	for _, path = range filepath.SplitList(os.Getenv("PATH")) {
		path, err = filepath.EvalSymlinks(path)
		if err != nil && !errors.Is(err, fs.ErrNotExist) {
			panic(err)
		}

		filepath.Walk(path, func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			err = unix.Access(path, unix.X_OK)
			if err == nil {
				bins[filepath.Base(path)] = path
			}

			if err == unix.EACCES {
				return nil
			}

			return err
		})
	}

	output, err = json.Marshal(bins)
	panicIf(err)

	fmt.Println(string(output))
}
