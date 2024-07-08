package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	head := l
	var isSorted bool
	for {
		isSorted = true
		curr := head
		for curr.Next != nil {
			if curr.Data > curr.Next.Data {
				isSorted = false
				curr.Data, curr.Next.Data = curr.Next.Data, curr.Data
			}
			curr = curr.Next
		}
		if isSorted {
			break
		}
	}
	return head
}
