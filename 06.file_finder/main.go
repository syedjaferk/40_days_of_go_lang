package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	var (
		dir     string
		ext     string
		verbose bool
		limit   int
	)

	flag.StringVar(&dir, "dir", "", "Directory to search")
	flag.StringVar(&ext, "ext", "", "Filter by file extension")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose output")
	flag.IntVar(&limit, "limit", 0, "No of results")

	flag.Usage = func() {
		fmt.Println("Usage: ")
		fmt.Println("go run main.go -dir=<dir_val> -ext=<ext_Val> -verbose -limit=<number>")
		fmt.Println()
		fmt.Println("Options : ")
		flag.PrintDefaults()
	}

	flag.Parse()

	// Validation
	if dir == "" {
		fmt.Println("Error: Directory is required")
		os.Exit(1)
	}

	info, err := os.Stat(dir)
	if err != nil {
		fmt.Println("Error: Directory does not exist.")
		os.Exit(1)
	}

	if !info.IsDir() {
		fmt.Println("Error: Provided path is not a directory")
		os.Exit(1)
	}

	// Verbose

	if verbose {
		fmt.Printf("Searcing in %s\n", dir)

		if ext != "" {
			fmt.Printf("Extension Filter  %s\n", ext)
		}
	}

	// Actual Search

	var results []string

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if ext != "" && filepath.Ext(info.Name()) != ext {
			return nil
		}

		results = append(results, path)

		return nil
	})

	if err != nil {
		fmt.Printf("Error while searching %v \n", err)
		os.Exit(1)
	}

	if limit > 0 && len(results) > limit {
		results = results[:limit]
	}

	// Print everything in results

	for _, file := range results {
		fmt.Println(filepath.Base(file))
	}

}
