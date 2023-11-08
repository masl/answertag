package tags

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/masl/answertag/storage"
	"github.com/masl/answertag/ws"
)

func Handle(store storage.Store, hub *ws.Hub) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		cloudID := ps.ByName("id")

		// check if cloud exists
		// TODO: separate exists function in storage?
		_, err := store.ReadByID(cloudID)
		if err == storage.ErrNotFound {
			http.NotFound(w, r)
			return
		}

		ws.ServeWs(hub, w, r, cloudID)
	}
}
