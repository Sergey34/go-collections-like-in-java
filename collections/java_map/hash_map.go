package java_map

import (
	_ "go-collections-like-in-java/collections"
	"go-collections-like-in-java/collections/list"
	"reflect"
)

type HashMap[K Hashable, V any] struct {
	table []list.LinkedList[Node[K, V]]
	size  int
}

type Node[K Hashable, V any] struct {
	hash  int
	key   K
	value V
}

type Hashable interface {
	hashCode() int
}

func NewHashMap[K Hashable, V any]() HashMap[K, V] {
	var table = make([]list.LinkedList[Node[K, V]], 10)
	return HashMap[K, V]{table: table}
}

func (hashMap HashMap[K, V]) Clear() {
	hashMap.table = make([]list.LinkedList[Node[K, V]], 10)
}

func (hashMap HashMap[K, V]) ContainsKey(key K) bool {
	bucket := hashMap.getBucket(key)
	nodes := hashMap.table[bucket]
	var null list.LinkedList[Node[K, V]]
	if nodes == null {
		return false
	}
	return nodes.ContainsByFilter(func(node any) bool {
		n := node.(Node[K, V])
		return n.hash == key.hashCode() && reflect.DeepEqual(n.key, key)
	})
}

func (hashMap HashMap[K, V]) Get(key K) (V, bool) {
	bucket := hashMap.getBucket(key)
	nodes := hashMap.table[bucket]
	var null list.LinkedList[Node[K, V]]
	if nodes == null {
		return nil, false
	}
	node, found := nodes.Find(func(node any) bool {
		n := node.(Node[K, V])
		return n.hash == key.hashCode() && reflect.DeepEqual(n.key, key)
	})
	if found {
		return node.value, true
	}
	return nil, false
}
func (hashMap HashMap[K, V]) GetNode(key K) (Node[K, V], bool) {
	bucket := hashMap.getBucket(key)
	nodes := hashMap.table[bucket]
	var null list.LinkedList[Node[K, V]]
	var n Node[K, V]
	if nodes == null {
		return n, false
	}
	node, found := nodes.Find(func(node any) bool {
		n := node.(Node[K, V])
		return n.hash == key.hashCode() && reflect.DeepEqual(n.key, key)
	})
	if found {
		return node, true
	}
	return n, false
}

func (hashMap HashMap[K, V]) GetOrDefault(key K, defaultValue V) V {
	node, found := hashMap.Get(key)
	if found {
		return node
	}
	return defaultValue
}

func (hashMap HashMap[K, V]) IsEmpty() bool {
	return hashMap.size == 0
}

func (hashMap HashMap[K, V]) KeySet() []K {
	result := make([]K, 0, hashMap.size)
	for _, nodes := range hashMap.table {
		for _, node := range nodes.Iterator() {
			result = append(result, node.key)
		}
	}
	return result
}

func (hashMap HashMap[K, V]) Put(key K, value V) V {
	bucket := hashMap.getBucket(key)
	nodes := hashMap.table[bucket]
	oldNode, found := hashMap.GetNode(key)
	if found {
		oldValue := oldNode.value
		oldNode.value = value
		return oldValue
	}
	nodes.Add(Node[K, V]{
		hash:  key.hashCode(),
		key:   key,
		value: value,
	})
	return nil
}

func (hashMap HashMap[K, V]) Remove(key K) V {
	bucket := hashMap.getBucket(key)
	nodes := hashMap.table[bucket]
	node, found := nodes.Find(func(node any) bool {
		n := node.(Node[K, V])
		return n.hash == key.hashCode() && reflect.DeepEqual(n.key, key)
	})
	if found {
		nodes.RemoveEntity(node)
		return node.value
	}
	return nil
}

func (hashMap HashMap[K, V]) Size() int {
	return hashMap.size
}

func (hashMap HashMap[K, V]) Values() []V {
	result := make([]V, 0, hashMap.size)
	for _, nodes := range hashMap.table {
		for _, node := range nodes.Iterator() {
			result = append(result, node.value)
		}
	}
	return result
}

func (hashMap HashMap[K, V]) getBucket(key K) int {
	return key.hashCode() % hashMap.Size()
}
