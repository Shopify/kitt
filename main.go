package main

import (
	"fmt"

	"github.com/Shopify/kitt/queue"
	"github.com/Shopify/kitt/set"
)

func main() {
	s1 := set.FromSlice([]int{1, 2, 3})
	s1.Add(3)
	fmt.Printf("Set members: %+v\n", s1.Values())

	s2 := set.FromSlice([]int{3, 4, 5})
	fmt.Printf("s1 union s2: %+v\n", s1.Union(s2).Values())
	fmt.Printf("s1 intersection s2: %+v\n", s1.Intersection(s2).Values())

	q := queue.NewQueue[int]()
	q.Enqueue(1, 2, 3)
	fmt.Printf("\nQueue values: %+v\n", q.Values())
	v, _ := q.Peek()
	fmt.Printf("peek: %d, values: %+v\n", v, q.Values())
	v, _ = q.Dequeue()
	fmt.Printf("dequeue: %d, values: %+v\n", v, q.Values())
	q.Enqueue(1)
	fmt.Printf("values: %+v\n", q.Values())
}
