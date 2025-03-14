package main

import (
	"encoding/json"
	"fmt"
	"github.com/PawelKowalcze/ApplicationRemitly/internal/auth"
	"github.com/PawelKowalcze/ApplicationRemitly/internal/database"
	"github.com/google/uuid"
	"net/http"
)

func (apiCfg *apiConfig) handlerSWIFTCode(w http.ResponseWriter, r *http.Request) {
	type message struct {
		Message string
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
