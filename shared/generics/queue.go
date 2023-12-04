package generics

import (
	"fmt"
)

type Queue[T any] struct {
  list   []T
  length int
}

func NewQueue[T any]() *Queue[T] {
  return &Queue[T]{make([]T, 0), 0}
}

func (q *Queue[T]) Len() int {
	return q.length
}

func (q *Queue[T]) IsEmpty() bool {
	return q.length == 0
}

func (q *Queue[T]) Enqueue(value T) {
  q.list = append(q.list, value)
	q.length++
}

func (q *Queue[T]) Dequeue() (T, error) {
  if q.IsEmpty() {
    return *new(T), fmt.Errorf("tried to dequeue from and empty queue")
  }
	value := q.list[0]
  q.list = q.list[1:]
  q.length--
	return value, nil
}
