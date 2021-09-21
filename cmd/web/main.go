package main

import (
	"chattex/pkg/application"
	"net/http"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func main() {
	app := application.NewApplication()

	s := http.Server{
		Handler:  app.Router,
		ErrorLog: app.ErrorLog,
		Addr:     app.Config.Addr,
	}

	app.Error(s.ListenAndServe())
}
