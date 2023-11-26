package tree

type BinaryTree struct {
	Root *Node
}
type Node struct {
	Value int
	Left  *Node
	Right *Node
}
