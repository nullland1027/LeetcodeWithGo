package dataStructure

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateTreeNode(data int) *TreeNode {
	return &TreeNode{data, nil, nil}
}

func PreorderTraverse(t *TreeNode, store []int) []int {
	if t != nil {
		store = append(store, t.Val)
		store = PreorderTraverse(t.Left, store)
		store = PreorderTraverse(t.Right, store)
	}
	return store
}

func InorderTraverse(t *TreeNode, store []int) []int {
	if t != nil {
		store = InorderTraverse(t.Left, store)
		store = append(store, t.Val)
		store = InorderTraverse(t.Right, store)
	}
	return store
}

func PostorderTraverse(t *TreeNode, store []int) []int {
	if t != nil {
		store = PostorderTraverse(t.Left, store)
		store = PostorderTraverse(t.Right, store)
		store = append(store, t.Val)
	}
	return store
}

func checkBST(node *TreeNode, min int, max int) bool {
	if node == nil {
		return true
	}
	value := node.Val
	return min < value && value < max && checkBST(node.Left, -math.MaxInt, value) && checkBST(node.Right, value, math.MaxInt)
}

func IsValidBST(root *TreeNode) bool {
	return checkBST(root, -math.MaxInt, math.MaxInt)
}

func IsValidBST2(root *TreeNode) bool {
	vals := InorderTraverse(root, []int{})
	for i := 0; i < len(vals)-1; i++ {
		if vals[i] >= vals[i+1] {
			return false
		}
	}
	return true
}

func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return make([][]int, 0)
	}
	current := []*TreeNode{root}
	next := []*TreeNode{}
	ans := [][]int{}
	for current != nil {
		layerAns := []int{}
		for _, node := range current {
			layerAns = append(layerAns, node.Val)
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		current = next
		next = []*TreeNode{}
		ans = append(ans, layerAns)
	}
	return ans
}

func buildTreeHelper(preorder []int, preStart, preEnd int, inorder []int, inStart, inEnd int, inMap map[int]int) *TreeNode {
	if preStart > preEnd {
		return nil
	}

	rootVal := preorder[preStart]
	root := &TreeNode{Val: rootVal}

	rootIndex := inMap[rootVal]
	leftSubtreeSize := rootIndex - inStart

	root.Left = buildTreeHelper(preorder, preStart+1, preStart+leftSubtreeSize,
		inorder, inStart, rootIndex-1, inMap)
	root.Right = buildTreeHelper(preorder, preStart+leftSubtreeSize+1, preEnd,
		inorder, rootIndex+1, inEnd, inMap)

	return root
}

func BuildTreeFromPreorderInorder(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// 创建中序遍历的值到索引的映射，用于O(1)时间找到根节点位置
	inMap := make(map[int]int)
	for i, val := range inorder {
		inMap[val] = i
	}

	return buildTreeHelper(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1, inMap)
}

