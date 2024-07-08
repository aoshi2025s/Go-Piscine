package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListRemoveIf(l *List, data_ref interface{}) {
	curr := l.Head
	result := &List{}
	for curr != nil {
		if curr.Data != data_ref {
			n := &NodeL{Data: curr.Data}
			ListPushBack(result, n.Data)
		}
		curr = curr.Next
	}
	l.Head = result.Head
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}
	l.Tail.Next = n
	l.Tail = n
}
