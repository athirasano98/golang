package main

import "fmt"

type List struct {
	head *Node
}
type Node struct {
	prev *Node
	data interface{}
	next *Node
}

func (list *List) Insert(data interface{}) {
	l := &Node{nil, data, nil}
	if list.head == nil {
		list.head = l
	} else {
		p := list.head
		for p.next != nil {
			p = p.next
		}
		p.next = l
		l.prev = p
	}

}
func (list *List) toCircular() {
	p := list.head
	for p.next != nil {
		p = p.next
	}
	p.next = list.head
	list.head.prev = p

}
func (list *List) Show() {
	p := list.head
	for {
		if p.next == list.head {
			fmt.Print("->", p.data)
			break

		}
		fmt.Print("->", p.data)
		p = p.next
	}
	fmt.Println()
}
func (list *List) Delete(item interface{}) {
	p := list.head
	if p.data == item {
		list.head = p.next
		return
	}
	for p != nil {
		if p.data == item {
			p.next.prev = p.prev
			p.prev.next = p.next
			return
		}
		p = p.next
	}

}
func main() {
	var list List
	list.Insert(10)
	list.Insert(20)
	list.Insert(30)
	list.Insert(40)
	list.Insert(50)
	list.toCircular()
	list.Show()
	list.Delete(10)
	fmt.Println("after deletion")
	list.Show()

}
