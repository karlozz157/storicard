package main

import (
	"os"

	"github.com/karlozz157/storicard/src/infrastructure/interfaces/aws"
	"github.com/karlozz157/storicard/src/infrastructure/interfaces/http"
)

func main() {
	if os.Getenv("ENVIRONMENT") == "aws" {
		aws.StartLambda()
	} else {
		http.StartServer()
	}
}
