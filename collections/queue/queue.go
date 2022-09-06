package queue

type Queue[E any] interface {
	Add(e E) bool
	Element() (E, error)
	Offer(e E) (bool, error)
	Peek() E
	Poll() E
	RemoveFirst() (E, error)
}
