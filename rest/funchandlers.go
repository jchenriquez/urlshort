package rest

import (
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
	"net/http"
)

type URLShort struct {
	Short string `yaml:"short,omitempty"`
	Long string `yaml:"long,omitempty"`
}

type Urls struct {
	Urls []URLShort `yaml:"urls"`
}

func buildMap (url Urls) map[string] string {
	mp := make(map[string]string)

	for _, url := range url.Urls {
		mp[url.Short] = url.Long
	}

	return mp
}

func MapHandler(mp map[string]string, fallback http.Handler) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {

		attrs := mux.Vars(r)

		url, ok := mp[attrs["name"]]

		if ok {
			http.Redirect(w, r, url, http.StatusFound)
			return
		}

		fallback.ServeHTTP(w, r)
	}
}

func YamlHandler(yBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var urls Urls

	err := yaml.Unmarshal(yBytes, &urls)

	if err != nil {
		return nil, err
	}

	mp := buildMap(urls)

	return MapHandler(mp, fallback), nil

}

func MakeRedirectingHandler() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("%v\n", *request)
		http.Redirect(writer, request, request.URL.Path, http.StatusPermanentRedirect)
	}
}
