package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"kbconcuradmin/api/programmesneufs"
	"kbconcuradmin/api/test"
	kbmiddlewares "kbconcuradmin/middlewares"
)

type todo struct {
	ID         string    `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
	Task       string    `json:"task"`
}

const PORT = "3333"

//noinspection ALL
func main() {
	fmt.Println("server is running on..." + PORT)
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(kbmiddlewares.CorsMiddleware.Handler)
	// permislocaux
	// r.Route("/api/permislocaux", permislocaux.Router)

	// programmesneufs
	r.Route("/api/programmesneufs", programmesneufs.Router)
	// test JSON on collection "test" etc ...
	r.Route("/api/test", test.Router)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	http.ListenAndServe(":"+PORT, r)

}
