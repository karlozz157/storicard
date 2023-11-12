package http

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/karlozz157/storicard/src/application"
	e "github.com/karlozz157/storicard/src/domain/errors"
	"github.com/karlozz157/storicard/src/utils"
)

const local = "127.0.0.1"

type StoriServer struct {
	http.Server
}

func StartServer() {
	server := &StoriServer{}
	server.Addr = fmt.Sprintf("%s:%s", local, os.Getenv("PORT"))
	server.dispatchHanlders()

	log.Fatal(server.ListenAndServe())
}

func (s *StoriServer) dispatchHanlders() {

	db := utils.InitMongoDb()
	handler := application.NewTransactionHandler(db)

	router := mux.NewRouter()
	router.HandleFunc("/storicard/{email}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		email := vars["email"]

		res, err := handler.CreateSummary(context.Background(), email, s.getReaderFromRequest(r))

		if err != nil {
			statusCode, _ := e.ParseError(err)
			w.WriteHeader(statusCode)
			json.NewEncoder(w).Encode(err)
			return
		}

		w.WriteHeader(res.StatusCode)
		json.NewEncoder(w).Encode(res)
	})

	s.Handler = router
}

func (s *StoriServer) getReaderFromRequest(r *http.Request) io.Reader {
	body, _ := io.ReadAll(r.Body)
	content, _ := base64.StdEncoding.DecodeString(string(body))

	return strings.NewReader(string(content))
}
