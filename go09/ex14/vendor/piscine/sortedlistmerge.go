package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

type List struct {
	Head *NodeI
	Tail *NodeI
}

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	result := &List{}

	it1 := n1
	it2 := n2
	for it1 != nil && it2 != nil {
		if it1.Data < it2.Data {
			ListPushBack(result, it1.Data)
			it1 = it1.Next
		} else if it2.Data <= it1.Data {
			ListPushBack(result, it2.Data)
			it2 = it2.Next
		}
	}
	
	if it1 != nil && it2 == nil {
		for it1 != nil {
			ListPushBack(result, it1.Data)
			it1 = it1.Next
		}
	} else if it1 == nil && it2 != nil {
		for it2 != nil {
			ListPushBack(result, it2.Data)
			it2 = it2.Next
		}
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
