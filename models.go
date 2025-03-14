package main

import (
	"github.com/PawelKowalcze/ApplicationRemitly/internal/database"
	"github.com/google/uuid"
)

type SwiftCode struct {
	ID             uuid.UUID `json:"id"`
	Countrycode    string    `json:"country_code"`
	Swiftcode      string    `json:"swift_code"`
	Codetype       string    `json:"code_type"`
	Name           string    `json:"name"`
	Address        string    `json:"address"`
	Townname       string    `json:"town_name"`
	Timezone       string    `json:"time_zone"`
	Countryname    string    `json:"country_name"`
	Isheadquarter  bool      `json:"is_headquarter"`
	Associatedwith int32     `json:"associated_with"`
}

func databaseSwiftCodeToSwiftCode(dbSwiftCode database.SwiftCode) SwiftCode {
	return SwiftCode{
		ID:             dbSwiftCode.ID,
		Countrycode:    dbSwiftCode.Countrycode,
		Swiftcode:      dbSwiftCode.Swiftcode,
		Codetype:       dbSwiftCode.Codetype,
		Name:           dbSwiftCode.Name,
		Address:        dbSwiftCode.Address,
		Townname:       dbSwiftCode.Townname,
		Timezone:       dbSwiftCode.Timezone,
		Countryname:    dbSwiftCode.Countryname,
		Isheadquarter:  dbSwiftCode.Isheadquarter,
		Associatedwith: dbSwiftCode.Associatedwith,
	}
}
