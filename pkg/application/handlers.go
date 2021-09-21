package application

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Главная страница")
	}
}

func Login(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			tmpl, err := template.ParseFiles(
				"./static/html/base.layout.html",
				"./static/html/header.partial.html",
				"./static/html/footer.partial.html",
				"./static/html/login.page.html",
			)
			if err != nil {
				app.Error(err)
				return
			}
	
			if err := tmpl.Execute(w, nil); err != nil {
				app.Error(err)
				return
			}
		}

		if r.Method == http.MethodPost {
			r.ParseForm()

			creds := make(map[string]string)
			creds[r.FormValue("username")] = r.FormValue("password")

			fmt.Println(creds)

			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func Register(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fmt.Fprintf(w, "Регистрация")
		}

		if r.Method == http.MethodPost {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}