package main

import (
	"errors"
	"fmt"
)

type Cell struct {
	element  string
	next     *Cell
	previous *Cell
}

type LinkedList struct {
	head *Cell
	tail *Cell
	len  int
}

func (l *LinkedList) Print() {
	if l.len == 0 {
		fmt.Println("No nodes in list")
	}

	pointer := l.head

	for i := 0; i < l.len; i++ {
		fmt.Println("Cell: ", pointer.element)
		pointer = pointer.next
	}
}

func (l *LinkedList) AddAtBegin(el string) {
	c := Cell{element: el}

	if l.len == 0 {
		l.head = &c
		l.tail = &c
	} else {
		head := l.head
		l.head = &c
		c.next = head
		head.previous = &c
	}

	l.len++
}

func (l *LinkedList) Add(el string) {
	c := Cell{element: el}

	if l.len == 0 {
		l.AddAtBegin(el)
		return
	}

	l.tail.next = &c
	c.previous = l.tail
	l.tail = &c
	l.len++
}

func (l *LinkedList) GetCell(position int) (*Cell, error) {
	if position < 0 || position > l.len {
		return nil, errors.New("inexistent position")
	}

	current := l.head

	for i := 0; i < position; i++ {
		current = current.next
	}

	return current, nil
}

func (l *LinkedList) AddAt(position int, el string) {
	if position == 0 {
		l.AddAtBegin(el)
		return
	}

	if position == l.len {
		l.Add(el)
		return
	}

	previous, _ := l.GetCell(position - 1)
	next := previous.next

	c := Cell{element: el}
	c.next = next
	c.previous = previous

	previous.next = &c
	next.previous = &c

	l.len++
}

func (l *LinkedList) RemoveFromBegin() error {
	if l.len == 0 {
		return errors.New("inexistent position")
	}

	l.head = l.head.next
	l.len--

	if l.len == 0 {
		l.tail = nil
	}

	return nil
}

func (l *LinkedList) Remove() {
	if l.len == 1 {
		l.RemoveFromBegin()
		return
	}

	penultimate := l.tail.previous
	penultimate.next = nil
	l.tail = penultimate
	l.len--
}

func (l *LinkedList) RemoveAt(position int) {
	if position == 0 {
		l.RemoveFromBegin()
		return
	}

	if position == l.len-1 {
		l.Remove()
		return
	}

	previous, _ := l.GetCell(position - 1)
	current := previous.next
	next := current.next

	previous.next = next
	next.previous = previous

	l.len--
}

func (l *LinkedList) Contains(element string) bool {
	current := l.head

	for current != nil {
		if current.element == element {
			return true
		}

		current = current.next
	}

	return false
}

func main() {
	list := LinkedList{}

	list.Add("2")
	list.AddAtBegin("1")
	list.AddAtBegin("0")
	list.Add("3")
	list.Add("4")
	list.Add("5")
	list.AddAt(1, "1.1")
	list.AddAt(4, "2.1")
	list.RemoveFromBegin()
	list.Remove()
	list.RemoveAt(1)

	fmt.Println(list.Contains("1"))
	fmt.Println(list.Contains("3"))

	list.Print()
}
