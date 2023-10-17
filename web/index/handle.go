package index

import (
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

func Handle(html *template.Template) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		testdata := struct {
			Title string
		}{
			Title: "AnswerTag",
		}

		err := html.ExecuteTemplate(w, "index.html", testdata)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
