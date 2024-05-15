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
	if err := validateInput(inputFile); err != nil {
		fmt.Println("Input file does not exist")
		os.Exit(0)
	}

	// validate output path exists
	validateOutput(outputPath)
}

func printUsage() {
	fmt.Println("Usage: fakegen [-i | --input] <input file> [-o | --output] <output path>")
	fmt.Println("-i, --input: Input of JSON file as a template")
	fmt.Println("-o, --output: Output JSON file for the generated data")
}

func validateInput(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return err
	}

	return nil
}

func validateOutput(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil
	}

	fmt.Println("Output file already exists")
	confirmOverwrite()

	return nil
}

func confirmOverwrite() {
	fmt.Print("Are you sure you want to overwrite the file? (y/n) ")

	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	response = strings.ToLower(strings.TrimSpace(response))

	if response != "y" && response != "yes" {
		fmt.Println("Aborting...")
		os.Exit(0)
	}
}
