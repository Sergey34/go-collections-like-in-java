package set

import (
	"go-collections-like-in-java/collections/java_map"
	"go/types"
)

type HashSet[E java_map.Hashable] struct {
	hashMap java_map.HashMap[E, types.Struct]
	t       types.Struct
}

func NewHashSet[E java_map.Hashable]() HashSet[E] {
	var hashMap = java_map.NewHashMap[E, types.Struct]()
	return HashSet[E]{hashMap: hashMap, t: types.Struct{}}
}

func (h HashSet[E]) Add(e E) {
	h.hashMap.Put(e, h.t)
}

func (h HashSet[E]) Clear() {
	h.hashMap.Clear()
}

func (h HashSet[E]) Contains(o E) bool {
	return h.hashMap.ContainsKey(o)
}

func (h HashSet[E]) IsEmpty() bool {
	return h.hashMap.IsEmpty()
}

func (h HashSet[E]) Iterator() []E {
	return h.hashMap.KeySet()
}

func (h HashSet[E]) Remove(o E) {
	h.hashMap.Remove(o)
}

func (h HashSet[E]) Size() int {
	return h.hashMap.Size()
}
