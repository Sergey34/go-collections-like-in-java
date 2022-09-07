package list

import (
	"errors"
	"fmt"
	"go-collections-like-in-java/collections"
	"reflect"
)

type LinkedList[E any] struct {
	first *Node[E]
	last  *Node[E]
	size  int
}

func (list *LinkedList[E]) Element() (E, error) {
	return list.Get(0)
}

func (list *LinkedList[E]) Offer(e E) (bool, error) {
	list.Add(e)
	return true, nil
}

func (list *LinkedList[E]) Peek() E {
	get, err := list.Get(0)
	if err != nil {
		var null E
		return null
	}
	return get
}

func (list *LinkedList[E]) Poll() E {
	get, err := list.Remove(0)
	if err != nil {
		var null E
		return null
	}
	return get
}

func (list *LinkedList[E]) RemoveFirst() (E, error) {
	return list.Remove(0)
}

type Node[E any] struct {
	data     *E
	next     *Node[E]
	previous *Node[E]
}

func NewLinkedList[E any]() LinkedList[E] {
	return LinkedList[E]{size: 0}
}

func (list *LinkedList[E]) Iterator() []E {
	result := make([]E, 0)
	if list.size != 0 {
		var current = list.first
		for i := 0; i < list.size; i++ {
			data := current.data
			result = append(result, *data)
			current = current.next
		}
	}
	return result
}

func (list *LinkedList[E]) Add(e E) {
	if list.size == 0 {
		newNode := Node[E]{
			data:     &e,
			next:     nil,
			previous: nil,
		}
		list.first = &newNode
		list.last = &newNode
		list.size++
		return
	}
	newNode := Node[E]{
		data:     &e,
		next:     nil,
		previous: list.last,
	}
	if list.size == 1 {
		list.last = &newNode
		list.first.next = &newNode
	} else {
		list.last.next = &newNode
		list.last = &newNode
	}
	list.size++
}

func (list *LinkedList[E]) Clear() {
	list.size = 0
	var null Node[E]
	list.first = &null
	list.last = &null
}

func (list *LinkedList[E]) Contains(o E) bool {
	var current = list.first
	for i := 0; i < list.size; i++ {
		data := current.data
		if reflect.DeepEqual(*data, o) {
			return true
		}
		current = current.next
	}
	return false
}

func (list *LinkedList[E]) ContainsAll(c Iterable[E]) bool {
	for _, v := range c.Iterator() {
		if !list.Contains(v) {
			return false
		}
	}
	return true
}

func (list *LinkedList[E]) IsEmpty() bool {
	return list.size == 0
}

func (list *LinkedList[E]) RemoveEntity(o any) bool {
	var current = list.first
	for i := 0; i < list.size; i++ {
		data := current.data
		if reflect.DeepEqual(*data, o) {
			if i == 0 {
				list.first = current.next
				current.previous = nil
			} else {
				current.previous.next = current.next
			}
			list.size--
			return true
		}
		current = current.next
	}
	return false
}

func (list *LinkedList[E]) RemoveAll(c Iterable[E]) bool {
	var result = false
	for _, v := range c.Iterator() {
		removed := list.RemoveEntity(v)
		if removed {
			result = true
		}
	}
	return result
}

func (list *LinkedList[E]) Size() int {
	return list.size
}

func (list *LinkedList[E]) Get(index int) (E, error) {
	if index >= list.size {
		var result E
		return result, errors.New(fmt.Sprintf("index out of range [%v] with length %v", index, list.size))
	}
	if index == 0 {
		return *list.first.data, nil
	}
	if index == list.size {
		return *list.last.data, nil
	}
	current := list.first
	for i := 0; i < index; i++ {
		current = current.next
	}
	return *current.data, nil
}

func (list *LinkedList[E]) getNode(index int) (Node[E], error) {
	if index >= list.size {
		var result Node[E]
		return result, errors.New(fmt.Sprintf("index out of range [%v] with length %v", index, list.size))
	}
	if index == 0 {
		return *list.first, nil
	}
	if index == list.size {
		return *list.last, nil
	}
	current := list.first
	for i := 0; i < index; i++ {
		current = current.next
	}
	return *current, nil
}

func (list *LinkedList[E]) IndexOf(o any) int {
	current := list.first
	for i := 0; i < list.size; i++ {
		if reflect.DeepEqual(*current.data, o) {
			return i
		}
		current = current.next
	}
	return -1
}

func (list *LinkedList[E]) LastIndexOf(o any) int {
	current := list.last
	for i := list.size - 1; i >= 0; i-- {
		if reflect.DeepEqual(*current.data, o) {
			return i
		}
		current = current.previous
	}
	return -1
}

func (list *LinkedList[E]) Remove(index int) (E, error) {
	size := list.size
	if index >= size {
		var result E
		return result, errors.New(fmt.Sprintf("index out of range [%v] with length %v", index, size))
	}
	node, e := list.getNode(index)
	if e != nil {
		var result E
		return result, e
	}
	node.previous.next = node.next
	list.size--
	return *node.data, nil
}

func (list *LinkedList[E]) Set(index int, element E) (E, error) {
	size := list.size
	var result E
	if size == index {
		list.Add(element)
		return result, nil
	}
	if size < index {
		return result, errors.New(fmt.Sprintf("index out of range [%v] with length %v", index, size))
	}
	node, e := list.getNode(index)
	if e != nil {
		var result E
		return result, e
	}
	oldData := *node.data
	*node.data = element
	return oldData, nil
}

func (list *LinkedList[E]) SubList(fromIndex int, toIndex int) (LinkedList[E], error) {
	if fromIndex > toIndex {
		var result LinkedList[E]
		return result, errors.New(fmt.Sprintf("toIndex should be > fromIndex. fromIndex: %v toIndex: %v", fromIndex, toIndex))
	}
	size := list.size
	if toIndex > size {
		var result LinkedList[E]
		return result, errors.New(fmt.Sprintf("index out of range [%v] with length %v", toIndex, size))
	}
	current := list.first
	var firstNode Node[E]
	for i := 0; i < toIndex; i++ {
		if i == fromIndex {
			firstNode = *current
		}
		current = current.next
	}
	firstNode.previous = nil
	if toIndex != list.size {
		current.next = nil
	}
	return LinkedList[E]{
		first: &firstNode,
		last:  current,
		size:  toIndex - fromIndex,
	}, nil
}

func (list *LinkedList[E]) ContainsByFilter(contains collections.Filter) bool {
	for _, v := range list.Iterator() {
		if contains(v) {
			return true
		}
	}
	return false
}

func (list *LinkedList[E]) Find(filter collections.Filter) (E, bool) {
	current := list.first
	for i := 0; i < list.size; i++ {
		if filter(current.data) {
			return *current.data, true
		}
		current = current.next
	}
	var null E
	return null, false
}
