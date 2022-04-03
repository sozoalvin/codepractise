// time complexity O(m * n) space complexity O(min(m + n))
package main

func main() {

}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return false
	}

	if isSameTree(root, subRoot) {
		return true
	} else {

		return isSubtree(root.Left, subRoot) || isSubTree(root.Right, subRoot)

	}
}

func isSameTree(root, subRoot *TreeNode) bool {

	if root == nil || subRoot == nil {
		return root == subRoot
	}

	if root.Val == subRoot.Val {
		resultL := isSameTree(root.Left, subRoot.Left)
		resultR := isSameTree(root.Right, subRoot.Right)

		return resultL && resultR
	} else {
		return false
	}
}
