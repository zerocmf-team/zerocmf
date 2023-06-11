/**
** @创建时间: 2022/3/13 17:00
** @作者　　: return
** @描述　　:
 */

package middleware

import (
	"net/http"
)

type InitMiddleware struct {
}

func NewInitMiddleware() *InitMiddleware {
	return &InitMiddleware{}
}

func (m *InitMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
	}
}
