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
	values := []string{"1", "7", "5", "2", "9", "8", "3", "6"}
	for _, v := range values {
		root = piscine.BTreeInsertData(root, v)
	}
	fmt.Println("=============================")
	VisualizeBTree(root, "", false)
	fmt.Println("=============================")
	min := piscine.BTreeMin(root)
	fmt.Println("min: ", min.Data)
	fmt.Println("==============================")
}
