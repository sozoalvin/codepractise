/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// time complexity O(n^2) and space complexity O(n^2))

func buildTree(preorder []int, inorder []int) *TreeNode {

	if len(preorder) != len(inorder) || len(preorder) == 0 || len(inorder) == 0 {
		return nil // not possible to build any tree
	}

	return dfs(preorder, inorder)
}

func dfs(preorder, inorder []int) *TreeNode {

	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}

	mid := findIndex(preorder[0], inorder) //helper function fto find index from preorder array

	root.Left = dfs(preorder[1:mid+1], inorder[:mid])
	root.Right = dfs(preorder[mid+1:], inorder[mid+1:])

	return root
}

func findIndex(i int, sequence []int) int {
	for k, v := range sequence {
		if v == i {
			return k
		}
	}
	return -1
}


