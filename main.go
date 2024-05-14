package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
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
	// if the output path only dir, use same name as input file
	// if output path is a file, use it as output file
	// if the file exist, confirm overwrite
}

func printUsage() {
	fmt.Println("Usage: fakegen [-i | --input] <input file> [-o | --output] <output path>")
	fmt.Println("-i, --input: Input of JSON file as a template")
	fmt.Println("-o, --output: Output JSON file for the generated data")
}

func confirmOverwrite() {
	fmt.Println("Are you sure you want to overwrite the file? (y/n)")

	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	response = strings.ToLower(strings.TrimSpace(response))

	if response != "y" && response != "yes" {
		fmt.Println("Aborting...")
		os.Exit(0)
	}
}
