package list

import (
	"errors"
	"fmt"
	"reflect"
)

type ArrayList[E any] struct {
	elementData []E
}

func (arrayList *ArrayList[E]) Iterator() []E {
	return arrayList.elementData
}

func NewArrayList[E any]() ArrayList[E] {
	var a = make([]E, 0)
	return ArrayList[E]{elementData: a}
}

func (arrayList *ArrayList[E]) Add(e E) {
	arrayList.elementData = append(arrayList.elementData, e)
}

func (arrayList *ArrayList[E]) Clear() {
	arrayList.elementData = arrayList.elementData[:0]
}

func (arrayList *ArrayList[E]) Contains(o E) bool {
	for _, v := range arrayList.Iterator() {
		if reflect.DeepEqual(v, o) {
			return true
		}
	}
	return false
}

func (arrayList *ArrayList[E]) ContainsAll(c Iterable[E]) bool {
	for _, v := range c.Iterator() {
		if !arrayList.Contains(v) {
			return false
		}
	}
	return true
}

func (arrayList *ArrayList[E]) IsEmpty() bool {
	return len(arrayList.elementData) == 0
}

func (arrayList *ArrayList[E]) RemoveEntity(o any) bool {
	index := arrayList.IndexOf(o)
	if index < 0 {
		return false
	}
	arrayList.Remove(index)
	return true
}

func (arrayList *ArrayList[E]) RemoveAll(c Iterable[E]) bool {
	var result = false
	for _, v := range c.Iterator() {
		removed := arrayList.RemoveEntity(v)
		if removed {
			result = true
		}
	}
	return result
}

func (arrayList *ArrayList[E]) Size() int {
	return len(arrayList.elementData)
}

func (arrayList *ArrayList[E]) Get(index int) (E, error) {
	size := len(arrayList.elementData)
	if index >= size {
		var result E
		return result, errors.New(fmt.Sprintf("index out of range [%v] with length %v", index, size))
	}
	return arrayList.elementData[index], nil
}

func (arrayList *ArrayList[E]) IndexOf(o any) int {
	for i, v := range arrayList.Iterator() {
		if reflect.DeepEqual(v, o) {
			return i
		}
	}
	return -1
}

func (arrayList *ArrayList[E]) LastIndexOf(o any) int {
	for i := arrayList.Size() - 1; i >= 0; i-- {
		if reflect.DeepEqual(arrayList.elementData[i], o) {
			return i
		}
	}
	return -1
}

func (arrayList *ArrayList[E]) Remove(index int) (E, error) {
	size := len(arrayList.elementData)
	if index >= size {
		var result E
		return result, errors.New(fmt.Sprintf("index out of range [%v] with length %v", index, size))
	}
	oldValue := arrayList.elementData[index]
	arrayList.elementData = append(arrayList.elementData[:index], arrayList.elementData[index+1:]...)
	return oldValue, nil
}

func (arrayList *ArrayList[E]) Set(index int, element E) (E, error) {
	size := len(arrayList.elementData)
	var result E
	if size == index {
		arrayList.Add(element)
		return result, nil
	}
	if size < index {
		return result, errors.New(fmt.Sprintf("index out of range [%v] with length %v", index, size))
	}
	oldValue := arrayList.elementData[index]
	arrayList.elementData[index] = element
	return oldValue, nil
}

func (arrayList *ArrayList[E]) SubList(fromIndex int, toIndex int) (ArrayList[E], error) {
	if fromIndex > toIndex {
		var result ArrayList[E]
		return result, errors.New(fmt.Sprintf("toIndex should be > fromIndex. fromIndex: %v toIndex: %v", fromIndex, toIndex))
	}
	size := len(arrayList.elementData)
	if toIndex > size {
		var result ArrayList[E]
		return result, errors.New(fmt.Sprintf("index out of range [%v] with length %v", toIndex, size))
	}
	newList := arrayList.elementData[fromIndex:toIndex]
	return ArrayList[E]{elementData: newList}, nil
}
