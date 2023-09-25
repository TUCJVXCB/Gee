package gee

type router struct {
	handlers map[string]HandlerFunc
}

func NewRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRoute(method string, url string, f HandlerFunc) {
	key := method + "-" + url
	r.handlers[key] = f
}
