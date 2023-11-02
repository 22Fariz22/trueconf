package app

import (
	"net/http"
	"time"

	"github.com/22Fariz22/trueconf/internal/app/api"
	"github.com/22Fariz22/trueconf/internal/config"
	"github.com/22Fariz22/trueconf/internal/user"
	"github.com/22Fariz22/trueconf/pkg/logger"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type App interface {
	Run() error
}

type app struct {
	cfg        *config.Config
	httpServer *http.Server
	UC         user.UseCaseUser
}

// NewApp create
func NewApp(cfg *config.Config) App {
	// var repo user.RepositoryUser

	return &app{
		cfg:        cfg,
		httpServer: &http.Server{},
		//		UC:         usecase.NewUseCase(repo),
	}
}

func (a *app) Run() error {
	l := logger.New(a.cfg.Log.Level)
	l.Infof("app start")

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(time.Now().String()))
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Route("/users", func(r chi.Router) {
				r.Get("/", api.SearchUsers)
				r.Post("/", api.CreateUser)

				r.Route("/{id}", func(r chi.Router) {
					r.Get("/", api.GetUser)
					r.Patch("/", api.UpdateUser)
					r.Delete("/", api.DeleteUser)
				})
			})
		})
	})

	a.httpServer.Handler = r
	a.httpServer.Addr = a.cfg.Port
	err := a.httpServer.ListenAndServe()
	if err != http.ErrServerClosed {
		return err
	}
	return nil
}
