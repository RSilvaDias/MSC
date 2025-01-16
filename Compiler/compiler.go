// Go program to read a file and compile it to a WebAssembly binary.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <source_file> <output_file.wasm>")
		return
	}

	sourceFile := os.Args[1]
	outputFile := os.Args[2]

	// Read the source file
	content, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		log.Fatalf("Failed to read the source file: %v", err)
	}

	fmt.Printf("Read source file '%s' successfully.\n", sourceFile)

	// Verify the file is not empty
	if len(content) == 0 {
		log.Fatalf("Source file '%s' is empty.", sourceFile)
	}

	// Determine the file extension
	ext := filepath.Ext(sourceFile)
	var cmd *exec.Cmd

	switch ext {
	case ".go":
		// Compile the source file to WebAssembly using TinyGo
		cmd = exec.Command("tinygo", "build", "-o", outputFile, "-target=wasm", sourceFile)
	case ".cpp":
		// Compile the source file to WebAssembly using Binaryen
		cmd = exec.Command("emcc", sourceFile, "-o", outputFile, "-s", "WASM=1")
	default:
		log.Fatalf("Unsupported file extension '%s'. Supported extensions are .go and .cpp.", ext)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		log.Fatalf("Failed to compile '%s' to WebAssembly: %v", sourceFile, err)
	}

	fmt.Printf("Compiled '%s' to WebAssembly binary '%s' successfully.\n", sourceFile, outputFile)
}
