package set_test

import (
	"github.com/mitchellh/hashstructure/v2"
	"github.com/voicera/tester/assert"
	"go-collections-like-in-java/collections/set"
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
	mySet := set.NewHashSet[User]()
	seko := User{
		name: "seko",
		age:  27,
	}
	mySet.Add(seko)

	assert.For(t).ThatActual(mySet.Size()).Equals(1)
	size := len(mySet.Iterator())
	assert.For(t).ThatActual(size).Equals(1)

}
