package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"githhub.com/warehouse-13/got/template"
)

const (
	suffix = "_test"
	ext    = ".go"
)

func main() {
	var packageName string
	flag.StringVar(&packageName, "name", "", "The name of the test file (without the '_test'). Default: current package name.")
	flag.Parse()

	if packageName != "" {
		packageName = strings.ReplaceAll(packageName, suffix, "")
		packageName = strings.ReplaceAll(packageName, ext, "")
	}

	if packageName == "" {
		path, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		packageName = filepath.Base(path)
	}

	tmpl := template.Generate(packageName)

	filename := fmt.Sprintf("%s%s%s", packageName, suffix, ext)
	if err := os.WriteFile(filename, tmpl, 0644); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Dummy test file written to %s.\n", filename)
	fmt.Println("Run `go mod tidy` to start using.")
}
