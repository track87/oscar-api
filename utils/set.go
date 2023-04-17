// Package utils declare something
// MarsDong 2023/3/31
package utils

import (
	"reflect"
	"sync"
)

type Set struct {
	sync.RWMutex
	kinds map[string]reflect.Type
}

func (s *Set) Register(kind string, obj interface{}) {
	s.Lock()
	defer s.Unlock()

	s.kinds[kind] = reflect.TypeOf(obj).Elem()
}

func (s *Set) Clone() map[string]reflect.Type {
	s.RLock()
	defer s.RUnlock()

	out := make(map[string]reflect.Type)
	for key, kind := range s.kinds {
		out[key] = kind
	}
	return out
}

func (s *Set) Exist(kind string) bool {
	s.RLock()
	defer s.RUnlock()

	_, exists := s.kinds[kind]
	return exists
}

func (s *Set) Get(kind string) reflect.Type {
	s.RLock()
	defer s.RUnlock()

	v, _ := s.kinds[kind]
	return v
}

func (s *Set) AllKeys() []string {
	keys := make([]string, 0)
	s.RLock()
	defer s.RUnlock()

	for key, _ := range s.kinds {
		keys = append(keys, key)
	}
	return keys
}

func NewSet() *Set {
	return &Set{
		kinds: make(map[string]reflect.Type),
	}
}
