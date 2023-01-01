package routes

type Router struct {
	config *Config
	router *chi.Mux
}

func NewRouter() *Router {
	return &Router{
		config: NewConfig().SetTimeout(serviceConfig.GetConfig().timeout),
		router: chi.NewRouter(),
	}
}

func (r *Router) SetRouters() *chi.Mux {}

func (r *Router) SetRouterConfigRouters() {}

func RouterHealth() {}

func RouterProduct() {}

func EnableTimeout() {}

func EnableCORS() {}

func EnableRecovery() {}

func EnableRequestId() {}

func EnableRealIP() {}
