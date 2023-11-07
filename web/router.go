package web

import (
	"io/fs"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
	"github.com/masl/answertag/storage"
	"github.com/masl/answertag/web/c"
	"github.com/masl/answertag/web/index"
	"github.com/masl/answertag/web/ping"
	"github.com/masl/answertag/web/start"
	"github.com/masl/answertag/web/tags"
	"github.com/masl/answertag/ws"
)

func GetRouter(store storage.Store, htmlTemplates *template.Template, staticFS fs.FS, hub *ws.Hub) *httprouter.Router {
	router := httprouter.New()

	// index page
	router.GET("/", index.Handle(htmlTemplates))

	// cloud page
	router.GET("/c/:id", c.Handle(htmlTemplates, store))

	// static files
	router.Handler("GET", "/static/*filepath", PreventDirectoryIndex(http.FileServer(http.FS(staticFS))))

	// redirect /favicon.ico to /static/favicon.ico
	router.GET("/favicon.ico", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		http.Redirect(w, r, "/static/favicon.ico", http.StatusMovedPermanently)
	})

	// api endpoints
	router.POST("/api/ping", ping.Handle())

	router.POST("/api/start", start.Handle(store))

	// websocket endpoint
	router.GET("/ws/:id", tags.Handle(htmlTemplates, store, hub))

	return router
}
