package main

import (
	"context"
	"database/sql"
	"github.com/PawelKowalcze/ApplicationRemitly/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database", err)
	}

	//Run migrations:
	if err = goose.Up(conn, "sql/schema"); err != nil {
		log.Printf("error running schema migrations: %v", err)
	}

	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
	}

	filePath := "Interns_2025_SWIFT_CODES.xlsx"
	swiftCodes, err := parseSwiftCodes(filePath)
	if err != nil {
		log.Printf("Failed to parse SWIFT codes: %v", err)
	}

	for i, code := range swiftCodes {
		exists, err := apiCfg.DB.CheckSWIFTCodeExists(context.Background(), code.SWIFTCode)
		if err != nil {
			log.Printf("Error checking if swift code exists: %v", err)
		}

		if exists {
			log.Printf("SWIFT code %s already exists, skipping", code.SWIFTCode)
			continue
		}

		if i == 0 {
			continue
		}

		_, err = apiCfg.DB.CreateSWIFTCodeEntry(context.Background(), database.CreateSWIFTCodeEntryParams{
			ID:             uuid.New(),
			Countrycode:    code.CountryCode,
			Swiftcode:      code.SWIFTCode,
			Codetype:       code.CodeType,
			Name:           code.Name,
			Address:        code.Address,
			Townname:       code.TownName,
			Timezone:       code.TimeZone,
			Countryname:    code.CountryName,
			Isheadquarter:  code.IsHeadquarter,
			Associatedwith: int32(code.IsAssociatedWith),
		})

		if err != nil {
			log.Printf("Failed to add SWIFT code: %v", err)
		}
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/swift-codes", apiCfg.handlerSWIFTCode)
	v1Router.Get("/swift-codes/{swift-code}", apiCfg.handlerGetEntryBySWIFTCode)
	v1Router.Delete("/swift-codes/{swift-code}", apiCfg.handlerDeleteEntryForSWIFTCode)
	v1Router.Get("/swift-codes/country/{countryISO2code}", apiCfg.handlerGetEntriesByCountryCode)
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Serving on port %v", portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
