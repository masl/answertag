package web

import (
	"io/fs"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"github.com/masl/answertag/storage"
	"github.com/masl/answertag/web/index"
	"github.com/masl/answertag/web/ping"
	"github.com/masl/answertag/web/start"
)

func GetRouter(store storage.Store, htmlTemplates *template.Template, staticFS fs.FS) *httprouter.Router {
	router := httprouter.New()

	// index page
	router.GET("/", index.Handle(htmlTemplates))

	// static files
	router.Handler("GET",  "/static/*filepath", http.FileServer(http.FS(staticFS)))

	// api endpoints
	router.POST("/api/ping", ping.Handle())

	router.POST("/api/start", start.Handle(store))

	return router
}
