package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
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

	// process the input
	var mapping map[string]string
	if err := readInput(inputFile, &mapping); err != nil {
		fmt.Println("Error reading input")
		os.Exit(0)
	}

	// validate the mapping includes only supported types
	if err := validateType(mapping); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	// generate the output
	result, err := generateOutput(mapping)
	if err != nil {
		fmt.Println("Error generating output")
		os.Exit(0)
	}

	// write the output
	if err := writeOutput(outputPath, result); err != nil {
		fmt.Println("Error writing output")
		os.Exit(0)
	}
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

func validateType(mapping map[string]string) error {
	// supported types, will change with from the randomizer
	supportedTypes := map[string]bool{
		"name":    true,
		"date":    true,
		"address": true,
		"phone":   true,
	}

	for _, value := range mapping {
		if !supportedTypes[value] {
			return fmt.Errorf("Unsupported type: %s", value)
		}
	}

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

func readInput(path string, mapping *map[string]string) error {
	if path == "" {
		return errors.New("path is empty")
	}

	if mapping == nil {
		return errors.New("mapping is nil")
	}

	// read file
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// parse file
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if len(fileBytes) == 0 {
		return errors.New("file is empty")
	}

	// unmarshal JSON
	if err := json.Unmarshal(fileBytes, &mapping); err != nil {
		return err
	}

	return nil
}

func generateOutput(mapping map[string]string) (map[string]any, error) {
	result := make(map[string]any)

	for key, value := range mapping {
		result[key] = value // will be randomized here
	}

	return result, nil
}

func writeOutput(path string, result map[string]any) error {
	if path == "" {
		return errors.New("path is empty")
	}

	// open file for write
	flags := os.O_RDWR | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile(path, flags, 0o644)
	if err != nil {
		return errors.New("error opening file")
	}
	defer file.Close()

	// json marshal
	resultBytes, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return err
	}

	// write to file
	_, err = file.Write(resultBytes)
	if err != nil {
		return err
	}

	return nil
}
