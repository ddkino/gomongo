package test

import (
	"github.com/go-chi/chi"
)

func Router(r chi.Router) {
	r.Get("/json", handleProfileGetJson)
	r.Get("/id", handleProfileGetByName)
	r.Post("/", handleProfileInsertOne)
}
