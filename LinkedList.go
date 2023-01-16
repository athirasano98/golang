package main

import "fmt"

type List struct {
	head *Node
}

type Node struct {
	item interface{}
	next *Node
}

func (list *List) Insert(item interface{}) {
	l := &Node{item, nil}
	if list.head == nil {
		list.head = l

	} else {
		p := list.head
		for p.next != nil {
			p = p.next
		}
		p.next = l
	}

}
func (list *List) Show() {
	p := list.head
	for p != nil {
		fmt.Print("->", p.item)
		p = p.next
	}
	fmt.Println()
}

func (list *List) Delete(item interface{}) interface{} {
	var p = list.head
	var data interface{}
	if p.item == item {
		list.head = p.next
		return p.item
	}
	for p != nil {
		if p.next.item == item {
			p.next = p.next.next
			data = item
			break
		}
		p = p.next
	}
	return data

}

func main() {
	var l List
	l.Insert(10)
	l.Insert(20)
	l.Insert(30)
	l.Insert(40)
	d := l.Delete(40)
	fmt.Println("the deleted element is", d)
	l.Show()

}
