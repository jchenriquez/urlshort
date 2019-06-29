package models

import (
	"fmt"
	"net/http"
)

type NotFoundHandler string

func (nf NotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	fmt.Fprintf(w,
		"<h1>%s</h1>" +
		"<a href='%s'>%s</a> was not found. <br/>Please add mapping or try again.", nf, url, url)
}
