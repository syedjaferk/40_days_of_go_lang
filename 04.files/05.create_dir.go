package main

import (
	"os"
)

func main() {
	os.Mkdir("logs", 0755)
	os.MkdirAll("data/2026/go/files/content", 0755)
}
