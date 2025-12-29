package dbmsgo

import (
	"fmt"
)

type AVLNode struct {
	Data   int
	Left   *AVLNode
	Right  *AVLNode
	height int
}

type AVLTree struct {
	name string
	root *AVLNode
}

func NewAVLTree(name string) *AVLTree {
	return &AVLTree{
		name: name,
		root: nil,
	}
}

func (a *AVLTree) height(node *AVLNode) int {
	if node == nil {
		return 0
	}
	return node.height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (a *AVLTree) balanceFactor(node *AVLNode) int {
	if node == nil {
		return 0
	}
	return a.height(node.Left) - a.height(node.Right)
}

func (a *AVLTree) rotateRight(y *AVLNode) *AVLNode {
	x := y.Left
	T2 := x.Right

	x.Right = y
	y.Left = T2

	y.height = max(a.height(y.Left), a.height(y.Right)) + 1
	x.height = max(a.height(x.Left), a.height(x.Right)) + 1

	return x
}

func (a *AVLTree) rotateLeft(x *AVLNode) *AVLNode {
	y := x.Right
	T2 := y.Left

	y.Left = x
	x.Right = T2

	x.height = max(a.height(x.Left), a.height(x.Right)) + 1
	y.height = max(a.height(y.Left), a.height(y.Right)) + 1

	return y
}

func (a *AVLTree) insertHelper(node *AVLNode, value int) *AVLNode {
	if node == nil {
		return &AVLNode{Data: value, height: 1}
	}

	if value < node.Data {
		node.Left = a.insertHelper(node.Left, value)
	} else if value > node.Data {
		node.Right = a.insertHelper(node.Right, value)
	} else {
		return node // Дубликаты не разрешены
	}

	node.height = 1 + max(a.height(node.Left), a.height(node.Right))

	balance := a.balanceFactor(node)

	// Left Left Case
	if balance > 1 && value < node.Left.Data {
		return a.rotateRight(node)
	}

	// Right Right Case
	if balance < -1 && value > node.Right.Data {
		return a.rotateLeft(node)
	}

	// Left Right Case
	if balance > 1 && value > node.Left.Data {
		node.Left = a.rotateLeft(node.Left)
		return a.rotateRight(node)
	}

	// Right Left Case
	if balance < -1 && value < node.Right.Data {
		node.Right = a.rotateRight(node.Right)
		return a.rotateLeft(node)
	}

	return node
}

func (a *AVLTree) Insert(value int) {
	a.root = a.insertHelper(a.root, value)
}

func (a *AVLTree) minValueNode(node *AVLNode) *AVLNode {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func (a *AVLTree) deleteHelper(node *AVLNode, value int) *AVLNode {
	if node == nil {
		return node
	}

	if value < node.Data {
		node.Left = a.deleteHelper(node.Left, value)
	} else if value > node.Data {
		node.Right = a.deleteHelper(node.Right, value)
	} else {
		if node.Left == nil || node.Right == nil {
			var temp *AVLNode
			if node.Left != nil {
				temp = node.Left
			} else {
				temp = node.Right
			}

			if temp == nil {
				return nil
			} else {
				node = temp
			}
		} else {
			temp := a.minValueNode(node.Right)
			node.Data = temp.Data
			node.Right = a.deleteHelper(node.Right, temp.Data)
		}
	}

	if node == nil {
		return node
	}

	node.height = 1 + max(a.height(node.Left), a.height(node.Right))
	balance := a.balanceFactor(node)

	// Left Left Case
	if balance > 1 && a.balanceFactor(node.Left) >= 0 {
		return a.rotateRight(node)
	}

	// Left Right Case
	if balance > 1 && a.balanceFactor(node.Left) < 0 {
		node.Left = a.rotateLeft(node.Left)
		return a.rotateRight(node)
	}

	// Right Right Case
	if balance < -1 && a.balanceFactor(node.Right) <= 0 {
		return a.rotateLeft(node)
	}

	// Right Left Case
	if balance < -1 && a.balanceFactor(node.Right) > 0 {
		node.Right = a.rotateRight(node.Right)
		return a.rotateLeft(node)
	}

	return node
}

func (a *AVLTree) Remove(value int) {
	a.root = a.deleteHelper(a.root, value)
}

func (a *AVLTree) searchHelper(node *AVLNode, value int) *AVLNode {
	if node == nil || node.Data == value {
		return node
	}

	if value < node.Data {
		return a.searchHelper(node.Left, value)
	}
	return a.searchHelper(node.Right, value)
}

func (a *AVLTree) Search(value int) *AVLNode {
	return a.searchHelper(a.root, value)
}

func (a *AVLTree) printInOrderHelper(node *AVLNode) {
	if node != nil {
		a.printInOrderHelper(node.Left)
		fmt.Printf("%d ", node.Data)
		a.printInOrderHelper(node.Right)
	}
}

func (a *AVLTree) PrintInOrder() {
	fmt.Printf("Дерево '%s' in-order: ", a.name)
	a.printInOrderHelper(a.root)
	fmt.Println()
}

func (a *AVLTree) countElementsHelper(node *AVLNode) int {
	if node == nil {
		return 0
	}
	return 1 + a.countElementsHelper(node.Left) + a.countElementsHelper(node.Right)
}

func (a *AVLTree) CountElements() int {
	return a.countElementsHelper(a.root)
}

func (a *AVLTree) saveTreeHelper(node *AVLNode, result *[]int) {
	if node != nil {
		a.saveTreeHelper(node.Left, result)
		*result = append(*result, node.Data)
		a.saveTreeHelper(node.Right, result)
	}
}

func (a *AVLTree) SaveTree() []int {
	result := make([]int, 0)
	a.saveTreeHelper(a.root, &result)
	return result
}

func (a *AVLTree) IsEmpty() bool {
	return a.root == nil
}

func (a *AVLTree) GetName() string {
	return a.name
}

func (a *AVLTree) GetRoot() *AVLNode {
	return a.root
}

func (a *AVLTree) SetRoot(root *AVLNode) {
	a.root = root
}

func (a *AVLTree) Cleanup() {
	a.root = nil
}
