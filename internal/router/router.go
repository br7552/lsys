package router

import (
	"context"
	"net/http"
	"regexp"
	"strings"
)

type Router struct {
	routes           []route
	NotFound         http.HandlerFunc
	MethodNotAllowed http.HandlerFunc
}

func New() *Router {
	return &Router{
		NotFound:         http.NotFound,
		MethodNotAllowed: methodNotAllowed,
	}
}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var allow []string

	for _, route := range rt.routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)

		if len(matches) == 0 {
			continue
		}

		if r.Method != http.MethodOptions && r.Method != route.method {
			allow = append(allow, route.method)
			continue
		}

		params := make(map[string]string)

		for k, v := range route.keys {
			params[k] = matches[v]
		}

		ctx := context.WithValue(r.Context(), paramsContextKey, params)

		route.ServeHTTP(w, r.WithContext(ctx))

		return
	}

	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		rt.MethodNotAllowed(w, r)
		return
	}

	rt.NotFound(w, r)
}

func (rt *Router) HandleFunc(method string, pattern string,
	handler func(http.ResponseWriter, *http.Request)) {
	route := newRoute(method, pattern, http.HandlerFunc(handler))

	rt.routes = append(rt.routes, route)
}

func (rt *Router) Handle(method string, pattern string,
	handler http.Handler) {
	route := newRoute(method, pattern, handler)

	rt.routes = append(rt.routes, route)
}

type route struct {
	method string
	regex  *regexp.Regexp
	keys   map[string]int
	http.Handler
}

func newRoute(method string, pattern string, handler http.Handler) route {
	slug := "([^/]+)"

	re := regexp.MustCompile("/:" + slug)

	keys := make(map[string]int)
	i := 1

	pattern = re.ReplaceAllStringFunc(pattern, func(s string) string {
		k := strings.TrimPrefix(s, "/:")
		keys[k] = i
		i++
		return "/" + slug
	})

	return route{
		method:  method,
		regex:   regexp.MustCompile("^" + pattern + "$"),
		keys:    keys,
		Handler: handler,
	}
}

type contextKey string

const paramsContextKey = contextKey("params")

func Params(ctx context.Context) map[string]string {
	params, ok := ctx.Value(paramsContextKey).(map[string]string)
	if !ok {
		return map[string]string{}
	}

	return params
}

func methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed),
		http.StatusMethodNotAllowed)
}
