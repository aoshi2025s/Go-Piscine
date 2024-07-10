package main

import (
	"fmt"
	"piscine"
)

func main() {

	root := &piscine.TreeNode{Data: "4"}
	piscine.BTreeInsertData(root, "1")
	piscine.BTreeInsertData(root, "7")
	piscine.BTreeInsertData(root, "5")
	fmt.Println(piscine.BTreeLevelCount(root))


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
	fmt.Println(piscine.BTreeLevelCount(root))

	// TODO: visualize BTree
}

