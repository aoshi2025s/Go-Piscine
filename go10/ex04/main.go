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
	values := []string{"4", "5", "3", "1", "2", "8", "7", "6", "9"}
	for _, v := range values {
		root = piscine.BTreeInsertData(root, v)
	}
	fmt.Printf("%s================================%s\n", pink, reset)
	VisualizeBTree(root, "", false)
	fmt.Printf("%s================================%s\n", pink, reset)
	fmt.Printf("%scount level%s %d\n", blue, reset, piscine.BTreeLevelCount(root))
	fmt.Printf("%s================================%s\n", pink, reset)
}

