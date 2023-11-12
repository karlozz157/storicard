package http

import (
	"context"
	"encoding/base64"
	"io"
	"log"
	"net/http"
	"strings"

	e "github.com/karlozz157/storicard/src/domain/errors"

	"github.com/karlozz157/storicard/src/application"
	"github.com/karlozz157/storicard/src/utils"
)

type StoriServer struct {
}

func StartServer() {
	s := StoriServer{}
	s.Serve()
}

func (s *StoriServer) Serve() {
	db := utils.InitMongoDb()
	handler := application.NewTransactionHandler(db)

	http.HandleFunc("/storicard", func(w http.ResponseWriter, r *http.Request) {
		res, err := handler.CreateSummary(context.Background(), s.getBody(r))

		w.Header().Set("Content-Type", "application/json")

		if err != nil {
			message, statusCode := s.parseError(err)
			http.Error(w, message, statusCode)
			return
		}

		w.WriteHeader(res.StatusCode)
		w.Write([]byte(res.Message))
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func (s *StoriServer) getBody(r *http.Request) io.Reader {
	body, _ := io.ReadAll(r.Body)
	content, _ := base64.StdEncoding.DecodeString(string(body))

	return strings.NewReader(string(content))
}

func (s *StoriServer) parseError(err error) (string, int) {
	message := "houston, we have a problem"
	statusCode := http.StatusInternalServerError

	if errStori, ok := err.(*e.ErrStori); ok {
		message = errStori.Mesage
		statusCode = errStori.StatusCode
	}

	return message, statusCode
}
