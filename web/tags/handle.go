package tags

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"text/template"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	"github.com/masl/answertag/storage"
)

var upgrader = websocket.Upgrader{}

func Handle(html *template.Template, store storage.Store) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			slog.Info("upgrade:", "error", err)
			return
		}

		defer c.Close()
		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				slog.Info("read:", "error", err)
				break
			}
			slog.Info("recv:", "message", message)

			// unmarshal message
			// TODO: correct struct for message
			var tagData struct {
				Tag string `json:"tag"`
			}
			err = json.Unmarshal(message, &tagData)
			if err != nil {
				slog.Info("unmarshal:", "error", err)
				break
			}

			// add tag to storage
			err = store.AddTag(tagData.Tag)
			if err != nil {
				slog.Info("add tag:", "error", err)
				break
			}

			// get all tags and write message with tags.html templates
			tagsData, err := store.GetAllTags()

			var buf bytes.Buffer
			err = html.ExecuteTemplate(&buf, "tags.html", tagsData)
			if err != nil {
				slog.Info("execute template:", "error", err)
				break
			}

			err = c.WriteMessage(mt, buf.Bytes())
			if err != nil {
				slog.Info("write:", "error", err)
				break
			}
		}
	}
}
