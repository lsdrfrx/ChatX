package application

import (
	"fmt"
	"net/http"
)

func Home(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Главная страница")
	}
}

func Login(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Логин")
	}
}

func Register(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Регистрация")
	}
}