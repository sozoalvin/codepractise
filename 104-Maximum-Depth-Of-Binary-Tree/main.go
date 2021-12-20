/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	count := 0
	result := dfs(root, count)
	return result
}

func dfs(root *TreeNode, count int) int {

	countL := 0
	countR := 0

	if root == nil {
		return count
	}

	count++

	if root.Left != nil {
		countL = dfs(root.Left, count)
	} else {
		countL = count // if there are no left, then we should return the count at the particular level
	}

	if root.Right != nil {
		countR = dfs(root.Right, count)
	} else {
		countR = count // if there are no left, then we should return the count at the particular level
	}

	// you can use math.Max and cast float64/int values to circumvent go's enforcement of math package's requirement of float64

	if countL > countR {
		return countL
	} else if countR > countL {
		return countR
	} else {
		return countL
	}

}
