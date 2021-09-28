package leetcode113
//offer 34
//二叉树搜索，回溯法：先序遍历+路径记录
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

//func pathSum(root *TreeNode, target int) [][]int {
//	if root==nil||target==0{
//		return nil
//	}
//	var(
//		res [][]int
//		path []int
//	)
//	dfs(root, 0, target, path,res)
//	return res
//
//}
//
//func dfs(root *TreeNode, sum int, target int, path []int, res [][]int){
//	if root==nil||target==0{
//		return
//	}
//	path=append(path, root.Val)
//	sum+=root.Val
//	if sum==target&&root.Left==nil&&root.Right==nil{
//		var temp []int
//		copy(temp,path)
//		res= append(res, temp)
//	}
//	if root.Left!=nil{
//		dfs(root.Left,sum, target, path, res)
//	}
//	if root.Right!=nil{
//		dfs(root.Right,sum, target,path,res)
//	}
//}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, target int) [][]int {
	res := make([][]int, 0, 0)
	if root == nil {
		return res
	}

	return dfs(root, target, []int{}, res)
}
// root 当前节点
// target 当前值
// path   当前走的路径
// lastRes 当前路径结果集
// return: 遍历后的结果集
func dfs(root *TreeNode, target int, path []int, lastRes [][]int) [][]int {
	if root == nil {
		return lastRes
	}
	// 计算新的值
	target = target - root.Val
	// 添加路径
	path = append(path, root.Val)
	// 如果找到了路径，并且是根节点就添加进结果集
	if target == 0 && root.Left == nil && root.Right == nil {
		// var tempPath []int
		// tempPath = append(tempPath, path...)
		// lastRes = append(lastRes, tempPath)
		// 添加path到结果集
		lastRes = append(lastRes, append([]int{}, path...))
	}

	// 先序遍历 依次为 根-> 左-> 右
	lastRes = dfs(root.Left, target, path, lastRes)
	lastRes = dfs(root.Right, target, path, lastRes)

	// 移除遍历后的路径
	path = path[:len(path) - 1]

	return lastRes
}











