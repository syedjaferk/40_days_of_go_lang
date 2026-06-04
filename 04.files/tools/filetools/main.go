package main

import (
	"filetools/tools"
	"fmt"
	// "strings"
)

func main() {

	// value := "apple apple apple"
	// new_value := strings.Replace(value, "apple", "Wijay", -1)

	// fmt.Println("New Value ", new_value)

	for {
		fmt.Println("\n--- File Based Tools----")
		fmt.Println("1. File Copier")
		fmt.Println("2. Copy file to folder")
		fmt.Println("3. Search file")
		fmt.Println("4. Notes")
		fmt.Println("5. Tree")
		fmt.Println("6. Word Count")

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
		case 4:
			tools.NotesManager()
		case 5:
			tools.DirectoryTree()
		case 6:
			tools.WordCounter()
		case 0:
			return
		}

	}
}
