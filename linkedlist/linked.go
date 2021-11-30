package linkedlist

type MyLinkedList struct {
	head *MyLinkedListNode
}

type MyLinkedListNode struct {
	next *MyLinkedListNode
	val  int
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (list *MyLinkedList) Get(index int) int {
	node := list.head
	for i := 0; i < index && node != nil; i++ {
		node = node.next
	}
	if node == nil {
		return -1
	} else {
		return node.val
	}
}

func (list *MyLinkedList) AddAtHead(val int) {
	list.head = &MyLinkedListNode{list.head, val}
}

func (list *MyLinkedList) AddAtTail(val int) {
	if list.head == nil {
		list.head = &MyLinkedListNode{nil, val}
		return
	}

	node := list.head
	for node.next != nil {
		node = node.next
	}
	node.next = &MyLinkedListNode{nil, val}
}

func (list *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		list.AddAtHead(val)
		return
	}

	node := list.head
	var i int
	for i = 0; i < index-1 && node != nil; i++ {
		node = node.next
	}
	if node == nil {
		return
	} else if node.next == nil && i+1 == index {
		list.AddAtTail(val)
	} else {
		node.next = &MyLinkedListNode{node.next, val}
	}
}

func (list *MyLinkedList) DeleteAtIndex(index int) {
	node := list.head
	if list.head == nil {
		return
	} else if index == 0 {
		list.head = list.head.next
		return
	}

	for i := 0; i < index-1 && node != nil; i++ {
		node = node.next
	}
	if node == nil || node.next == nil {
		return
	} else {
		node.next = node.next.next
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
