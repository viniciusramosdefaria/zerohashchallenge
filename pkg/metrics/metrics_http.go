package metrics

import (
	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/viniciusramosdefaria/zerohashchallenge/pkg/phttp"
	"net/http"
)

type httpHandler struct {
}

func NewHTTPHandler() phttp.HttpHandler {
	return httpHandler{}
}

func (h httpHandler) Handler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		promhttp.Handler().ServeHTTP(w, r)
	}
}