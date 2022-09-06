package java_map

type Map[K comparable, V any] interface {
	Clear()
	Compute(key K, remappingFunction any)
	ContainsKey(key K) bool
	Get(key K) (V, bool)
	GetOrDefault(key K, defaultValue V) V
	IsEmpty() bool
	KeySet() []K
	Put(key K, value V) V
	Remove(key K) V
	Size() int
	Values() []V
}
