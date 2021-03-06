package phttp

import "github.com/julienschmidt/httprouter"

type HttpHandler interface {
	Handler() httprouter.Handle
}