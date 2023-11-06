package linked_list

import "fmt"

type ILinkedList[T interface{}] interface {
	Length() uint
	InsertAt(item T, index int)
	RemoveAt(index int) (T, error)
	Append(item T)
	Prepend(item T)
	Get(index int) (T, error)
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
	ll.len++
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

func (ll *SingleLinkedList[T]) Prepend(item T) {
	newNode := &Node[T]{value: item}
	ll.len++

	if ll.head == nil {
		ll.head = newNode
		return
	}

	newNode.next = ll.head
	ll.head = newNode
}

func (ll *SingleLinkedList[T]) Get(index int) (*T, error) {
	curr := ll.head

	for i := 0; i < index; i++ {
		if curr == nil {
			return nil, fmt.Errorf("Value not found")
		}
		curr = curr.next
	}

	return &curr.value, nil
}
