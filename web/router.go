package web

import (
	"io/fs"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"github.com/masl/answertag/web/index"
)


func GetRouter(htmlTemplates *template.Template, cssFS fs.FS) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", index.Handle(htmlTemplates))
	router.Handler("GET", "/css/output.css", http.FileServer(http.FS(cssFS)))

	return router
}
