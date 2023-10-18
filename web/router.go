package web

import (
	"io/fs"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"github.com/masl/answertag/web/index"
	"github.com/masl/answertag/web/ping"
)

func GetRouter(htmlTemplates *template.Template, staticFS fs.FS) *httprouter.Router {
	router := httprouter.New()

	// index page
	router.GET("/", index.Handle(htmlTemplates))

	// static files
	router.Handler("GET", "/static/output.css", http.FileServer(http.FS(staticFS)))
	router.Handler("GET", "/static/htmx.min.js", http.FileServer(http.FS(staticFS)))
	router.Handler("GET", "/static/ext/json-enc.js", http.FileServer(http.FS(staticFS)))

	// api endpoints
	router.POST("/api/ping", ping.Handle())

	return router
}
