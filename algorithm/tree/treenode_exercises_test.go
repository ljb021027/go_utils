package tree

import (
	"container/list"
	"fmt"
	"testing"
)

func TestDepth(t *testing.T) {
	//ints := []int{0, 1, 2, 3, 4, 5}
	ints := []int{0, 3, 9, 20, 0, 0, 15, 7}
	array := buildWithArray(ints)
	min, max := depth(array)
	fmt.Println(min, max) // 2,3
}

//例如：给定二叉树 [3,9,20,null,null,15,7],
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//则它的最小深度  2，最大深度 3。
//
//解题思路：这题比较简单，只不过是深度优先遍历的一种变形，只要递归求出左右子树的最大/最小深度即可，深度怎么求，每递归调用一次函数，深度加一。不难写出如下代码：
func depth(node *TreeNode) (min, max int) {
	return dfsMinDepth(node), dfsMaxDepth(node)
}

func dfsMinDepth(node *TreeNode) int {
	left := 0
	right := 0
	if node == nil {
		return 0
	}
	left = dfsMinDepth(node.Left) + 1
	right = dfsMinDepth(node.Right) + 1
	if left <= right {
		return left
	} else {
		return right
	}
}

func dfsMaxDepth(node *TreeNode) int {
	left := 0
	right := 0
	if node == nil {
		return 0
	}
	left = dfsMaxDepth(node.Left) + 1
	right = dfsMaxDepth(node.Right) + 1
	if left >= right {
		return left
	} else {
		return right
	}
}

//leetcode 102: 给你一个二叉树，请你返回其按层序遍历得到的节点值。（即逐层地，从左到右访问所有节点）。示例，给定二叉树：[3,9,20,null,null,15,7]
func TestLevelResult(t *testing.T) {
	//ints := []int{0, 3, 9, 20, 0, 0, 15, 7}
	ints := []int{0, 1, 2, 3, 4, 5, 6, 7}
	array := buildWithArray(ints)
	level := bfsLevelResult(array)
	fmt.Println(level) // [[1] [2 3] [4 5 6 7]]

	result := dfsLevelResult(array)
	fmt.Println(result)
}

//bfsLevel
func bfsLevelResult(node *TreeNode) [][]int {
	if node == nil {
		return [][]int{}
	}
	l := list.List{}
	l.PushFront(node)
	levelResult := make([][]int, 0)
	for l.Back() != nil {
		num := l.Len()
		curResult := make([]int, num)
		for i := 0; i < num; i++ {
			cur := l.Back().Value.(*TreeNode)
			l.Remove(l.Back())
			curResult[i] = cur.val
			if cur.Left != nil {
				l.PushFront(cur.Left)
			}
			if cur.Right != nil {
				l.PushFront(cur.Right)
			}
		}
		levelResult = append(levelResult, curResult)

	}
	return levelResult
}

func dfsLevelResult(node *TreeNode) map[int][]int {
	levelResult := make(map[int][]int, 0)
	dfsLevel(node, 0, levelResult)
	return levelResult
}

func dfsLevel(node *TreeNode, level int, m map[int][]int) {
	if node == nil {
		return
	}
	curLevel := m[level]
	if curLevel == nil {
		curLevel = make([]int, 0)
	}
	curLevel = append(curLevel, node.val)
	m[level] = curLevel
	dfsLevel(node.Left, level+1, m)
	dfsLevel(node.Right, level+1, m)

}

//https://leetcode-cn.com/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/submissions/
func Test(t *testing.T) {
	i := grid(36, 11, 15)
	fmt.Println(i)
}

func grid(m, n, k int) int {
	l := list.List{}
	node := GridNode{0, 0}
	l.PushFront(node)
	existMap := make(map[GridNode]struct{})
	for l.Back() != nil {
		cur := l.Back().Value.(GridNode)
		l.Remove(l.Back())
		if cur.calVal() > k {
			continue
		} else {
			_, ok := existMap[cur]
			if ok {
				continue
			} else {
				existMap[cur] = struct{}{}
			}
			if cur.x+1 < m {
				l.PushFront(GridNode{cur.x + 1, cur.y})
			}
			if cur.y+1 < n {
				l.PushFront(GridNode{cur.x, cur.y + 1})
			}
		}

	}
	result := 0
	for range existMap {
		result++
	}
	return result
}

type GridNode struct {
	x int
	y int
}

func (n GridNode) calVal() int {
	return cal(n.x, n.y)
}

func cal(x, y int) int {
	return calOne(x) + calOne(y)
}

func calOne(x int) int {
	xTotal := 0
	for x > 0 {
		xTotal += x % 10
		x = x / 10
	}
	return xTotal
}
