package list_test

import (
	"github.com/voicera/tester/assert"
	"go-collections-like-in-java/collections/list"
	"testing"
)

func TestArrayList_Add(t *testing.T) {
	l := list.NewArrayList[int]()
	l.Add(42)
	assert.For(t).ThatActual(l.Size()).Equals(1)
	get, _ := l.Get(0)
	assert.For(t).ThatActual(get).Equals(42)
}

func TestArrayList_MultiAdd(t *testing.T) {
	l := list.NewArrayList[int]()
	l.Add(42)
	l.Add(42)
	l.Add(42)
	assert.For(t).ThatActual(l.Size()).Equals(3)
	get, _ := l.Get(1)
	assert.For(t).ThatActual(get).Equals(42)
}

func TestArrayList_GetOutOfBound(t *testing.T) {
	l := list.NewArrayList[int]()
	l.Add(42)
	l.Add(42)
	l.Add(42)
	_, err := l.Get(4)
	assert.For(t).ThatActualError(err).IsNotNil()
}

func TestArrayList_Clear(t *testing.T) {
	l := list.NewArrayList[int]()
	l.Add(42)
	l.Add(42)
	l.Add(42)
	assert.For(t).ThatActual(l.Size()).Equals(3)
	l.Clear()
	assert.For(t).ThatActual(l.Size()).Equals(0)
	empty := l.IsEmpty()
	assert.For(t).ThatActual(empty).Equals(true)
}

func TestArrayList_Set(t *testing.T) {
	l := list.NewArrayList[int]()
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

func TestArrayList_SetNoOldValue(t *testing.T) {
	l := list.NewArrayList[int]()
	l.Add(42)
	l.Add(42)
	l.Add(42)
	assert.For(t).ThatActual(l.Size()).Equals(3)
	_, e := l.Set(3, 22)
	assert.For(t).ThatActual(e).IsNil()
	get, _ := l.Get(3)
	assert.For(t).ThatActual(get).Equals(22)
}

func TestArrayList_IsEmpty(t *testing.T) {
	l := list.NewArrayList[int]()
	l.Add(42)
	l.Add(42)
	l.Add(42)
	assert.For(t).ThatActual(l.Size()).Equals(3)
	empty := l.IsEmpty()
	assert.For(t).ThatActual(empty).IsFalse()
}

func TestArrayList_Remove(t *testing.T) {
	l := list.NewArrayList[int]()
	l.Add(42)
	l.Add(43)
	l.Add(44)
	assert.For(t).ThatActual(l.Size()).Equals(3)
	oldValue, e := l.Remove(1)
	assert.For(t).ThatActual(e).IsNil()
	assert.For(t).ThatActual(oldValue).Equals(43)
	assert.For(t).ThatActual(l.Size()).Equals(2)
}

func TestArrayList_SubList(t *testing.T) {
	l := list.NewArrayList[int]()
	l.Add(42)
	l.Add(43)
	l.Add(44)
	assert.For(t).ThatActual(l.Size()).Equals(3)
	newList, e := l.SubList(1, 3)
	assert.For(t).ThatActual(e).IsNil()
	assert.For(t).ThatActual(newList.Size()).Equals(2)
	assert.For(t).ThatActual(l.Size()).Equals(3)
}

func TestArrayList_Contains(t *testing.T) {
	l := list.NewArrayList[int]()
	l.Add(42)
	l.Add(43)
	l.Add(44)
	assert.For(t).ThatActual(l.Size()).Equals(3)
	assert.For(t).ThatActual(l.Contains(42)).IsTrue()
	assert.For(t).ThatActual(l.Contains(46)).IsFalse()
}

func TestArrayList_ContainsAll(t *testing.T) {
	l := list.NewArrayList[int]()
	l.Add(42)
	l.Add(43)
	l.Add(44)
	assert.For(t).ThatActual(l.Size()).Equals(3)

	ll := list.NewArrayList[int]()
	ll.Add(42)
	ll.Add(43)
	assert.For(t).ThatActual(l.ContainsAll(&ll)).IsTrue()
	lll := list.NewArrayList[int]()
	lll.Add(42)
	lll.Add(42)
	lll.Add(42)
	lll.Add(45)
	assert.For(t).ThatActual(l.ContainsAll(&lll)).IsFalse()
}

func TestArrayList_IndexOf(t *testing.T) {
	l := list.NewArrayList[int]()
	l.Add(42)
	l.Add(43)
	l.Add(44)
	l.Add(44)
	assert.For(t).ThatActual(l.Size()).Equals(4)
	assert.For(t).ThatActual(l.IndexOf(42)).Equals(0)
	assert.For(t).ThatActual(l.IndexOf(44)).Equals(2)
	assert.For(t).ThatActual(l.IndexOf(46)).Equals(-1)
}

func TestArrayList_LastIndexOf(t *testing.T) {
	l := list.NewArrayList[int]()
	l.Add(42)
	l.Add(43)
	l.Add(44)
	l.Add(44)
	assert.For(t).ThatActual(l.Size()).Equals(4)
	assert.For(t).ThatActual(l.LastIndexOf(42)).Equals(0)
	assert.For(t).ThatActual(l.LastIndexOf(44)).Equals(3)
	assert.For(t).ThatActual(l.LastIndexOf(46)).Equals(-1)
}

func TestArrayList_RemoveEntity(t *testing.T) {
	l := list.NewArrayList[int]()
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

func TestArrayList_RemoveAll(t *testing.T) {
	l := list.NewArrayList[int]()
	l.Add(42)
	l.Add(43)
	l.Add(44)
	assert.For(t).ThatActual(l.Size()).Equals(3)

	ll := list.NewArrayList[int]()
	ll.Add(42)
	ll.Add(43)
	assert.For(t).ThatActual(l.RemoveAll(&ll)).IsTrue()
	assert.For(t).ThatActual(l.Size()).Equals(1)
	lll := list.NewArrayList[int]()
	lll.Add(42)
	lll.Add(42)
	lll.Add(44)
	lll.Add(45)
	assert.For(t).ThatActual(l.RemoveAll(&lll)).IsTrue()
	assert.For(t).ThatActual(l.Size()).Equals(0)
}
