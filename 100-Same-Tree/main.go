/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// time complexity O(n) since we have to visit every node
// space complexity:
// worse case is O(n) for a skewed tree
// avg case is O(log n ) for a avg tree since max recursion is height of tree

func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false // if either one is nil, comparison stops
	}

	result := dfs(p, q)
	return result

}

func dfs(p *TreeNode, q *TreeNode) bool {

	var res1 bool
	var res2 bool

	if p.Val != q.Val {
		return false
	}

	if p.Left != nil && q.Left != nil {
		res1 = dfs(p.Left, q.Left)
	} else if p.Left == nil && q.Left == nil {
		res1 = true
	}

	if p.Right != nil && q.Right != nil {
		res2 = dfs(p.Right, q.Right)
	} else if p.Right == nil && q.Right == nil {
		res2 = true
	}

	return res1 && res2

}

