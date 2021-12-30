/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// time complexity - O(n)
// space complexity - O(n)
func kthSmallest(root *TreeNode, k int) int {

	if root == nil {
		return -1
	}

	resultArray := make([]int, 1)

	dfs(root, &resultArray, &k)
	return resultArray[0]
}

func dfs(root *TreeNode, resultArray *[]int, k *int) {

	if root == nil {
		return
	}

	dfs(root.Left, resultArray, k)
	fmt.Println(root.Val)
	*k--
	if *k == 0 {
		(*resultArray)[0] = root.Val
		return
	}
	// *resultArray = append(*resultArray, root.Val)
	dfs(root.Right, resultArray, k)

}

