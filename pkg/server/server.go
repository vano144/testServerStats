package server

import (
	"github.com/go-chi/chi"
	"testServerStats/pkg/db"
)

func GetRouter() *chi.Mux {
	r := chi.NewRouter()
	db.InitDb()
	InitValidators()
	initializeMiddlewares(r)
	initializeRoutes(r)
	return r
}
