package main

import (
	"fmt"
	"image/jpeg"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/webp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: webptojpg <path to .webp file>")
		os.Exit(1)
	}

	inputFilePath := os.Args[1]
	outputFilePath := generateOutputPath(inputFilePath)

	// Open the source file
	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	// Decode webp
	img, err := webp.Decode(inputFile)
	if err != nil {
		panic(err)
	}

	// Create the output file
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// Encode to jpg
	err = jpeg.Encode(outputFile, img, &jpeg.Options{Quality: 80})
	if err != nil {
		panic(err)
	}

	fmt.Println("Conversion completed:", outputFilePath)
}

// generateOutputPath creates the output file path by changing the extension to .jpg
func generateOutputPath(inputPath string) string {
	dir, file := filepath.Split(inputPath)
	filename := strings.TrimSuffix(file, filepath.Ext(file))
	return filepath.Join(dir, filename+".jpg")
}
