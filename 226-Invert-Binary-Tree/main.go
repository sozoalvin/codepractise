/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// time complexity : O(n) since we touch every node once
// space comlexity:
// worse case: skewed tree = o(n)
// avg case: normal tree = O(log n)

func invertTree(root *TreeNode) *TreeNode {

	dfs(root)
	return root

}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}

	if root.Left != nil {
		dfs(root.Left)
	}

	if root.Right != nil {
		dfs(root.Right)
	}

	root.Left, root.Right = root.Right, root.Left
}
