package tree

import (
	"github.com/kazu-peh3/alGorithm/constraints"
)

type BinarySearchTree[T constraints.Ordered] struct {
	Root *Node[T]
}

type Node[T constraints.Ordered] struct {
	Data   T
	Parent *Node[T]
	Left   *Node[T]
	Right  *Node[T]
}

func NewBinarySearchTree[T constraints.Ordered]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{
		Root: nil,
	}
}

func NewNode[T constraints.Ordered](data T) *Node[T] {
	return &Node[T]{
		Data: data,
	}
}

func (t *BinarySearchTree[T]) Insert(data T) {
	if t.Root == nil {
		// Rootノードがnilの場合、Rootノードを新規作成
		t.Root = NewNode[T](data)
	} else {
		// Rootノードが存在する場合、再帰的にノードを作成
		t.Root.insert(data)
	}
}

func (n *Node[T]) insert(data T) {
	if n.Data > data {
		if n.Left == nil {
			// 挿入する値がノード値より小さいかつ左側のノードが存在しない場合、左側に新規ノード作成
			n.Left = NewNode[T](data)
			n.Left.Parent = n
		} else {
			// 挿入する値がノード値より小さいかつ左側のノードが存在する場合、再帰的に左側へノード作成
			n.Left.insert(data)
		}
	} else {
		if n.Right == nil {
			// 挿入する値がノード値以上かつ右側のノードが存在しない場合、右側に新規ノード作成
			n.Right = NewNode(data)
			n.Right.Parent = n
		} else {
			// 挿入する値がノード値以上かつ右側のノードが存在する場合、再帰的に右側へノード作成
			n.Right.insert(data)
		}
	}
}

func (t *BinarySearchTree[T]) Search(data T) bool {
	return t.Root.search(data)
}

func (n *Node[T]) search(data T) bool {
	if n == nil {
		// ノードがnilのときはfalse
		return false
	}
	if n.Data > data {
		// 検索値がノード値より小さい場合、左側のノードから検索
		return n.Left.search(data)
	} else if n.Data < data {
		// 検索値がノード値より大きい場合、右側のノードから検索
		return n.Right.search(data)
	} else {
		// 検索値とノード値が一致する時はtrue
		return true
	}
}

func (t *BinarySearchTree[T]) Delete(data T) {
	t.Root = t.Root.delete(data)
}

func (n *Node[T]) delete(data T) *Node[T] {
	// ノードがnilの場合はnilを返す
	if n == nil {
		return nil
	}

	if n.Data > data {
		// 削除データがノード値より小さい場合、再帰的に左側のノードから削除
		n.Left = n.Left.delete(data)
	} else if n.Data < data {
		// 削除データがノード値より大きい場合、再帰的に右側のノードから削除
		n.Right = n.Right.delete(data)
	} else {
		// 削除データとノード値を等しい場合、ノードを削除
		// ケース1: 子ノードがいない場合、そのノードを削除
		if n.Left == nil && n.Right == nil {
			return nil
		}
		// ケース2: 右側にのみ子ノードが存在する場合、ノードを削除して右側の子ノードを返す
		if n.Left == nil {
			return n.Right
		}
		// ケース3: 左側にのみ子ノードが存在する場合、ノードを削除して左側の子ノードを返す
		if n.Right == nil {
			return n.Left
		}
		// ケース4: 両側にノードが存在する場合、右側の部分木から最小ノードを取得し、値を置き換えて、その最小ノードを削除
		minNode := n.Right
		for minNode.Left != nil {
			minNode = minNode.Left
		}
		n.Data = minNode.Data
		n.Right = n.Right.delete(minNode.Data)
	}

	return n
}
