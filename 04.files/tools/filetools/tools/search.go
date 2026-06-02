package tools

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func FileSearch() {

	var dir string
	var keyword string

	fmt.Print("Directory : ")
	fmt.Scan(&dir)

	fmt.Print("Keyword : ")
	fmt.Scan(&keyword)

	filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if strings.Contains(
				strings.ToLower(info.Name()),
				strings.ToLower(keyword),
			) {
				fmt.Println(path)
			}

			return nil
		})
}
