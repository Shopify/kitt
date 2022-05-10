package queue

import "fmt"

type BoundedQueue[T any] struct {
	capacity int
	queue    *Queue[T]
}

func NewBoundedQueue[T any](capacity int) *BoundedQueue[T] {
	q := make(Queue[T], 0, capacity)
	bq := &BoundedQueue[T]{
		capacity: capacity,
		queue:    &q,
	}
	return bq
}

func (bq *BoundedQueue[T]) IsFull() bool {
	return bq.Len() == bq.Cap()
}

func (bq *BoundedQueue[T]) IsEmpty() bool {
	return bq.queue.IsEmpty()
}

func (bq *BoundedQueue[T]) Cap() int {
	return bq.capacity
}

func (bq *BoundedQueue[T]) Len() int {
	return bq.queue.Len()
}

func (bq *BoundedQueue[T]) Enqueue(values ...T) error {
	if len(values)+bq.queue.Len() > (*bq).capacity {
		return fmt.Errorf("queue would overflow")
	}
	bq.queue.Enqueue(values...)
	return nil
}

func (bq *BoundedQueue[T]) Dequeue() (T, error) {
	return bq.queue.Dequeue()
}

func (bq *BoundedQueue[T]) Peek() (T, error) {
	return bq.queue.Peek()
}

func (bq *BoundedQueue[T]) Values() []T {
	return bq.queue.Values()
}
