package routes

import (
	"errors"
	"net/http"
	"quoteapp/view"
	
)

func (r *Route) get(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			next.ServeHTTP(w, r)
		} else {
			view.AnyErrorRespond(w, http.StatusMethodNotAllowed, errors.New("method is not allowed"))
		}
	})
}

func (r *Route) post(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			next.ServeHTTP(w, r)
		} else {
			view.AnyErrorRespond(w, http.StatusMethodNotAllowed, errors.New("method is not allowed"))
		}
	})
}

func (r *Route) put(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			next.ServeHTTP(w, r)
		} else {
			view.AnyErrorRespond(w, http.StatusMethodNotAllowed, errors.New("method is not allowed"))
		}
	})
}

func (r *Route) delete(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			next.ServeHTTP(w, r)
		} else {
			view.AnyErrorRespond(w, http.StatusMethodNotAllowed, errors.New("method is not allowed"))
		}
	})
}




