package tools

import (
	"fmt"
	"os"
	"path/filepath"
)

func DirectoryTree() {
	var root string

	fmt.Print("Directory : ")
	fmt.Scan(&root)

	filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {

			if info.IsDir() {
				fmt.Println("D: ", info.Name())
				// last_folder := info.Name()
			} else {
				// new_path := strings.Replace(path, root, "", -1)

				fmt.Println("  F: ", info.Name())
			}

			// fmt.Println(path, info.Name())

			return nil
		})
}
