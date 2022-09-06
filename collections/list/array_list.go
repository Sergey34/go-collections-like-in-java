package list

import (
	"errors"
	"fmt"
	"go-collections-like-in-java/collections"
	"reflect"
)

type ArrayList[E any] struct {
	elementData []E
}

func (list *ArrayList[E]) Iterator() []E {
	return list.elementData
}

func NewArrayList[E any]() ArrayList[E] {
	var a = make([]E, 0)
	return ArrayList[E]{elementData: a}
}

func (list *ArrayList[E]) Add(e E) {
	list.elementData = append(list.elementData, e)
}

func (list *ArrayList[E]) Clear() {
	list.elementData = list.elementData[:0]
}

func (list *ArrayList[E]) Contains(o E) bool {
	for _, v := range list.Iterator() {
		if reflect.DeepEqual(v, o) {
			return true
		}
	}
	return false
}

func (list *ArrayList[E]) ContainsAll(c Iterable[E]) bool {
	for _, v := range c.Iterator() {
		if !list.Contains(v) {
			return false
		}
	}
	return true
}

func (list *ArrayList[E]) IsEmpty() bool {
	return len(list.elementData) == 0
}

func (list *ArrayList[E]) RemoveEntity(o any) bool {
	index := list.IndexOf(o)
	if index < 0 {
		return false
	}
	list.Remove(index)
	return true
}

func (list *ArrayList[E]) RemoveAll(c Iterable[E]) bool {
	var result = false
	for _, v := range c.Iterator() {
		removed := list.RemoveEntity(v)
		if removed {
			result = true
		}
	}
	return result
}

func (list *ArrayList[E]) Size() int {
	return len(list.elementData)
}

func (list *ArrayList[E]) Get(index int) (E, error) {
	size := len(list.elementData)
	if index >= size {
		var result E
		return result, errors.New(fmt.Sprintf("index out of range [%v] with length %v", index, size))
	}
	return list.elementData[index], nil
}

func (list *ArrayList[E]) IndexOf(o any) int {
	for i, v := range list.Iterator() {
		if reflect.DeepEqual(v, o) {
			return i
		}
	}
	return -1
}

func (list *ArrayList[E]) LastIndexOf(o any) int {
	for i := list.Size() - 1; i >= 0; i-- {
		if reflect.DeepEqual(list.elementData[i], o) {
			return i
		}
	}
	return -1
}

func (list *ArrayList[E]) Remove(index int) (E, error) {
	size := len(list.elementData)
	if index >= size {
		var result E
		return result, errors.New(fmt.Sprintf("index out of range [%v] with length %v", index, size))
	}
	oldValue := list.elementData[index]
	list.elementData = append(list.elementData[:index], list.elementData[index+1:]...)
	return oldValue, nil
}

func (list *ArrayList[E]) Set(index int, element E) (E, error) {
	size := len(list.elementData)
	var result E
	if size == index {
		list.Add(element)
		return result, nil
	}
	if size < index {
		return result, errors.New(fmt.Sprintf("index out of range [%v] with length %v", index, size))
	}
	oldValue := list.elementData[index]
	list.elementData[index] = element
	return oldValue, nil
}

func (list *ArrayList[E]) SubList(fromIndex int, toIndex int) (ArrayList[E], error) {
	if fromIndex > toIndex {
		var result ArrayList[E]
		return result, errors.New(fmt.Sprintf("toIndex should be > fromIndex. fromIndex: %v toIndex: %v", fromIndex, toIndex))
	}
	size := len(list.elementData)
	if toIndex > size {
		var result ArrayList[E]
		return result, errors.New(fmt.Sprintf("index out of range [%v] with length %v", toIndex, size))
	}
	newList := list.elementData[fromIndex:toIndex]
	return ArrayList[E]{elementData: newList}, nil
}

func (list *ArrayList[E]) ContainsByFilter(contains collections.Filter) bool {
	for _, v := range list.Iterator() {
		if contains(v) {
			return true
		}
	}
	return false
}

func (list *ArrayList[E]) Find(filter collections.Filter) (E, bool) {
	for _, v := range list.Iterator() {
		if filter(v) {
			return v, true
		}
	}
	return nil, false
}
