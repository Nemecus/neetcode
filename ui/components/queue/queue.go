package queue

import "reflect"

type dequeue struct {
	items []int
}

var Queue dequeue

func createQueue() dequeue {
	return Queue
}

func (d *dequeue) isEmpty() bool {
	return reflect.ValueOf(d).IsZero()
}

func (d *dequeue) append(item int) {
	d.items = append(d.items, item)
}

func (d *dequeue) appendleft(item int) {
	d.items = append([]int{item}, d.items...)
}

func (d *dequeue) pop() int {
	if len(d.items) == 0 {
		return -1
	}

	last := d.items[len(d.items)-1]
	d.items = d.items[:len(d.items)-1]
	return last
}

func (d *dequeue) popleft() int {
	if len(d.items) == 0 {
		return -1
	}

	first := d.items[0]
	d.items = d.items[1:]
	return first
}
