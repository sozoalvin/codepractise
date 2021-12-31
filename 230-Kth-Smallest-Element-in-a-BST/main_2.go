/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// time complexity O(k) // worse case O(n) if kth smallest element is the last biggest element
// space complexity O(h) /l worse case O(n) if you get a skewed tree

/*
this solution uses pointers in golang to track, store and evaluate
*/

func kthSmallest(root *TreeNode, k int) int {

	if root == nil { //we cannot small anything if a nil tree is provided
		return -1
	}

	resultArray := make([]int, 3) // you make a slice of size 3, to store 1) a variable to track/ increment, 2) store your actual value you are interested in finding, 3)the k value itself

	resultArray[2] = k // you can pass this variable along as a function argument if you like. But since we wil be using the array already, might as well use it too.

	dfs(root, &resultArray) //driver function

	return resultArray[1]
}

func dfs(root *TreeNode, resultArray *[]int) {

	if root == nil { //this base case stops your recursion
		return
	}

	dfs(root.Left, resultArray) //recursive call on the left

	/*we use resultArray[0] to evaluate if a function should continue recursing. Once kth smallest have been found, we can stop the recrusion thanks to the else portion of the code. If however kth smallest is not found, then we will count and assign the value of resultArray[0] and continue with the rest of the code*/

	if (*resultArray)[0] != (*resultArray)[2] {
		(*resultArray)[0]++
		if (*resultArray)[0] == (*resultArray)[2] { //we found the kth smallest
			(*resultArray)[1] = root.Val
			(*resultArray)[0] = (*resultArray)[2]
			return
		}
	} else {
		return // if your value has been found, why still need recurse the right side? we end it
	}
	dfs(root.Right, resultArray) //recursive call on the right
}
