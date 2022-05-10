package queue

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func ExampleNewQueue() {
	q := NewQueue[int]()
	fmt.Println(q.Len(), q.IsEmpty())
	// Output: 0 true
}

func ExampleQueue_IsEmpty() {
	q := NewQueue[int]()
	fmt.Println(q.IsEmpty())
	// Output: true
}

func ExampleQueue_Len() {
	q := NewQueue[int]()
	q.Enqueue(1, 2, 3)
	fmt.Println(q.Len())
	// Output: 3
}

func ExampleQueue_Enqueue() {
	q := NewQueue[int]()
	q.Enqueue(3)
	q.Enqueue(1, 2)
	fmt.Println(q.Values())
	// Output: [3 1 2]
}

func ExampleQueue_Dequeue() {
	q := NewQueue[int]()
	q.Enqueue(1)
	fmt.Println(q.Dequeue())
	// Output: 1 <nil>
}

func ExampleQueue_Dequeue_error() {
	q := NewQueue[int]()
	fmt.Println(q.Dequeue())
	// Output: 0 queue is empty
}

func ExampleQueue_Peek() {
	q := NewQueue[int]()
	q.Enqueue(1)
	fmt.Println(q.Peek())
	// Output: 1 <nil>
}

func ExampleQueue_Peek_error() {
	q := NewQueue[int]()
	fmt.Println(q.Peek())
	// Output: 0 queue is empty
}

func ExampleQueue_Values() {
	q := NewQueue[int]()
	q.Enqueue(1, 2, 3)
	fmt.Println(q.Values())
	// Output: [1 2 3]
}

func TestQueue_Basics(t *testing.T) {
	q := NewQueue[int]()
	require.True(t, q.IsEmpty())

	for i := 0; i < 10_000; i++ {
		q.Enqueue(i)
	}
	require.False(t, q.IsEmpty())
	require.Equal(t, 10_000, q.Len())

	for i := 0; i < 10_000; i++ {
		dq, err := q.Dequeue()
		require.NoError(t, err)
		require.Equal(t, i, dq)
	}
	require.True(t, q.IsEmpty())

	_, err := q.Dequeue()
	require.Error(t, err)
}

func TestQueue_Peek(t *testing.T) {
	q := NewQueue[int]()
	_, err := q.Peek()
	require.Error(t, err)

	q.Enqueue(1)
	v, err := q.Peek()
	require.NoError(t, err)
	require.Equal(t, 1, v)
}
