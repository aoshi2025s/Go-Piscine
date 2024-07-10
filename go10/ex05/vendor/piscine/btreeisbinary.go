package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data string
}

/*
左部分木のすべてのノードは親ノードよりも小さく
右部分木のすべてのノードは親ノードよりも大きいという条件を満たすかを判定
isBinary(root, nil, nil)で開始
左のノードにとって親ノードがmax値になり -> if node.Data >= max.Data: return false
右のノードにとって親ノードがmin値になる -> if node.Data <= min.Data: return false
*/

func BTreeIsBinary(root *TreeNode) bool {
	return isBinary(root, nil, nil)
}

func isBinary(node, min, max *TreeNode) bool {
	//base case
	if node == nil {
		return true
	}
	
	// check left node
	if min != nil && node.Data < min.Data {
		return false
	}

	// check right node
	if max != nil && node.Data > max.Data {
		return false
	}
	
	return isBinary(node.Left, min, node) && isBinary(node.Right, node, max)
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
	}

	if data >= root.Data {
		if root.Right == nil {
			root.Right = &TreeNode{Data: data, Parent: root}
		} else {
			BTreeInsertData(root.Right, data)
		}
	}

	return root
}
