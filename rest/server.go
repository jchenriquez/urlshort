package rest

import (
	"errors"
	"fmt"
	"github.com/alemjc/gophercises/urlshort/rest/models"
	"github.com/boltdb/bolt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
)

type config struct {
	DBPath string `yaml:"db_file_path"`
}

func Start(configFilePath string) error {
	var cnf config
	bytes, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(bytes, &cnf)

	if err != nil {
		return err
	}

	db, err := bolt.Open(cnf.DBPath, 0666, nil)

	if err != nil {
		return err
	}

	defer db.Close()

	router := mux.NewRouter()
	router.Host("http://localhost")
	mapHandler := MapHandler(db, models.NotFoundHandler("Not Found"))
	dir, err := os.Getwd()
	if err != nil {
		return errors.New("could not locate web directory")
	}
	fs := http.FileServer(http.Dir(fmt.Sprintf("%s/%s", dir, "web/build")))
	router.PathPrefix("/static/").Handler(fs)
	router.Handle("/admin", http.StripPrefix("/admin", fs))
	router.HandleFunc("/admin/urls", makeAdminHandler(db))
	router.HandleFunc("/{name:[a-zA-z]+}", mapHandler)
	allowedOptions := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "OPTIONS"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})

	return http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), handlers.CORS(allowedOptions, allowedHeaders, allowedOrigins)(router))
}
