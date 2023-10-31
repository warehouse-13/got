package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/warehouse-13/got/template"
)

const (
	suffix = "_test"
	ext    = ".go"
)

func main() {
	var filename string
	flag.StringVar(&filename, "name", "", "The name of the test file (without the '_test'). Default: current package name.")
	flag.Parse()

	if filename != "" {
		filename = strings.ReplaceAll(filename, suffix, "")
		filename = strings.ReplaceAll(filename, ext, "")
	}

	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	packageName := filepath.Base(path)

	if filename == "" {
		filename = packageName
	}

	tmpl := template.Generate(packageName)

	filename = fmt.Sprintf("%s%s%s", filename, suffix, ext)
	if err := os.WriteFile(filename, tmpl, 0644); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Dummy test file written to %s.\n", filename)
	fmt.Println("Run `go mod tidy` to start using.")
}
