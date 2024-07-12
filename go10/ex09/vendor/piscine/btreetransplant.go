package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data string
}

func BTreeTransparent(root, node, rplc *TreeNode) *TreeNode {

}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
}

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
}

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
}
