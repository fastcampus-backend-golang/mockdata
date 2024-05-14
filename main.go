package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// parse input and output flags
	var help bool
	var inputFile, outputPath string

	flag.BoolVar(&help, "h", false, "Show help")
	flag.BoolVar(&help, "help", false, "Show help")

	flag.StringVar(&inputFile, "i", "", "Input JSON file")
	flag.StringVar(&inputFile, "input", "", "Input JSON file")
	flag.StringVar(&outputPath, "o", "", "Output JSON file")
	flag.StringVar(&outputPath, "output", "", "Output JSON file")

	flag.Parse()

	if help || inputFile == "" || outputPath == "" {
		printUsage()
		os.Exit(0)
	}

	// validate input file exists

	// validate output path exists
}

func printUsage() {
	fmt.Println(`Usage: fakegen [-i | --input] <input file> [-o | --output] <output path>`)
	fmt.Println(`-i, --input: Input JSON file`)
	fmt.Println(`-o, --output: Output JSON file`)
}
