package programmesneufs

import (
	"github.com/go-chi/chi"
)


func Router(r chi.Router) {
	r.Get("/test", HandleProgrammesneufsTest)
	r.Post("/test", HandleProgrammesneufsTest)
	r.Post("/bydate", HandleProgrammesneufsFindByDate)
	r.Get("/", HandleProgrammesneufsGetAll)
	// r.Get("/{id}", HandlePermislocauxGetById)
}
