package start

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/masl/answertag/cloud"
	"github.com/masl/answertag/storage"
)

func Handle(store storage.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		c := cloud.New()

		err := store.Create(c)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		responseBody := ReseponseBody{
			CloudID: c.ID.String(),
		}

		responseBytes, err := json.Marshal(responseBody)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// htmx redirect to cloud page
		w.Header().Set("HX-Redirect", "/c/"+c.ID.String())

		slog.Info("cloud created", "cloud", c.ID.String())

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseBytes)
	}
}
