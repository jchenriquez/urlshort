package rest

import (
	"fmt"
	"github.com/alemjc/gophercises/urlshort/rest/models"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
)

type config struct {
	Mappings string `yaml:"mappings_file,omitempty"`
	Port string `yaml:"server_port,omitempty"`
}

func Start(confFilePath string) error {
	var conf config

	bytes,err := ioutil.ReadFile(confFilePath)
	if err != nil{
		return err
	}

	err = yaml.Unmarshal(bytes, &conf)
	if err != nil {
		return err
	}

	bytes, err = ioutil.ReadFile(conf.Mappings)

	router := mux.NewRouter()

	yamlHandler, err := YamlHandler(bytes, models.NotFoundHandler("Not Found"))
	if err != nil {
		return err
	}

	router.HandleFunc("/{name:[a-zA-z]+}", yamlHandler)
	allowedOptions := handlers.AllowedMethods([]string{"GET"})

	return http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), handlers.CORS(allowedOptions)(router))
}
