package main 

type middleware func(http.Handler)http.Handler

type Router struct {
	middlewareChain []middleware
	mux map[string]http.Handler
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router)Use(m middleware){
	r.middlewareChain = append(r.middlewareChain, m)
}

func (r *Router)Add(route, string, h http.Handler){
	var mergeHandler = h 

	for i := len(r.middlewareChain) -1; i >= 0; i-- {
		mergeHandler = r.middlewareChain[i](mergeHandler)
	}
	r.mux[route] = mergeHandler
}


