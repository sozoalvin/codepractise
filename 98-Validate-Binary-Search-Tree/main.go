/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
time complexity of O(n) for worst case and best case because you have to visit every node
space complexity of this is O(h) for average case and for worse case where you get a tree like a linked list, you will have O(n) time complexity

Please leave questions if you have any.
*/

func isValidBST(root *TreeNode) bool {

	if root == nil { //if you dont have a valid tree, then surely you can't continue. Just return true
		return true
	}

	return dfs(root, math.Inf(-1), math.Inf(+1))

}

func dfs(root *TreeNode, min, max float64) bool {

	if root == nil { // base case to stop the recursion
		return true
	}

	if float64(root.Val) <= min || float64(root.Val) >= max { //this checks if there are any forms of violation. if any of these rules are violated, then it will return a false
		return false
	}

	resultL := dfs(root.Left, min, float64(root.Val))
	resultR := dfs(root.Right, float64(root.Val), max)

	return resultL && resultR // this ensures that at any point in time, if you have a recursive call that returns a false, we will always return a false by the end of all the recusive calls

}
