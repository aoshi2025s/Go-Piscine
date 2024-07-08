package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool ) *interface{} {
	curr := l.Head
	for curr != nil {
		if comp(curr.Data, ref) {
			return &(curr.Data)
		}
		curr = curr.Next
	}
	return nil
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
	}
	l.Tail.Next = n
	l.Tail = n
}
