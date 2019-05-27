package server

import (
	"github.com/go-chi/chi"
	"net/http"
	"testServerStats/pkg/statistics"
	log "github.com/sirupsen/logrus"
)

func initializeRoutes(r *chi.Mux) {
	r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("ok"))
		w.WriteHeader(http.StatusOK)
		if err != nil {
			log.WithError(err).Error("Failed to write to output")
		}
	})

	r.Route("/api/users", func(r chi.Router) {
		r.Post("/", statistics.CreateUser)

		r.Route("/stats", func(r chi.Router) {
			r.Post("/", statistics.CreateStat)
			r.Get("/topAccum", statistics.GetAccumulateStats)
			r.Get("/top", statistics.GetStatsPerDay)
		})
	})
}


