package tree

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	val   int
	Left  *TreeNode
	Right *TreeNode
}

//dfs 递归实现
func dfs(node *TreeNode) {
	if node != nil {
		fmt.Println(node.val)
	}
	if node.Left != nil {
		dfs(node.Left)
	}
	if node.Right != nil {
		dfs(node.Right)
	}
}

//对于每个节点来说，先遍历当前节点，然后把右节点压栈，再压左节点（这样弹栈的时候会先拿到左节点遍历，符合深度优先遍历要求）
//弹栈，拿到栈顶的节点，如果节点不为空，重复步骤 1， 如果为空，结束遍历。
//dfs 栈实现
func dfsStack(node *TreeNode) {
	stack := list.List{}
	stack.PushBack(node)
	for stack.Back() != nil {
		cur := stack.Back().Value.(*TreeNode)
		stack.Remove(stack.Back())
		fmt.Println(cur.val)
		if cur.Right != nil {
			stack.PushBack(cur.Right)
		}
		if cur.Left != nil {
			stack.PushBack(cur.Left)
		}

	}
}

//bfs 队列实现
func bfs(node *TreeNode) {
	l := list.List{}
	l.PushFront(node)
	for l.Back() != nil {
		cur := l.Back().Value.(*TreeNode)
		l.Remove(l.Back())
		fmt.Println(cur.val)
		if cur.Left != nil {
			l.PushFront(cur.Left)
		}
		if cur.Right != nil {
			l.PushFront(cur.Right)
		}
	}

}
