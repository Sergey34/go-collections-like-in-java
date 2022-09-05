package list

import (
	"errors"
	"fmt"
	"reflect"
)

type LinkedList[E any] struct {
	first *Node[E]
	last  *Node[E]
	size  int
}

type Node[E any] struct {
	data     *E
	next     *Node[E]
	previous *Node[E]
}

func NewLinkedList[E any]() LinkedList[E] {
	return LinkedList[E]{size: 0}
}

func (linkedList *LinkedList[E]) Iterator() []E {
	result := make([]E, 0)
	if linkedList.size != 0 {
		var current = linkedList.first
		for i := 0; i < linkedList.size; i++ {
			data := current.data
			result = append(result, *data)
			current = current.next
		}
	}
	return result
}

func (linkedList *LinkedList[E]) Add(e E) {

	if linkedList.size == 0 {
		newNode := Node[E]{
			data:     &e,
			next:     nil,
			previous: nil,
		}
		linkedList.first = &newNode
		linkedList.last = &newNode
		linkedList.size++
		return
	}
	newNode := Node[E]{
		data:     &e,
		next:     nil,
		previous: linkedList.last,
	}
	if linkedList.size == 1 {
		linkedList.last = &newNode
		linkedList.first.next = &newNode
	} else {
		linkedList.last.next = &newNode
		linkedList.last = &newNode
	}
	linkedList.size++
}

func (linkedList *LinkedList[E]) Clear() {
	linkedList.size = 0
	var null Node[E]
	linkedList.first = &null
	linkedList.last = &null
}

func (linkedList *LinkedList[E]) Contains(o E) bool {
	var current = linkedList.first
	for i := 0; i < linkedList.size; i++ {
		data := current.data
		if reflect.DeepEqual(*data, o) {
			return true
		}
		current = current.next
	}
	return false
}

func (linkedList *LinkedList[E]) ContainsAll(c Iterable[E]) bool {
	for _, v := range c.Iterator() {
		if !linkedList.Contains(v) {
			return false
		}
	}
	return true
}

func (linkedList *LinkedList[E]) IsEmpty() bool {
	return linkedList.size == 0
}

func (linkedList *LinkedList[E]) RemoveEntity(o any) bool {
	var current = linkedList.first
	for i := 0; i < linkedList.size; i++ {
		data := current.data
		if reflect.DeepEqual(*data, o) {
			if i == 0 {
				linkedList.first = current.next
				current.previous = nil
			} else {
				current.previous.next = current.next
			}
			linkedList.size--
			return true
		}
		current = current.next
	}
	return false
}

func (linkedList *LinkedList[E]) RemoveAll(c Iterable[E]) bool {
	var result = false
	for _, v := range c.Iterator() {
		removed := linkedList.RemoveEntity(v)
		if removed {
			result = true
		}
	}
	return result
}

func (linkedList *LinkedList[E]) Size() int {
	return linkedList.size
}

func (linkedList *LinkedList[E]) Get(index int) (E, error) {
	if index >= linkedList.size {
		var result E
		return result, errors.New(fmt.Sprintf("index out of range [%v] with length %v", index, linkedList.size))
	}
	if index == 0 {
		return *linkedList.first.data, nil
	}
	if index == linkedList.size {
		return *linkedList.last.data, nil
	}
	current := linkedList.first
	for i := 0; i < index; i++ {
		current = current.next
	}
	return *current.data, nil
}

func (linkedList *LinkedList[E]) getNode(index int) (Node[E], error) {
	if index >= linkedList.size {
		var result Node[E]
		return result, errors.New(fmt.Sprintf("index out of range [%v] with length %v", index, linkedList.size))
	}
	if index == 0 {
		return *linkedList.first, nil
	}
	if index == linkedList.size {
		return *linkedList.last, nil
	}
	current := linkedList.first
	for i := 0; i < index; i++ {
		current = current.next
	}
	return *current, nil
}

func (linkedList *LinkedList[E]) IndexOf(o any) int {
	current := linkedList.first
	for i := 0; i < linkedList.size; i++ {
		if reflect.DeepEqual(*current.data, o) {
			return i
		}
		current = current.next
	}
	return -1
}

func (linkedList *LinkedList[E]) LastIndexOf(o any) int {
	current := linkedList.last
	for i := linkedList.size - 1; i >= 0; i-- {
		if reflect.DeepEqual(*current.data, o) {
			return i
		}
		current = current.previous
	}
	return -1
}

func (linkedList *LinkedList[E]) Remove(index int) (E, error) {
	size := linkedList.size
	if index >= size {
		var result E
		return result, errors.New(fmt.Sprintf("index out of range [%v] with length %v", index, size))
	}
	node, e := linkedList.getNode(index)
	if e != nil {
		var result E
		return result, e
	}
	node.previous.next = node.next
	linkedList.size--
	return *node.data, nil
}

func (linkedList *LinkedList[E]) Set(index int, element E) (E, error) {
	size := linkedList.size
	var result E
	if size == index {
		linkedList.Add(element)
		return result, nil
	}
	if size < index {
		return result, errors.New(fmt.Sprintf("index out of range [%v] with length %v", index, size))
	}
	node, e := linkedList.getNode(index)
	if e != nil {
		var result E
		return result, e
	}
	oldData := *node.data
	*node.data = element
	return oldData, nil
}

func (linkedList *LinkedList[E]) SubList(fromIndex int, toIndex int) (LinkedList[E], error) {
	if fromIndex > toIndex {
		var result LinkedList[E]
		return result, errors.New(fmt.Sprintf("toIndex should be > fromIndex. fromIndex: %v toIndex: %v", fromIndex, toIndex))
	}
	size := linkedList.size
	if toIndex > size {
		var result LinkedList[E]
		return result, errors.New(fmt.Sprintf("index out of range [%v] with length %v", toIndex, size))
	}
	current := linkedList.first
	var firstNode Node[E]
	for i := 0; i < toIndex; i++ {
		if i == fromIndex {
			firstNode = *current
		}
		current = current.next
	}
	firstNode.previous = nil
	if toIndex != linkedList.size {
		current.next = nil
	}
	return LinkedList[E]{
		first: &firstNode,
		last:  current,
		size:  toIndex - fromIndex,
	}, nil
}
