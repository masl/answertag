package web

import (
	"io/fs"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"github.com/masl/answertag/web/index"
)

func GetRouter(htmlTemplates *template.Template, staticFS fs.FS) *httprouter.Router {
	router := httprouter.New()

	// index page
	router.GET("/", index.Handle(htmlTemplates))

	// static files
	router.Handler("GET", "/static/output.css", http.FileServer(http.FS(staticFS)))
	router.Handler("GET", "/static/htmx.min.js", http.FileServer(http.FS(staticFS)))

	return router
}
