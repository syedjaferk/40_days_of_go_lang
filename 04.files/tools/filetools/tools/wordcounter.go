package tools

import (
	"fmt"
	"os"
	"strings"
)

func WordCounter() {
	var file string

	fmt.Print("file : ")
	fmt.Scan(&file)

	data, err := os.ReadFile(file)

	if err != nil {
		fmt.Println(err)
		return
	}

	count := len(strings.Fields(string(data)))

	fmt.Println("Words : ", count)
}
