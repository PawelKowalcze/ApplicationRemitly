package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/PawelKowalcze/ApplicationRemitly/internal/auth"
	"github.com/PawelKowalcze/ApplicationRemitly/internal/database"
	"github.com/google/uuid"
	"net/http"
)

func (apiCfg *apiConfig) handlerSWIFTCode(w http.ResponseWriter, r *http.Request) {
	type message struct {
		Message string `json:"message"`
	}

	type parameters struct {
		Address       string `json:"address"`
		BankName      string `json:"bankName"`
		CountryISO2   string `json:"countryISO2"`
		CountryName   string `json:"countryName"`
		IsHeadquarter bool   `json:"isHeadquarter"`
		SwiftCode     string `json:"swiftCode"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	_, err = apiCfg.DB.CreateSWIFTCodeEntry(r.Context(), database.CreateSWIFTCodeEntryParams{
		ID:             uuid.New(),
		Countrycode:    params.CountryISO2,
		Swiftcode:      params.SwiftCode,
		Codetype:       "BIC11",
		Name:           params.BankName,
		Address:        params.Address,
		Townname:       "",
		Timezone:       "",
		Countryname:    params.CountryName,
		Isheadquarter:  params.IsHeadquarter,
		Associatedwith: 0,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creating SWIFT code: %v", err))
		return
	}

	responseMessage := message{Message: "SWIFT code created successfully"}

	respondWithJSON(w, 201, responseMessage)
}

func (apiCfg *apiConfig) handlerGetEntryBySWIFTCode(w http.ResponseWriter, r *http.Request) {
	Swiftcode, err := auth.GetSWIFTCode(r.URL.Path)
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Error getting SWIFT code: %v", err))
		return
	}
	code, err := apiCfg.DB.GetEntryBySWIFTCode(r.Context(), Swiftcode)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get SWIFT code entry: %v", err))
		return
	}

	if code.Isheadquarter {
		branches, err := apiCfg.DB.GetBranchesByAssociatedWith(r.Context(), code.Associatedwith)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldn't get branches: %v", err))
			return
		}
		headquarter := databaseSwiftCodeToSwiftCode_Headquarter(code)
		headquarter.Branches = make([]SwiftCode_BranchForHeadquarter, len(branches))
		for i, branch := range branches {
			headquarter.Branches[i] = databaseSwiftCodeToSwiftCode_BranchForHeadquarter(branch)
		}
		respondWithJSON(w, 200, headquarter)
		return
	}

	if code.Associatedwith != 0 {
		respondWithJSON(w, 200, databaseSwiftCodeToSwiftCode_Branch(code))
		return
	}
	respondWithJSON(w, 400, fmt.Sprintf("SWIFT code is not headquarter nor branch"))
}

func (apiCfg *apiConfig) handlerGetEntriesByCountryCode(w http.ResponseWriter, r *http.Request) {
	CountryCode, err := auth.GetCountryCode(r.URL.Path)
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Error getting country code: %v", err))
		return
	}
	CodeSlice, err := apiCfg.DB.GetEntryByCountryCode(r.Context(), CountryCode)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get country code entry: %v", err))
		return
	}
	mainInfo := Country_Code{
		Countrycode: CountryCode,
		Countryname: CodeSlice[0].Countryname, // Assuming all entries have the same country name
		SwiftCodes:  make([]SwiftCode_BranchForHeadquarter, len(CodeSlice)),
	}

	for i, code := range CodeSlice {
		mainInfo.SwiftCodes[i] = databaseSwiftCodeToSwiftCode_BranchForHeadquarter(code)
	}
	respondWithJSON(w, 200, mainInfo)
	return
}

func (apiCfg *apiConfig) handlerDeleteEntryForSWIFTCode(w http.ResponseWriter, r *http.Request) {
	type mess struct {
		Message string `json:"message"`
	}

	Swiftcode, err := auth.GetSWIFTCode(r.URL.Path)

	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Error getting SWIFT code: %v", err))
		return
	}

	exists, err := apiCfg.DB.CheckSWIFTCodeExists(context.Background(), Swiftcode)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Failed to check if SWIFT code exists: %v", err))
		return
	}
	if !exists {
		respondWithError(w, 400, fmt.Sprintf("SWIFT code %s does not exist", Swiftcode))
		return
	}

	err = apiCfg.DB.DeleteSWIFTCodeEntry(r.Context(), Swiftcode)

	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Couldn't delete entry for given SWIFT code: %v", err))
		return
	}

	responseMessage := mess{Message: "SWIFT code entry deleted successfully"}
	respondWithJSON(w, 200, responseMessage)

}
