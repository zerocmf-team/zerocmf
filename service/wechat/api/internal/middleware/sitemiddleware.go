package middleware

import "net/http"

type SiteMiddleware struct {
}

func NewSiteMiddleware() *SiteMiddleware {
	return &SiteMiddleware{}
}

func (m *SiteMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		// Passthrough to next handler if need
		next(w, r)
	}
}
