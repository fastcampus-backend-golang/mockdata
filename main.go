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

	"github.com/fastcampus-backend-golang/mockdata/data"
)

func main() {
	// baca flag help, input, dan output
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

	// validasi input ada
	if err := validateInput(inputFile); err != nil {
		fmt.Printf("invalid input: %s\n", err.Error())
		os.Exit(0)
	}

	// validasi output ada
	if err := validateOutput(outputPath); err != nil {
		fmt.Printf("invalid output: %s\n", err.Error())
		os.Exit(0)
	}

	// proses input
	var mapping map[string]string
	if err := readInput(inputFile, &mapping); err != nil {
		fmt.Printf("failed reading input: %s\n", err.Error())
		os.Exit(0)
	}

	// validasi tipedata memuat hanya tipedata yang didukung
	if err := validateType(mapping); err != nil {
		fmt.Printf("invalid type: %s\n", err.Error())
		os.Exit(0)
	}

	// membuat data palsu
	result, err := generateOutput(mapping)
	if err != nil {
		fmt.Printf("failed generating output: %s\n", err.Error())
		os.Exit(0)
	}

	// menulis hasil ke file
	if err := writeOutput(outputPath, result); err != nil {
		fmt.Printf("failed writing output: %s\n", err.Error())
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
	for _, value := range mapping {
		if !data.Supported[value] {
			return fmt.Errorf("%s type is not supported", value)
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

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if len(fileBytes) == 0 {
		return errors.New("file is empty")
	}

	if err := json.Unmarshal(fileBytes, &mapping); err != nil {
		return err
	}

	return nil
}

func generateOutput(mapping map[string]string) (map[string]any, error) {
	result := make(map[string]any)

	for key, dataType := range mapping {
		result[key] = data.Generate(dataType)
	}

	return result, nil
}

func writeOutput(path string, result map[string]any) error {
	if path == "" {
		return errors.New("path is empty")
	}

	// flag untuk membaca, menulis, membuat file baru, dan menghapus isi file jika sudah ada
	flags := os.O_RDWR | os.O_CREATE | os.O_TRUNC // READ WRITE | CREATE | TRUNCATE
	file, err := os.OpenFile(path, flags, 0o644)  // 0644 = rw-r--r-- = owner read write, group read, other read
	if err != nil {
		return errors.New("error opening file")
	}
	defer file.Close()

	// marshal result dengan indentasi
	resultBytes, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return err
	}

	// tulis ke file
	_, err = file.Write(resultBytes)
	if err != nil {
		return err
	}

	return nil
}
