package generics

import (
	"sync"
)

type Set[K comparable] struct {
  m   map[K]struct{}
  mut *sync.RWMutex
}

func NewSet[K comparable]() *Set[K] {
  return &Set[K]{make(map[K]struct{}, 0), &sync.RWMutex{}}
}

func (s *Set[K]) Len() int {
	return len(s.m) 
}

func (s *Set[K]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Set[K]) Add(value K) {
  s.mut.Lock()
  defer s.mut.Unlock()
  s.m[value] = struct{}{}
}

func (s *Set[K]) Get(value K) bool {
  s.mut.RLock()
  defer s.mut.RUnlock()
  _, ok := s.m[value]
  return ok
}
