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
	p := list.head
	if list.head == nil {
		list.head = l
	} else {
		for p.next != nil {
			p = p.next
		}
		p.next = l
	}
}

func (list *List) toCircular() {
	p := list.head
	for p.next != nil {
		p = p.next
	}
	p.next = list.head
}
func (list *List) Show() {
	p := list.head
	for {
		if p.next == list.head {
			fmt.Print("->", p.item)
			break
		}
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
	var list List
	list.Insert(10)
	list.Insert(20)
	list.Insert(30)
	list.Insert(40)
	list.toCircular()
	list.Show()
	list.Delete(40)
	fmt.Println("after deletion")
	list.Show()

}
