package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 修剪二叉树，使不存在 Val 为 0 且无孩子的节点
func pruneTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if root.Left != nil {
		root.Left = pruneTree(root.Left)
	}
	if root.Right != nil {
		root.Right = pruneTree(root.Right)
	}
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}
	return root
}
