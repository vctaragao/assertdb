package web

import (
	"log"
	"net/http"

	"github.com/vctaragao/assertdb/internal"
	"github.com/vctaragao/assertdb/internal/entity"
)

func RegisterRoutes(mux *http.ServeMux, createService *internal.CreateBananaService) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b := &entity.Banana{
			Name:  "Banana",
			Color: "Yellow",
		}

		if err := createService.Execute(r.Context(), b); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		if _, err := w.Write([]byte("Banana created!")); err != nil {
			log.Println(err)
		}
	})

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}
