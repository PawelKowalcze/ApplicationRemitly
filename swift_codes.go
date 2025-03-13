package main

import (
	"errors"
	"github.com/tealeg/xlsx"
	"strings"
)

type SwiftCode struct {
	Code             string
	CountryCode      string
	CountryName      string
	IsHeadquarter    bool
	IsAssociatedWith int
}

func contains(slice []string, item string) (int, bool) {
	for i, s := range slice {
		if s == item {
			return i, true
		}
	}
	return 0, false
}

func parseSwiftCodes(filePath string) ([]SwiftCode, error) {
	var swiftCodes []SwiftCode
	var headquarterCodes []string

	xlFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		return nil, errors.New("Failed to open file")
	}

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			if len(row.Cells) < 3 {
				continue // Skip rows with insufficient columns
			}

			countryCode := strings.ToUpper(row.Cells[0].String())
			code := row.Cells[1].String()
			countryName := strings.ToUpper(row.Cells[6].String())

			isHeadquarter := strings.HasSuffix(code, "XXX")
			if isHeadquarter {
				headquarterCodes = append(headquarterCodes, code[:8])
			} else {
				headquarterCodes = append(headquarterCodes, "")
			}

			swiftCodes = append(swiftCodes, SwiftCode{
				Code:          code,
				CountryCode:   countryCode,
				CountryName:   countryName,
				IsHeadquarter: isHeadquarter,
			})
		}
		for i, row := range sheet.Rows {
			code := row.Cells[1].String()

			index, ok := contains(headquarterCodes, code[:8])
			if ok {
				swiftCodes[i].IsAssociatedWith = index
			}
		}

	}

	return swiftCodes, nil
}
