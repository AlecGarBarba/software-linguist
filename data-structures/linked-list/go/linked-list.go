package linked_list

/*
*
 */
type ILinkedList[T interface{}] interface {
	Length() uint
	InsertAt(item T, index uint)
	RemoveAt(index uint) (T, error)
	Append(item T)
	Prepend(item T)
	Get(index uint) (T, error)
}

type Node[T interface{}] struct {
	value T
	next  *Node[T]
}

type SingleLinkedList[T interface{}] struct {
	head *Node[T]
	len  uint
}

func (ll *SingleLinkedList[T]) Length() uint {
	return ll.len
}

func (ll *SingleLinkedList[T]) Append(item T) {
	newNode := &Node[T]{value: item}
	ll.len = ll.len + 1
	if ll.head == nil {
		ll.head = newNode
		return
	}
	curr := ll.head

	for curr.next != nil {
		curr = curr.next
	}

	curr.next = newNode
}
