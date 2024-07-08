package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

type List struct {
	Head *NodeI
	Tail *NodeI
}

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	result := &List{}
	inserted := false
	curr := l
	for curr != nil {
		if data_ref <= curr.Data && inserted == false {
			ListPushBack(result, data_ref)
			ListPushBack(result, curr.Data)
			inserted = true
		} else {
			ListPushBack(result, curr.Data)
		}
		curr = curr.Next
	}
	return result.Head
}

func ListPushBack(l *List, data int) {
	n := &NodeI{Data: data}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}
	l.Tail.Next = n
	l.Tail = n
}
