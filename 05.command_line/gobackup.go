package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func copyFile(source string, destination string) error {
	err := os.MkdirAll(filepath.Dir(destination), 0755)

	if err != nil {
		return err
	}

	//  Opened a source file for copy
	src, err := os.Open(source)

	if err != nil {
		return err
	}

	defer src.Close()

	// Opened a destination file
	dst, err := os.Create(destination)
	if err != nil {
		return err
	}

	defer dst.Close()

	_, err = io.Copy(dst, src)

	if err != nil {
		return err
	}
	return err
}

func main() {

	source := flag.String("source", "", "Source Directory")
	destination := flag.String("destination", "", "Destination Directory")
	verbose := flag.Bool("verbose", false, "Verbsoe Logging")

	flag.Parse()

	// Validation

	if *source == "" {
		fmt.Println("Error: Source is required")
		os.Exit(1)
	}

	if *destination == "" {
		fmt.Println("Error: Destination is required")
		os.Exit(1)
	}

	err := filepath.Walk(
		*source,
		func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			relativePath, err := filepath.Rel(*source, path)
			if err != nil {
				return err
			}

			target := filepath.Join(*destination, relativePath)

			if *verbose {
				fmt.Println("Copying", path, target)
			}
			err = copyFile(
				path,
				target,
			)

			if err != nil {
				fmt.Println("Error copying %s: %v\n", path, err)
			}
			return nil

		})

	if err != nil {
		fmt.Println("Error ", err)
		os.Exit(1)
	}

	fmt.Println("Backup Completed")

}
