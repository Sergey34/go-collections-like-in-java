package queue

import "errors"

type ArrayDeque[E any] struct {
	elementData []E
	head        int
	tail        int
}

func NewArrayDeque[E any](size int) ArrayDeque[E] {
	var a = make([]E, 0, size)
	return ArrayDeque[E]{elementData: a, head: 0, tail: 0}
}

func (queue *ArrayDeque[E]) Iterator() []E {
	return queue.elementData
}

func (queue *ArrayDeque[E]) Add(e E) bool {
	queue.elementData = append(queue.elementData, e)
	queue.tail++
	return true
}

func (queue *ArrayDeque[E]) Element() (E, error) {
	if queue.head == queue.tail {
		return nil, errors.New("NoSuchElementException")
	}
	return queue.elementData[queue.head], nil
}

func (queue *ArrayDeque[E]) Offer(e E) bool {
	if len(queue.elementData) < cap(queue.elementData) {
		queue.Add(e)
		return true
	}
	return false
}

func (queue *ArrayDeque[E]) Peek() E {
	if queue.head == queue.tail {
		return nil
	}
	return queue.elementData[queue.head]
}

func (queue *ArrayDeque[E]) Poll() E {
	if queue.head == queue.tail {
		return nil
	}
	e := queue.elementData[queue.head]
	queue.elementData[queue.head] = nil
	queue.head++
	return e
}

func (queue *ArrayDeque[E]) RemoveFirst() (E, error) {
	if queue.head == queue.tail {
		return nil, errors.New("NoSuchElementException")
	}
	e := queue.elementData[queue.head]
	queue.elementData[queue.head] = nil
	queue.head++
	return e, nil
}
