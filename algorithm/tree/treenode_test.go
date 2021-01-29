package tree

import (
	"container/list"
	"fmt"
	"testing"
)

func build() *TreeNode {
	return &TreeNode{
		val: 1,
		Left: &TreeNode{
			val: 2,
			Left: &TreeNode{
				val:   5,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				val:   6,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			val:   4,
			Left:  nil,
			Right: nil,
		},
	}
}

// 从下标1开始存储，0表示空节点
func buildWithArray(arr []int) *TreeNode {
	m := map[int]*TreeNode{}
	for i := 1; i < len(arr); i++ {
		//int 0值当做空节点
		if arr[i] == 0 {
			continue
		}
		node := TreeNode{
			val:   arr[i],
			Left:  nil,
			Right: nil,
		}
		m[i] = &node
		if i/2 > 0 && i%2 == 0 {
			m[i/2].Left = &node
		}
		if (i-1)/2 > 0 && i%2 == 1 {
			m[(i-1)/2].Right = &node
		}
	}

	return m[1]
}

func TestBuildWithArray(t *testing.T) {
	array := buildWithArray([]int{0, 1, 2, 3, 4, 5, 6, 7})
	fmt.Println(array)
}

func TestDfs(t *testing.T) {
	array := buildWithArray([]int{0, 1, 2, 3, 4, 5, 6, 7})
	dfs(array)
	fmt.Println("===")
	dfsStack(array)
}

func TestBfs(t *testing.T) {
	array := buildWithArray([]int{0, 1, 2, 3, 4, 5, 6, 7})
	bfs(array)
}

func TestStack(t *testing.T) {
	stack := list.List{}

	stack.PushBack(1)
	stack.PushBack(2)
	stack.PushBack(3)

	for top := stack.Back(); top != nil; top = stack.Back() {
		stack.Remove(top)
		fmt.Println(top.Value)
	}

}
