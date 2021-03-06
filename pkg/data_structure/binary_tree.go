package data_structure

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//εεΊιε
func preorderTraversal(root *TreeNode) []int {
	var preorder func(node *TreeNode)
	var result []int
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	return result
}
