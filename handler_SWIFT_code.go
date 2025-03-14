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
	Swiftcode, err := auth.GetSWIFTCode(r.Header)
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Error getting SWIFT code: %v", err))
		return
	}
	code, err := apiCfg.DB.GetEntryBySWIFTCode(r.Context(), Swiftcode)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get SWIFT code entry: %v", err))
		return
	}
	respondWithJSON(w, 200, databaseSwiftCodeToSwiftCode(code))

}

//
//func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
//	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
//		UserID: user.ID,
//		Limit:  10})
//	if err != nil {
//		respondWithError(w, 400, fmt.Sprintf("Couldn't get posts for user: %v", err))
//		return
//	}
//
//	respondWithJSON(w, 200, databasePostsToPosts(posts))
//}
