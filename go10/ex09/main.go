package main

import (
	"fmt"
	"piscine"
)

const (
	blue = "\033[34m"
	green = "\033[32m"
	pink  = "\033[35m"
	reset = "\033[0m"
)

func VisualizeBTree(root *piscine.TreeNode, prefix string, isLeft bool) {
	if root != nil {
		if root.Parent == nil {
			fmt.Printf("%s%s%s\n", pink, root.Data, reset)
		} else {
			if isLeft {
				fmt.Printf("%s%s├── %s%s%s%s\n", green,  prefix, reset,  pink, root.Data, reset)
			} else {
				fmt.Printf("%s%s└── %s%s%s%s\n", green, prefix, reset, pink, root.Data, reset)
			}
		}

		newPrefix := prefix
		if root.Parent != nil {
			if isLeft {
				newPrefix += "│   "
			} else {
				newPrefix += "    "
			}
		}

		VisualizeBTree(root.Left, newPrefix, true)
		VisualizeBTree(root.Right, newPrefix, false)
	}
}

func main() {
	root := &piscine.TreeNode{Data: "4"}
	piscine.BTreeInsertData(root, "1")
	piscine.BTreeInsertData(root, "7")
	piscine.BTreeInsertData(root, "5")
	node := BTreeSearchItem(root, "1")
	replacement := &piscine.TreeNode{Data, "3"}
	root = piscine.BTreeTransplant(root, node, replacement)
	piscine.BTreeApplyInorder(root, fmt.Println)
}

/*
go mod init ex09
go run .
3
4
5
7
*/
