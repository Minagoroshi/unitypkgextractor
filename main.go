package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <path-to-unitypackage> [output-directory]")
		os.Exit(1)
	}
	filePath := os.Args[1]
	outPath := ""
	if len(os.Args) > 2 {
		outPath = os.Args[2]
	}

	err := extractPackage(filePath, outPath)
	if err != nil {
		log.Println("Error extracting package:", err)
		log.Println("Press enter to exit...")
		fmt.Scanln()
		os.Exit(1)
	}

}
