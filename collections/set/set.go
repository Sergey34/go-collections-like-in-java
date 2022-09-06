package set

import "go-collections-like-in-java/collections/java_map"

type Set[E java_map.Hashable] interface {
	Add(e E)
	Clear()
	Contains(o E) bool
	IsEmpty() bool
	Iterator() []E
	Remove(o E)
	Size() int
}
