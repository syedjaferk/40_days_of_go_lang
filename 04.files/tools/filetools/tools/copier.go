package tools

import (
	"fmt"
	"io"
	"os"
)

func FileCopier() {
	var src, dest string

	fmt.Println("Source File : ")
	fmt.Scan(&src)

	fmt.Println("Destination File: ")
	fmt.Scan(&dest)

	source, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer source.Close()

	dst_fo, err := os.Create(dest)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dst_fo.Close()

	io.Copy(dst_fo, source)
	fmt.Println("File copied successfully")

}
