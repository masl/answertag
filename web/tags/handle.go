package tags

import (
	"net/http"
	"text/template"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	"github.com/masl/answertag/storage"
	"github.com/masl/answertag/ws"
)

var upgrader = websocket.Upgrader{}

func Handle(html *template.Template, store storage.Store, hub *ws.Hub) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ws.ServeWs(hub, w, r)
	}
}
