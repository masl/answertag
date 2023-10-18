package ping

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Ping struct {
	Name string `json:"name"`
}

func Handle() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var p Ping

		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			slog.Error("ping error:", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Write([]byte("Hello, " + p.Name + "!"))
	}
}
