package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	ServerConfig "github.com/Chakravarthy7102/serverless/config"
	"github.com/Chakravarthy7102/serverless/internal/repository/adapter"
)

type Router struct {
	config *Config
	router *chi.Mux
}

func NewRouter() *Router {
	return &Router{
		config: NewConfig().SetTimeout(ServerConfig.GetConfig().Timeout),
		router: chi.NewRouter(),
	}
}

func (r *Router) SetRouters(repository adapter.Interface) *chi.Mux {
	r.SetRouterConfigRouters()
	r.RouterHealth(repository)
	r.RouterProduct(repository)

	return r.router
}

func (r *Router) SetRouterConfigRouters() {
	r.EnableCORS()
	r.EnableTimeout()
	r.EnableLogger()
	r.EnableRecovery()
	r.EnableRealIP()
	r.EnableRequestId()
}

func (r *Router) RouterHealth(repository adapter.Interface) {
	handler := HealthHandler.newHandler(repository)
	r.router.Route("/health", func(route chi.Route) {
		route.Post("/", handler.Post)
		route.Get("/", handler.Get)
		route.Put("/", handler.Put)
		route.Delete("/", handler.Delete)
		route.Options("/", handler.Options)

	})
}

func (r *Router) RouterProduct(repository adapter.Interface) {
	handler := ProductHandler.newHandler(repository)

	r.router.Route("/product", func(route chi.Route) {
		route.Post("/", handler.Post)
		route.Get("/", handler.Get)
		route.Put("/{ID}", handler.Put)
		route.Delete("/{ID}", handler.Delete)
		route.Options("/", handler.Options)
	})
}

func (r *Router) EnableTimeout() *Router {
	r.router.Use(middleware.Timeout(r.config.GetTimeout()))
	return r
}

func (r *Router) EnableCORS() *Router {
	r.router.Use(r.config.Cors)
	return r
}

func (r *Router) EnableRecovery() *Router {
	r.router.Use(middleware.Recoverer)
	return r
}

func (r *Router) EnableRequestId() *Router {
	r.router.Use(middleware.RequestId)
	return r
}

func (r *Router) EnableRealIP() *Router {
	r.router.Use(middleware.RealIP)
	return r
}

func (r *Router) EnableLogger() *Router {
	r.router.Use(middleware.Logger)
	return r
}
