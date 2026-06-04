package tools

import (
	"bufio"
	"fmt"
	"os"
)

func NotesManager() {
	fmt.Println("Enter Notes: ")
	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')

	file, _ := os.OpenFile(
		"notes.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644,
	)

	defer file.Close()

	file.WriteString(text)

	fmt.Println("Saved")
}
