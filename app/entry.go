package main

import (
	"fmt"
	queue "godsa/queue"
)

func main() {
	// q := queue.Int
	// queue.Int is not an expression
	var q queue.ItemQueue

	q.Enqueue(123)
	q.Enqueue(43)
	q.Enqueue(99)
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q)
	fmt.Println(q.Dequeue())
	fmt.Println(q)
}
