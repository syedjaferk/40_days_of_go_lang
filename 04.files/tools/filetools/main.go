package main

import (
	"filetools/tools"
	"fmt"
)

func main() {

	for {
		fmt.Println("\n--- File Based Tools----")
		fmt.Println("1. File Copier")
		fmt.Println("2. Copy file to folder")
		fmt.Println("3. Search file")

		fmt.Println("0. Exit")

		var choice int
		fmt.Println("Enter the Choice : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			// File copier
			tools.FileCopier()
		case 2:
			tools.FileToFolderCopier()
		case 3:
			tools.FileSearch()
		case 0:
			return
		}

	}
}
