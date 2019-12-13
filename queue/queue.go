// many ways to implement a queue: linked list, slice, fixed size array (aka circular queue
// length of slice will have length of items in queue, the capacity can change
// uses for queue: BFS

package queue

// int is inside the queue package call it by saying queue.Int
type Int struct {
	slice []int
}

//enqueue adds the integer provided to the back of the Queue
func (q *Int) Enqueue(i int) {
	// adds item to end of slice
	q.slice = append(q.slice, i)
}

// dequeue return the first item in the queue
// and removes that item from the front of the queue
// or panics if there isn't one
func (q *Int) Dequeue() int {
	// return the first item in the queue
	ret := q.slice[0]
	// remove the first item from the queue
	q.slice = q.slice[1:len(q.slice)]
	return ret
}
