package main

import (
	"fmt"
	"log"
)

func main() {
	filePath := "Interns_2025_SWIFT_CODES.xlsx"
	swiftCodes, err := parseSwiftCodes(filePath)
	if err != nil {
		log.Fatalf("Failed to parse SWIFT codes: %v", err)
	}

	for _, code := range swiftCodes {
		fmt.Printf("%+v\n", code)
	}
}
