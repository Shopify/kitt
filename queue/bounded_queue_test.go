package queue

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func ExampleNewBoundedQueue() {
	bq := NewBoundedQueue[int](3)
	fmt.Println(bq.Cap(), bq.Len(), bq.IsEmpty())
	// Output: 3 0 true
}

func ExampleBoundedQueue_IsFull() {
	bq := NewBoundedQueue[int](1)
	_ = bq.Enqueue(1)
	fmt.Println(bq.IsFull())
	// Output: true
}

func ExampleBoundedQueue_IsEmpty() {
	bq := NewBoundedQueue[int](1)
	fmt.Println(bq.IsEmpty())
	// Output: true
}

func ExampleBoundedQueue_Cap() {
	bq := NewBoundedQueue[int](3)
	fmt.Println(bq.Cap())
	// Output: 3
}

func ExampleBoundedQueue_Len() {
	bq := NewBoundedQueue[int](3)
	_ = bq.Enqueue(1, 2)
	fmt.Println(bq.Len())
	// Output: 2
}

func ExampleBoundedQueue_Enqueue() {
	bq := NewBoundedQueue[int](3)
	fmt.Println(bq.Enqueue(1, 2, 3))
	// Output: <nil>
}

func ExampleBoundedQueue_Enqueue_error() {
	bq := NewBoundedQueue[int](3)
	fmt.Println(bq.Enqueue(1, 2, 3, 4))
	// Output: queue would overflow
}

func ExampleBoundedQueue_Dequeue() {
	bq := NewBoundedQueue[int](1)
	_ = bq.Enqueue(1)
	v, err := bq.Dequeue()
	fmt.Println(v, err)
	// Output: 1 <nil>
}

func ExampleBoundedQueue_Dequeue_error() {
	bq := NewBoundedQueue[int](1)
	v, err := bq.Dequeue()
	fmt.Println(v, err)
	// Output: 0 queue is empty
}

func ExampleBoundedQueue_Peek() {
	bq := NewBoundedQueue[int](1)
	_ = bq.Enqueue(1)
	fmt.Println(bq.Peek())
	// Output: 1 <nil>
}

func ExampleBoundedQueue_Values() {
	bq := NewBoundedQueue[int](3)
	_ = bq.Enqueue(1, 2, 3)
	fmt.Println(bq.Values())
	// Output: [1 2 3]
}

func TestBoundedQueue_Enqueue(t *testing.T) {
	bq := NewBoundedQueue[int](3)
	require.Equal(t, 3, bq.Cap())

	err := bq.Enqueue(1, 2, 3)
	require.NoError(t, err)
	require.Equal(t, 3, bq.Len())

	err = bq.Enqueue(4)
	require.Error(t, err)
	require.Equal(t, 3, bq.Len())
}

func TestBoundedQueue_Dequeue(t *testing.T) {
	bq := NewBoundedQueue[int](3)
	require.Equal(t, 3, bq.Cap())

	err := bq.Enqueue(1, 2, 3)
	require.NoError(t, err)
	require.Equal(t, 3, bq.Len())
	require.True(t, bq.IsFull())
	require.False(t, bq.IsEmpty())

	v, err := bq.Dequeue()
	require.NoError(t, err)
	require.Equal(t, 1, v)
	require.Equal(t, bq.Len(), 2)
	require.False(t, bq.IsFull())
	require.False(t, bq.IsEmpty())

	v, err = bq.Dequeue()
	require.NoError(t, err)
	require.Equal(t, 2, v)
	require.Equal(t, bq.Len(), 1)
	require.False(t, bq.IsFull())
	require.False(t, bq.IsEmpty())

	v, err = bq.Dequeue()
	require.NoError(t, err)
	require.Equal(t, 3, v)
	require.Equal(t, bq.Len(), 0)
	require.False(t, bq.IsFull())
	require.True(t, bq.IsEmpty())

	v, err = bq.Dequeue()
	require.Error(t, err)
}
