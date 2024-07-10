package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data string
}

// same meaning foreach function

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	// base case
	if root == nil {
		return
	}
	
	f(root.Data)
	BTreeApplyByLevel(root.Left, f)
	BTreeApplyByLevel(root.Right, f)
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
	} else if data >= root.Data {
		if root.Right == nil {
			root.Right = &TreeNode{Data: data, Parent: root}
		} else {
			BTreeInsertData(root.Right, data)
		}
	}
	return root
}
