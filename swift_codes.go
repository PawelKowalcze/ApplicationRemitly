package main

import (
	"errors"
	"github.com/tealeg/xlsx"
	"strings"
)

type SWIFTCode struct {
	CountryCode      string
	SWIFTCode        string
	CodeType         string
	Name             string
	Address          string
	TownName         string
	CountryName      string
	TimeZone         string
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

func parseSwiftCodes(filePath string) ([]SWIFTCode, error) {
	var swiftCodes []SWIFTCode
	var headquarterCodes []string

	xlFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		return nil, errors.New("Failed to open file")
	}

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			if len(row.Cells) < 8 {
				continue // Skip rows with insufficient columns
			}

			countryCode := strings.ToUpper(row.Cells[0].String())
			swiftCode := row.Cells[1].String()
			codeType := row.Cells[2].String()
			name := row.Cells[3].String()
			address := row.Cells[4].String()
			townName := row.Cells[5].String()
			countryName := strings.ToUpper(row.Cells[6].String())
			timeZone := row.Cells[7].String()

			isHeadquarter := strings.HasSuffix(swiftCode, "XXX")
			if isHeadquarter {
				headquarterCodes = append(headquarterCodes, swiftCode[:8])
			} else {
				headquarterCodes = append(headquarterCodes, "")
			}

			swiftCodes = append(swiftCodes, SWIFTCode{
				CountryCode:   countryCode,
				SWIFTCode:     swiftCode,
				CodeType:      codeType,
				Name:          name,
				Address:       address,
				TownName:      townName,
				TimeZone:      timeZone,
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
