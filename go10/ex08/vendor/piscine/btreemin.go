package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data string
}

// 一番左最深部のノードを返す
func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil {
		return root
	}
	return BTreeMin(root.Left)
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	if root == nil {
		root = &TreeNode{Data: data}
		return root
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
