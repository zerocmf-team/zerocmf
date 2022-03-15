/**
** @创建时间: 2022/3/13 17:42
** @作者　　: return
** @描述　　:
 */

package data

import "sync"

type Context struct {
	// This mutex protect Keys map
	mu sync.RWMutex

	// Keys is a key/value pair exclusively for the context of each request.
	Keys map[string]interface{}
}

// Set is used to store a new key/value pair exclusively for this context.
// It also lazy initializes  c.Keys if it was not used previously.
func (s *Context) Set(key string, value interface{}) {
	s.mu.Lock()
	if s.Keys == nil {
		s.Keys = make(map[string]interface{})
	}

	s.Keys[key] = value
	s.mu.Unlock()
}

// Get returns the value for the given key, ie: (value, true).
// If the value does not exists it returns (nil, false)
func (s *Context) Get(key string) (value interface{}, exists bool) {
	s.mu.RLock()
	value, exists = s.Keys[key]
	s.mu.RUnlock()
	return
}


