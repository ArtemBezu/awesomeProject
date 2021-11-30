package linkedlist

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	list := Constructor()
	if list.head != nil {
		t.Error("List head is not nil")
	}
}

func TestMyLinkedList_AddAtHead(t *testing.T) {
	list := Constructor()
	list.AddAtHead(1)
	if list.head == nil || list.head.val != 1 {
		t.Error("List list is not 1")
	}

	list.AddAtHead(2)
	if list.head == nil || list.head.val != 2 {
		t.Error("List list is not 2")
	}
}

func TestMyLinkedList_AddAtIndex(t *testing.T) {
	list := Constructor()
	list.AddAtIndex(0, 1)
	if list.head == nil || list.head.val != 1 {
		t.Error("List list is not 1")
	}
	list.AddAtIndex(1, 2)
	if list.head == nil || list.head.next.val != 2 {
		t.Error("List tail is not 2")
	}
	list.AddAtIndex(3, 3)
	if list.head == nil || list.head.val != 1 || list.head.next.val != 2 || list.head.next.next != nil {
		t.Error("List tail is not 2")
	}
	list.AddAtIndex(0, 3)
	if list.head == nil || list.head.val != 3 {
		t.Error("List list is not 3")
	}
}

func TestMyLinkedList_AddAtTail(t *testing.T) {
	list := Constructor()
	list.AddAtTail(1)
	if list.head == nil || list.head.val != 1 {
		t.Error("List tail is not 1")
	}
	list.AddAtTail(2)
	if list.head == nil || list.head.next.val != 2 {
		t.Error("List tail is not 2")
	}
}

func TestMyLinkedList_DeleteAtIndex(t *testing.T) {
	list := Constructor()
	list.AddAtTail(1)
	list.DeleteAtIndex(0)
	if list.head != nil {
		t.Error("List list is not nil")
	}

	list = Constructor()
	list.AddAtTail(1)
	list.AddAtTail(2)
	list.DeleteAtIndex(0)
	if list.head == nil || list.head.val != 2 {
		t.Error("List head is not 2")
	}

	myLinkedList := Constructor()
	myLinkedList.AddAtHead(1)
	myLinkedList.AddAtTail(3)
	myLinkedList.AddAtIndex(1, 2) // linked list becomes 1->2->3
	myLinkedList.Get(1)           // return 2
	myLinkedList.DeleteAtIndex(1) // now the linked list is 1->3
	myLinkedList.Get(1)           // return 3
}

func TestMyLinkedList_Get(t *testing.T) {
	list := Constructor()
	res := list.Get(0)
	if res != -1 {
		t.Error("List get is not -1")
	}
	list.AddAtTail(1)
	res = list.Get(0)
	if res != 1 {
		t.Error("List get is not 1")
	}

}

func TestMyLinkedList_SiteCase(t *testing.T) {
	list := Constructor()
	list.AddAtHead(2)
	list.DeleteAtIndex(1)
	list.AddAtHead(2)
	list.AddAtHead(7)
	list.AddAtHead(3)
	list.AddAtHead(2)
	list.AddAtHead(5)
	list.AddAtTail(5)
	list.Get(5)
	list.DeleteAtIndex(6)
	list.DeleteAtIndex(4)
}
