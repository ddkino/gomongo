package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"hellomongo/api/permislocaux"
	"hellomongo/api/test"
)

type todo struct {
	ID         string    `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
	Task       string    `json:"task"`
}

//noinspection ALL
func main() {
	fmt.Println("server is running ...")
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	// permislocaux CRUD
	r.Route("/api/permislocaux", permislocaux.Router)
	// test CRUD + JSON on collection "test" etc ...
	r.Route("/api/test", test.Router)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	http.ListenAndServe(":3333", r)

}
