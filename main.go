package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var (
	outputFilename   = flag.String("out", "", "Name of the output ffmpeg metadata file")
	templateFilename = flag.String("template", "metadata.tpl", "Template for the ffmpeg metadata file")
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Usage: gen-ffmetadata [OPTIONS] METADATA_FILE")
		flag.PrintDefaults()
		os.Exit(1)
	}
	metadataPath := args[0]

	// Use the name of the input file as output file name
	if *outputFilename == "" {
		*outputFilename = trimFileExtension(metadataPath) + ".metadata"
	}

	f, err := os.Open(metadataPath)
	if err != nil {
		fmt.Printf("open metadata file: %v", err)
		os.Exit(2)
	}

	var vm VideoMetadata
	if err = vm.ParseYAML(f); err != nil {
		fmt.Printf("parse metadata file: %v", err)
		os.Exit(2)
	}

	tpl, err := template.ParseFiles(*templateFilename)
	if err != nil {
		fmt.Printf("parse output template: %v", err)
		os.Exit(2)
	}
	g, err := os.Create(*outputFilename)
	if err != nil {
		fmt.Printf("create output file: %v", err)
		os.Exit(2)
	}
	err = tpl.Execute(g, vm)
	if err != nil {
		fmt.Printf("render template: %v", err)
		os.Exit(2)
	}
}

// trimFileExtension returns the name of a file without its extension
func trimFileExtension(filename string) string {
	return strings.TrimSuffix(filename, filepath.Ext(filename))
}
