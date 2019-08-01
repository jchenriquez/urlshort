package rest

import (
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/gorilla/mux"
	"net/http"
)

type URLShort struct {
	Short string `json:"short,omitempty"`
	Long  string `json:"long,omitempty"`
}

type Urls struct {
	Urls []URLShort `yaml:"urls"`
}

func MapHandler(db *bolt.DB, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var urlLong []byte

		attrs := mux.Vars(r)

		err := db.View(func(tx *bolt.Tx) error {
			bucket := tx.Bucket([]byte("url_maps"))

			urlLong = bucket.Get([]byte(attrs["name"]))
			fmt.Println(string(urlLong))
			if urlLong != nil {
				http.Redirect(w, r, fmt.Sprintf("http://%s", string(urlLong)), http.StatusMovedPermanently)
			} else {
				fallback.ServeHTTP(w, r)
			}

			return nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}
}

func urlWriteHandler(db *bolt.DB, mapper URLShort) error {

	return db.Update(func(tx *bolt.Tx) error {
		bucket, _ := tx.CreateBucketIfNotExists([]byte("url_maps"))

		err := bucket.Put([]byte(mapper.Short), []byte(mapper.Long))

		if err != nil {
			return err
		}

		return nil

	})

}

func urlReadAllHandler(db *bolt.DB) ([]URLShort, error) {
	result := make([]URLShort, 0, 0)
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("url_maps"))

		if bucket == nil {
			return nil
		}

		cursor := bucket.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			result = append(result, URLShort{
				Long:  string(v),
				Short: string(k),
			})
		}

		return nil

	})

	return result, err
}

func urlDeleteHandler(db *bolt.DB, urlDef URLShort) error {

	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("url_maps"))
		return bucket.Delete([]byte(urlDef.Short))
	})

}

func makeAdminHandler(db *bolt.DB) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		var payload []URLShort
		var err error
		writer.Header().Set("Content-Type", "application/json")
		switch request.Method {
		case "GET":
			payload, err = urlReadAllHandler(db)

			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}

		case "POST":
			var urlPayload URLShort
			err := json.NewDecoder(request.Body).Decode(&urlPayload)

			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}

			err = urlWriteHandler(db, urlPayload)

			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			} else {
				payload, err = urlReadAllHandler(db)
				writer.WriteHeader(http.StatusOK)
			}

		case "DELETE":

			query := request.URL.Query()
			var urlPayload URLShort
			urlPayload.Short = query.Get("short")
			urlPayload.Long = query.Get("Long")

			err = urlDeleteHandler(db, urlPayload)

			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			} else {
				payload, err = urlReadAllHandler(db)
				writer.WriteHeader(http.StatusOK)
			}
		}

		err = json.NewEncoder(writer).Encode(payload)

		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
		}

	}

}
