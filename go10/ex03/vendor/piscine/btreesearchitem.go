package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data string
}

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil || root.Data == elem {
		return root
	}
	if elem < root.Data {
		return BTreeSearchItem(root.Left, elem)
	}
	return BTreeSearchItem(root.Right, elem)
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		return &TreeNode{Data: data}
	}
	if data < root.Data {
		if root.Left == nil {
			root.Left = &TreeNode{Data: data, Parent: root}
		} else {
			BTreeInsertData(root.Left, data)
		}
	} else {
		if root.Right == nil {
			root.Right = &TreeNode{Data: data, Parent: root}
		} else {
			BTreeInsertData(root.Right, data)
		}
	}
	return root
}
