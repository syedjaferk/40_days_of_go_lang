package main

import "os"

func main() {
	file, _ := os.Create("notes.txt")
	defer file.Close()

	file.WriteString("Hello from Go Lang with Parottasalna !\n")
	file.WriteString("Learning file system operations")
}
