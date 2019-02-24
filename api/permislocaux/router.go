package permislocaux

import (
	"github.com/go-chi/chi"
)


func Router(r chi.Router) {
	r.Get("/", HandlePermislocauxGetAll)
	r.Get("/{id}", HandlePermislocauxGetById)
}
