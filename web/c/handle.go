package c

import (
	"log/slog"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"github.com/masl/answertag/storage"
)

func Handle(html *template.Template, store storage.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		data, err := store.ReadByID(ps.ByName("id"))
		if err != nil {
			if err == storage.ErrNotFound {
				http.NotFound(w, r)
				return
			}

			slog.Error("error reading cloud", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		err = html.ExecuteTemplate(w, "cloud.html", data)
		if err != nil {
			slog.Error("error executing template", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
