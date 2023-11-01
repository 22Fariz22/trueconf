package app

import (
	"net/http"
	"refactoring/internal/config"
	"refactoring/internal/user"
	filejson "refactoring/internal/user/repo/file-json"
	"refactoring/pkg/logger"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)


type app struct {
	cfg        *config.Config
	httpServer *http.Server
	UC         user.UseCaseUser
}

// NewApp create
func NewApp(cfg *config.Config) *app {

	// Repository
		var repo user.RepositoryUser

		repo = filejson.NewInFileRepository(cfg)


	return &app{
		cfg:        cfg,
		httpServer: nil,
		UC:         usecase.NewUseCase(repo),
	}
}

func (a *app) Run() {
	l := logger.New(a.cfg.Log.Level)
	l.Info("app start")

	

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	hd := handler.NewHandler(l,a.cfg, a.UC)

	r.Post("/", hd.someFunc())

	http.ListenAndServe(a.cfg.Port, r)
}