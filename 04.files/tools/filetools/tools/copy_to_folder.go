package tools

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func FileToFolderCopier() {
	var src, destFolder string

	fmt.Println("Source File : ")
	fmt.Scan(&src)

	fmt.Println("Destination Folder: ")
	fmt.Scan(&destFolder)

	srcFilename := filepath.Base(src)
	destPath := filepath.Join(destFolder, srcFilename)

	source, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer source.Close()

	dst_fo, err := os.Create(destPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dst_fo.Close()

	io.Copy(dst_fo, source)
	fmt.Println("File copied successfully")

}
