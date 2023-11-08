package c

import (
	"log/slog"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/masl/answertag/cloud"
	"github.com/masl/answertag/storage"
	"github.com/masl/answertag/tmpl"
)

func Handle(tm *tmpl.TemplateManager, store storage.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		cld, err := store.ReadByID(ps.ByName("id"))
		if err != nil {
			if err == storage.ErrNotFound {
				http.NotFound(w, r)
				return
			}

			slog.Error("error reading cloud", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		data := struct {
			CloudID string
			Tags    []cloud.TagWithFontSize
		}{
			CloudID: cld.ID.String(),
			Tags:    cloud.SupplementTagsWithFontSizes(cld.Tags),
		}

		if err := tm.RenderTemplate(w, "cloud", data, nil); err != nil {
			slog.Error("error executing template", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
