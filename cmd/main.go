package main

import (
	"github.com/alemjc/gophercises/urlshort/rest"
	"log"
	"os"
)

func main() {
	log.Fatal(rest.Start(os.Getenv("CONFIG_PATH")))
}
