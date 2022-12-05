package routes

type Route interface {
	Setup()
}

type Routes interface {
	Setup()
}

type routes []Route

func NewRoutes(postRoutes PostRoutes, userRoutes UserRoutes) Routes {
	return routes{postRoutes, userRoutes}
}

func (r routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
