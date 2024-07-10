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
	fmt.Printf("%s================================%s\n", pink, reset)
	VisualizeBTree(root, "", false)
	fmt.Printf("%s================================%s\n", pink, reset)
	fmt.Printf("%scount level%s %d\n", blue, reset, piscine.BTreeLevelCount(root))
	fmt.Printf("%s================================%s\n", pink, reset)


	root = &piscine.TreeNode{Data: "1"}
	piscine.BTreeInsertData(root, "2")
	piscine.BTreeInsertData(root, "3")
	piscine.BTreeInsertData(root, "4")
	piscine.BTreeInsertData(root, "5")
	piscine.BTreeInsertData(root, "6")
	piscine.BTreeInsertData(root, "7")
	piscine.BTreeInsertData(root, "8")
	piscine.BTreeInsertData(root, "9")
	piscine.BTreeInsertData(root, "10")
	piscine.BTreeInsertData(root, "11")
	piscine.BTreeInsertData(root, "12")
	piscine.BTreeInsertData(root, "13")
	piscine.BTreeInsertData(root, "14")
	piscine.BTreeInsertData(root, "15")
	piscine.BTreeInsertData(root, "16")
	piscine.BTreeInsertData(root, "17")
	piscine.BTreeInsertData(root, "18")
	piscine.BTreeInsertData(root, "19")
	piscine.BTreeInsertData(root, "20")
	piscine.BTreeInsertData(root, "21")
	fmt.Printf("%s================================%s\n", pink, reset)
	VisualizeBTree(root, "", false)
	fmt.Printf("%s================================%s\n", pink, reset)
	fmt.Printf("%scount level%s %d\n", blue, reset, piscine.BTreeLevelCount(root))
	fmt.Printf("%s================================%s\n", pink, reset)

	// TODO: visualize BTree
}

