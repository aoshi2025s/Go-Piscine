package piscine

// for debug
// import "fmt"

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data string
}

// 右の最新部のノードがmaxのはず、、

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil || root.Right == nil {
		return root 
	}
	return BTreeMax(root.Right)
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
    if root == nil {
		// fmt.Println("rootノードに", data, "を挿入")
        return &TreeNode{Data: data}
    }

    curr := root

    for curr != nil {
		// fmt.Println("curr.Data: ", curr.Data, "とdata: ", data, "を比較")
        if data < curr.Data {
            if curr.Left == nil {
                curr.Left = &TreeNode{Data: data, Parent: curr}
				// fmt.Println(data, "<", curr.Data, "なので", curr.Data, "のLeftNodeに", data, "を挿入")
                return root
            } else {
                curr = curr.Left
            }
        } else {
            if curr.Right == nil {
                curr.Right = &TreeNode{Data: data, Parent: curr}
				// fmt.Println(data, ">=", curr.Data, "なので", curr.Data, "のRightNodeに", data, "を挿入")
                return root
            } else {
                curr = curr.Right
            }
        }
    }
    return root
}
/*
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
*/
