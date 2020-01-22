/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func widthOfBinaryTree(root *TreeNode) int {
    return levelTraverse(root)
}

func levelTraverse(root *TreeNode) int {
    if root == nil {
        return 0
    }
    max := 1
    queue1, queue2 := newQueue(), newQueue()
    queue1.push(&Node{root, 0})
    startIndex, endIndex := -1, -2
    for !queue1.empty() || !queue2.empty() {
        for !queue1.empty() {
            n := queue1.pop()
            if n == nil || n.node == nil {
                continue
            }
            if startIndex == -1 {
                startIndex = n.index
            }
            endIndex = n.index
            queue2.push(&Node{n.node.Left, 2*n.index})
            queue2.push(&Node{n.node.Right, 2*n.index+1})
        }
        count := endIndex - startIndex + 1
        if count > max {
            max = count
        }
        queue1, queue2 = queue2, queue1
        startIndex, endIndex = -1, -2
    }
    
    return max
}

type Node struct {
    node  *TreeNode
    index int
}

type queue struct {
    slice []*Node
}

func newQueue() queue {
    return queue{
        slice: nil,
    }
}

func (q *queue) push(n *Node) {
    q.slice = append(q.slice, n)
}

func (q *queue) pop() *Node {
    if len(q.slice) == 0 {
        return nil
    }
    result := q.slice[0]
    q.slice = q.slice[1:]
    return result
}

func (q *queue) length() int {
    return len(q.slice)
}

func (q *queue) empty() bool {
    return len(q.slice) == 0
}
