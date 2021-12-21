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
func levelOrder(root *TreeNode) [][]int {

	queue := []*TreeNode{} // for queue since we using for bfs
	result := [][]int{}    // overall result , made up of levelResult
	count := 1             // we start with 1 and will assume root has value

	if root == nil {
		return nil // means there is nothing to traverse at all
	}

	queue = append(queue, root)

	for len(queue) > 0 {
		levelResult := []int{} //resets levelResult
		levelCount := 0

		for count > 0 {
			popped := queue[0]
			queue = queue[1:] //removes the first element from the queue
			count--

			if popped.Left != nil {
				queue = append(queue, popped.Left)
				levelCount++
			}

			if popped.Right != nil {
				queue = append(queue, popped.Right)
				levelCount++
			}

			levelResult = append(levelResult, popped.Val)
		}
		count = levelCount
		result = append(result, levelResult)
	}

	return result

}
