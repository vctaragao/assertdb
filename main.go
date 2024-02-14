package main

import (
	"log"
	"net/http"

	"github.com/vctaragao/assertdb/internal"
	"github.com/vctaragao/assertdb/internal/entity"
	"github.com/vctaragao/assertdb/internal/infra/database"
)

func main() {
	db := database.Connect()
	createService := internal.NewCreateBananaService(database.NewBananaAdapter(db))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b := &entity.Banana{
			Name:  "Banana",
			Color: "Yellow",
		}

		if err := createService.Execute(r.Context(), b); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		if _, err := w.Write([]byte("Banana created!")); err != nil {
			log.Println(err)
		}
	})

	log.Println("Server running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
