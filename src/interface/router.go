package handler

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Service interface {
	Server(ctx context.Context)
}

func (s *ServiceDriver) Server(ctx context.Context) {

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Route("/coffee", func(r chi.Router) {
			r.Get("/get", func(w http.ResponseWriter, r *http.Request) {
				s.CoffeeGet(ctx, w, r)
			})

			r.Get("/create", func(w http.ResponseWriter, r *http.Request) {
				s.CoffeeCreate(ctx, w, r)
			})
		})

		r.Route("/user", func(r chi.Router) {
			r.Get("/get", func(w http.ResponseWriter, r *http.Request) {
				s.UserGet(ctx, w, r)
			})
		})

		r.Route("/join", func(r chi.Router) {
			r.Get("/delete", func(w http.ResponseWriter, r *http.Request) {
				s.JoinDelete(ctx, w, r)
			})

			r.Get("/create", func(w http.ResponseWriter, r *http.Request) {
				s.JoinCreate(ctx, w, r)
			})

			r.Route("/get", func(r chi.Router) {
				r.Get("/user", func(w http.ResponseWriter, r *http.Request) {
					s.JoinGetUser(ctx, w, r)
				})

				r.Get("/coffee", func(w http.ResponseWriter, r *http.Request) {
					s.JoinGetCoffee(ctx, w, r)
				})
			})
		})
	})

	addr := os.Getenv("Addr")
	if addr == "" {
		addr = ":8080"
	}

	log.Printf("listen: %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("!! %+v", err)
	}
}
