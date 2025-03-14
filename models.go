package main

import (
	"github.com/PawelKowalcze/ApplicationRemitly/internal/database"
)

type SwiftCode_Branch struct {
	Address       string `json:"address"`
	Name          string `json:"bankName"`
	Countrycode   string `json:"countryISO2"`
	Countryname   string `json:"countryName"`
	Isheadquarter bool   `json:"isHeadquarter"`
	Swiftcode     string `json:"swiftCode"`
}

type SwiftCode_BranchForHeadquarter struct {
	Address       string `json:"address"`
	Name          string `json:"bankName"`
	Countrycode   string `json:"countryISO2"`
	Isheadquarter bool   `json:"isHeadquarter"`
	Swiftcode     string `json:"swiftCode"`
}

type SwiftCode_Headquarter struct {
	Address       string                           `json:"address"`
	Name          string                           `json:"bankName"`
	Countrycode   string                           `json:"countryISO2"`
	Countryname   string                           `json:"countryName"`
	Isheadquarter bool                             `json:"isHeadquarter"`
	Swiftcode     string                           `json:"swiftCode"`
	Branches      []SwiftCode_BranchForHeadquarter `json:"branches"`
}

type Country_Code struct {
	Countrycode string                           `json:"countryISO2"`
	Countryname string                           `json:"countryName"`
	SwiftCodes  []SwiftCode_BranchForHeadquarter `json:"swiftCodes"`
}

func databaseCountryCodeToCountryCode(dbSwiftCode database.SwiftCode) Country_Code {
	return Country_Code{
		Countrycode: dbSwiftCode.Countrycode,
		Countryname: dbSwiftCode.Countryname,
	}
}

func databaseSwiftCodeToSwiftCode_Headquarter(dbSwiftCode database.SwiftCode) SwiftCode_Headquarter {
	return SwiftCode_Headquarter{
		Countrycode:   dbSwiftCode.Countrycode,
		Swiftcode:     dbSwiftCode.Swiftcode,
		Name:          dbSwiftCode.Name,
		Address:       dbSwiftCode.Address,
		Countryname:   dbSwiftCode.Countryname,
		Isheadquarter: dbSwiftCode.Isheadquarter,
	}
}

func databaseSwiftCodeToSwiftCode_Branch(dbSwiftCode database.SwiftCode) SwiftCode_Branch {
	return SwiftCode_Branch{
		Address:       dbSwiftCode.Address,
		Name:          dbSwiftCode.Name,
		Countrycode:   dbSwiftCode.Countrycode,
		Countryname:   dbSwiftCode.Countryname,
		Isheadquarter: dbSwiftCode.Isheadquarter,
		Swiftcode:     dbSwiftCode.Swiftcode,
	}
}

func databaseSwiftCodeToSwiftCode_BranchForHeadquarter(dbSwiftCode database.SwiftCode) SwiftCode_BranchForHeadquarter {
	return SwiftCode_BranchForHeadquarter{
		Address:       dbSwiftCode.Address,
		Name:          dbSwiftCode.Name,
		Countrycode:   dbSwiftCode.Countrycode,
		Isheadquarter: dbSwiftCode.Isheadquarter,
		Swiftcode:     dbSwiftCode.Swiftcode,
	}
}
