package singlylinkedlist

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (ll *LinkedList) Get(index int) int {
	currentNode := ll.head
	current := -1
	i := 0

	for i <= index && currentNode.next != nil {
		if i == index {
			current = currentNode.data
		} else {
			currentNode = currentNode.next
		}
		i++
	}
	return current
}

func (ll *LinkedList) InsertHead(value int) {
	newNode := &Node{data: value}

	if ll.head == nil {
		ll.head = newNode
		ll.tail = ll.head
	} else {
		newNode.next = ll.head
		ll.head = newNode
	}
}

func (ll *LinkedList) InsertTail(value int) {
	newNode := &Node{data: value}
	ll.tail.next = newNode
	ll.tail = ll.tail.next
}

func (ll *LinkedList) Remove(index int) {
	currentNode := ll.head
	i := 0

	for i < index {
		currentNode = currentNode.next
		i++
	}
	currentNode.data = currentNode.next.data
	currentNode.next = currentNode.next.next
}

func (ll *LinkedList) GetValues() []int {
	var combinedList []int
	currentNode := ll.head

	for currentNode != nil {
		combinedList = append(combinedList, currentNode.data)
		currentNode = currentNode.next
	}
	return combinedList
}
