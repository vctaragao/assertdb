package main

import (
	"log"
	"net/http"

	"github.com/vctaragao/assertdb/internal"
	"github.com/vctaragao/assertdb/internal/infra/database"
	"github.com/vctaragao/assertdb/internal/web"
)

func main() {
	db := database.Connect()
	createService := internal.NewCreateBananaService(database.NewBananaAdapter(db))

	mux := http.NewServeMux()

	web.RegisterRoutes(mux, createService)

	log.Println("Server running on port 8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
