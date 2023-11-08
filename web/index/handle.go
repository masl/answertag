package index

import (
	"log/slog"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/masl/answertag/tmpl"
)

func Handle(tm *tmpl.TemplateManager) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		data := struct {
			Title string
		}{
			Title: "AnswerTag",
		}

		if err := tm.RenderTemplate(w, "index", data, nil); err != nil {
			slog.Error("error executing template", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
