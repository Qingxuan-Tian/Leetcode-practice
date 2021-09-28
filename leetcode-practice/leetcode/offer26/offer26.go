package offer26

import "../binaryTree"
//func isSame(A *binaryTree.TreeNode, B *binaryTree.TreeNode)bool{
//	if B==nil
//}
func isSame(A *binaryTree.TreeNode, B *binaryTree.TreeNode)bool {
	if B==nil{
		return true
	}
	if A==nil{
		return false
	}
	if A.Val!=B.Val{
		return false
	}
	return isSame(A.Right,B.Right)&&isSame(A.Left,B.Left)
}
func isSubStructure(A *binaryTree.TreeNode, B *binaryTree.TreeNode) bool {
	if A==nil&&B==nil{
		return true
	}
	if B==nil||A==nil{
		return false
	}
	var ret bool
	if A.Val==B.Val{
		ret=isSame(A.Left,B.Left)&&isSame(A.Right,B.Right)
		// ok:= isSubStructure(A.Left,B.Left)
		// ok1:= isSubStructure(A.Right,B.Right)
		// if ok&&ok1{
		// 	ret=true
		// }
	}
	if !ret{
		ret=isSubStructure(A.Right,B)||isSubStructure(A.Left,B)
	}
	//else {
	//	ok2:= isSubStructure(A.Right,B)
	//	ok3:= isSubStructure(A.Left,B)
	//	if ok2||ok3{
	//		return true
	//	}else {
	//		return false
	//	}
	//}
	return ret
}
