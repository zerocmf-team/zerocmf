/**
** @创建时间: 2022/8/24 12:24
** @作者　　: return
** @描述　　:
 */

package sessions

import (
	"github.com/gorilla/sessions"
	"os"
)

var (
	store *sessions.CookieStore
)

func NewStore() *sessions.CookieStore {
	if store == nil {
		store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	}
	return store
}
