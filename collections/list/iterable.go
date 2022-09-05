package list

type Iterable[E any] interface {
	Iterator() []E
}
