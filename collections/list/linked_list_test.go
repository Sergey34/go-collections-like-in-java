package list_test

import (
	"github.com/voicera/tester/assert"
	"go-collections-like-in-java/collections/list"
	"testing"
)

func TestLinkedList_Add(t *testing.T) {
	l := list.NewLinkedList[int]()
	l.Add(42)
	assert.For(t).ThatActual(l.Size()).Equals(1)
	get, _ := l.Get(0)
	assert.For(t).ThatActual(get).Equals(42)
}

func TestLinkedList_MultiAdd(t *testing.T) {
	l := list.NewLinkedList[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)
	assert.For(t).ThatActual(l.Size()).Equals(3)
	get, _ := l.Get(1)
	assert.For(t).ThatActual(get).Equals(2)
}

func TestLinkedList_GetOutOfBound(t *testing.T) {
	l := list.NewLinkedList[int]()
	l.Add(42)
	l.Add(42)
	l.Add(42)
	_, err := l.Get(4)
	assert.For(t).ThatActualError(err).IsNotNil()
}

func TestLinkedList_Clear(t *testing.T) {
	l := list.NewLinkedList[int]()
	l.Add(42)
	l.Add(42)
	l.Add(42)
	assert.For(t).ThatActual(l.Size()).Equals(3)
	l.Clear()
	assert.For(t).ThatActual(l.Size()).Equals(0)
	empty := l.IsEmpty()
	assert.For(t).ThatActual(empty).Equals(true)
}

func TestLinkedList_Set(t *testing.T) {
	l := list.NewLinkedList[int]()
	l.Add(42)
	l.Add(42)
	l.Add(42)
	assert.For(t).ThatActual(l.Size()).Equals(3)
	oldValue, e := l.Set(1, 22)
	assert.For(t).ThatActual(e).IsNil()
	assert.For(t).ThatActual(oldValue).Equals(42)
	get, _ := l.Get(1)
	assert.For(t).ThatActual(get).Equals(22)
}

func TestLinkedList_SetNoOldValue(t *testing.T) {
	l := list.NewLinkedList[int]()
	l.Add(42)
	l.Add(42)
	l.Add(42)
	assert.For(t).ThatActual(l.Size()).Equals(3)
	_, e := l.Set(3, 22)
	assert.For(t).ThatActual(e).IsNil()
	get, _ := l.Get(3)
	assert.For(t).ThatActual(get).Equals(22)
}

func TestLinkedList_IsEmpty(t *testing.T) {
	l := list.NewLinkedList[int]()
	l.Add(42)
	l.Add(42)
	l.Add(42)
	assert.For(t).ThatActual(l.Size()).Equals(3)
	empty := l.IsEmpty()
	assert.For(t).ThatActual(empty).IsFalse()
}

func TestLinkedList_Remove(t *testing.T) {
	l := list.NewLinkedList[int]()
	l.Add(42)
	l.Add(43)
	l.Add(44)
	assert.For(t).ThatActual(l.Size()).Equals(3)
	oldValue, e := l.Remove(1)
	assert.For(t).ThatActual(e).IsNil()
	assert.For(t).ThatActual(oldValue).Equals(43)
	assert.For(t).ThatActual(l.Size()).Equals(2)
}

func TestLinkedList_SubList(t *testing.T) {
	l := list.NewLinkedList[int]()
	l.Add(42)
	l.Add(43)
	l.Add(44)
	assert.For(t).ThatActual(l.Size()).Equals(3)
	newList, e := l.SubList(1, 3)
	assert.For(t).ThatActual(e).IsNil()
	assert.For(t).ThatActual(newList.Size()).Equals(2)
	assert.For(t).ThatActual(l.Size()).Equals(3)
}

func TestLinkedList_Contains(t *testing.T) {
	l := list.NewLinkedList[int]()
	l.Add(42)
	l.Add(43)
	l.Add(44)
	assert.For(t).ThatActual(l.Size()).Equals(3)
	assert.For(t).ThatActual(l.Contains(42)).IsTrue()
	assert.For(t).ThatActual(l.Contains(46)).IsFalse()
}

func TestLinkedList_ContainsAll(t *testing.T) {
	l := list.NewLinkedList[int]()
	l.Add(42)
	l.Add(43)
	l.Add(44)
	assert.For(t).ThatActual(l.Size()).Equals(3)

	ll := list.NewLinkedList[int]()
	ll.Add(42)
	ll.Add(43)
	assert.For(t).ThatActual(l.ContainsAll(&ll)).IsTrue()
	lll := list.NewLinkedList[int]()
	lll.Add(42)
	lll.Add(42)
	lll.Add(42)
	lll.Add(45)
	assert.For(t).ThatActual(l.ContainsAll(&lll)).IsFalse()
}

func TestLinkedList_IndexOf(t *testing.T) {
	l := list.NewLinkedList[int]()
	l.Add(42)
	l.Add(43)
	l.Add(44)
	l.Add(44)
	assert.For(t).ThatActual(l.Size()).Equals(4)
	assert.For(t).ThatActual(l.IndexOf(42)).Equals(0)
	assert.For(t).ThatActual(l.IndexOf(44)).Equals(2)
	assert.For(t).ThatActual(l.IndexOf(46)).Equals(-1)
}

func TestLinkedList_LastIndexOf(t *testing.T) {
	l := list.NewLinkedList[int]()
	l.Add(42)
	l.Add(43)
	l.Add(44)
	l.Add(44)
	assert.For(t).ThatActual(l.Size()).Equals(4)
	assert.For(t).ThatActual(l.LastIndexOf(42)).Equals(0)
	assert.For(t).ThatActual(l.LastIndexOf(44)).Equals(3)
	assert.For(t).ThatActual(l.LastIndexOf(46)).Equals(-1)
}

func TestLinkedList_RemoveEntity(t *testing.T) {
	l := list.NewLinkedList[int]()
	l.Add(42)
	l.Add(43)
	l.Add(44)
	l.Add(44)
	assert.For(t).ThatActual(l.Size()).Equals(4)
	assert.For(t).ThatActual(l.RemoveEntity(42)).Equals(true)
	assert.For(t).ThatActual(l.Size()).Equals(3)
	assert.For(t).ThatActual(l.RemoveEntity(44)).Equals(true)
	assert.For(t).ThatActual(l.Size()).Equals(2)
	assert.For(t).ThatActual(l.RemoveEntity(46)).Equals(false)
}

func TestLinkedList_RemoveAll(t *testing.T) {
	l := list.NewLinkedList[int]()
	l.Add(42)
	l.Add(43)
	l.Add(44)
	assert.For(t).ThatActual(l.Size()).Equals(3)

	ll := list.NewLinkedList[int]()
	ll.Add(42)
	ll.Add(43)
	assert.For(t).ThatActual(l.RemoveAll(&ll)).IsTrue()
	assert.For(t).ThatActual(l.Size()).Equals(1)
	lll := list.NewLinkedList[int]()
	lll.Add(42)
	lll.Add(42)
	lll.Add(44)
	lll.Add(45)
	assert.For(t).ThatActual(l.RemoveAll(&lll)).IsTrue()
	assert.For(t).ThatActual(l.Size()).Equals(0)
}
