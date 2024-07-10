package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data string
}

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftCount := BTreeLevelCount(root.Left)
	rightCount := BTreeLevelCount(root.Right)
	if leftCount > rightCount {
		return leftCount + 1
	}
	return rightCount + 1
}

/*
BTreeLevelCount(1) -> 0 + 1 + 1 + 1 = 3
	BTreeLevelCount(2) -> 0 + 1 + 1 = 2
		BTreeLevelCount(5) -> 0 + 1 = 1
			BTreeLevelCount(nil) -> 0
		BTreeLevelCount(6) -> 0 + 1
			BTreeLevelCount(nil) -> 0
	BTreeLevelCount(3) -> 0 + 1 + 1 = 1
		BTreeLevelCount(7) -> 0 + 1 = 1
			BTreeLevelCount(nil) -> 0
		BTreeLevelCount(8) -> 0 + 1 = 1
			BTreeLevelCount(nil) -> 0
	1
  /   \
  2   3
/  \ / \
5  6 7  8
*/

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	// base caaaase ??
	if root == nil {
		root := &TreeNode{Data: data}
		return root
	}
	
	// insert left if data < root.data
	if data < root.Data {
		if root.Left == nil {
			root.Left = &TreeNode{Data: data, Parent: root}
		} else {
			BTreeInsertData(root.Left, data)
		}
	}

	// insert right if data >= root.data
	if data >= root.Data {
		if root.Right == nil {
			root.Right = &TreeNode{Data: data, Parent: root}
		} else {
			BTreeInsertData(root.Right, data)
		}
	}

	return root
}
