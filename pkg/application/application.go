package application

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Application struct {
	Config *Config

	Router   *mux.Router
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Database *storage.DB
}

func NewApplication() *Application {
	app := &Application{
		Config: NewConfig(),
	}

	app.configureLogger()
	app.configureRouter()
	err := app.configureDatabase()
	if err != nil {
		panic(err)
	}

	return app
}

func (app *Application) configureLogger() {
	app.InfoLog = log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	app.ErrorLog = log.New(os.Stderr, "[ERRO]\t", log.Ldate|log.Ltime|log.Lshortfile)
}

func (app *Application) configureRouter() {
	router := mux.NewRouter()

	fileServer := http.FileServer(http.Dir("./static"))

	router.HandleFunc("/", Home(app))
	router.HandleFunc("/log", Login(app))
	router.HandleFunc("/reg", Register(app))

	router.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", fileServer),
	)

	app.Router = router
}

func (app *Application) configureDatabase() error {
	return nil
}

func (app *Application) Info(content ...interface{}) {
	app.InfoLog.Println(content...)
}

func (app *Application) Error(content ...interface{}) {
	app.ErrorLog.Println(content...)
}
