package oas3

import (
	"net/http"
)

func (o *Oas3) Add(path string, method string, operation *Operation) {
	if o.Paths == nil {
		o.Paths = make(Paths, 0)
	}

	originalPath := o.Paths[path]

	switch method {
	case http.MethodGet:
		originalPath.Get = operation
		break
	case http.MethodPost:
		originalPath.Post = operation
		break
	case http.MethodPut:
		originalPath.Put = operation
		break
	case http.MethodPatch:
		originalPath.Patch = operation
		break
	case http.MethodDelete:
		originalPath.Delete = operation
		break
	case http.MethodOptions:
		originalPath.Options = operation
		break
	case http.MethodTrace:
		originalPath.Trace = operation
		break
	default:
	}

	o.Paths[path] = originalPath
}
