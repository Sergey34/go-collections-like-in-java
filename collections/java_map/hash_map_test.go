package java_map_test

import (
	"github.com/mitchellh/hashstructure/v2"
	"github.com/voicera/tester/assert"
	"go-collections-like-in-java/collections/java_map"
	"testing"
)

type User struct {
	name string
	age  int
}

func (u User) HashCode() int {
	result, _ := hashstructure.Hash(u, hashstructure.FormatV2, nil)
	return int(result)
}

func TestArrayList_Add(t *testing.T) {
	myMap := java_map.NewHashMap[User, string]()
	seko := User{
		name: "seko",
		age:  27,
	}
	oldValue := myMap.Put(seko, "asd")

	assert.For(t).ThatActual(oldValue).Equals("")
	assert.For(t).ThatActual(myMap.Size()).Equals(1)
	get, found := myMap.Get(seko)
	assert.For(t).ThatActual(found).IsTrue()
	assert.For(t).ThatActual(get).Equals("asd")
}
