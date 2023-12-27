package tree

import (
	"testing"
)

func TestInsert(t *testing.T) {
	bst := NewBinarySearchTree[int]()
	bst.Insert(90)
	bst.Insert(80)
	bst.Insert(100)

	if bst.Root.Data != 90 {
		t.Errorf("Root node should have value = 90")
	}
	if bst.Root.Left.Data != 80 {
		t.Errorf("Left node should have value = 80")
	}
	if bst.Root.Right.Data != 100 {
		t.Errorf("Right node should have value = 100")
	}
}

func TestDelete(t *testing.T) {
	t.Run("Delete a node with no child", func(t *testing.T) {
		bst := NewBinarySearchTree[int]()

		// Insert
		bst.Insert(100)
		bst.Insert(80)
		bst.Insert(120)

		// Delete
		bst.Delete(120)

		if bst.Search(120) {
			t.Errorf("Failed to delete expected data")
		}
		if !bst.Search(100) {
			t.Errorf("Failed to delete expected data")
		}
		if !bst.Search(80) {
			t.Errorf("Failed to delete expected data")
		}
	})
	t.Run("Delete a node with single child", func(t *testing.T) {
		bst := NewBinarySearchTree[int]()

		// Insert
		bst.Insert(100)
		bst.Insert(80)
		bst.Insert(120)
		bst.Insert(140)

		// Delete
		bst.Delete(120)

		if bst.Search(120) {
			t.Errorf("Failed to delete expected data")
		}
		if !bst.Search(100) {
			t.Errorf("Failed to delete expected data")
		}
		if !bst.Search(80) {
			t.Errorf("Failed to delete expected data")
		}
		if !bst.Search(140) {
			t.Errorf("Failed to delete expected data")
		}
	})
	t.Run("Delete a node with both children", func(t *testing.T) {
		bst := NewBinarySearchTree[int]()

		// Insert
		bst.Insert(100)
		bst.Insert(80)
		bst.Insert(120)
		bst.Insert(110)
		bst.Insert(140)
		bst.Insert(105)
		bst.Insert(115)

		// Delete
		bst.Delete(100)

		if bst.Search(100) {
			t.Errorf("Failed to delete expected data")
		}
		if !bst.Search(80) {
			t.Errorf("Failed to delete expected data")
		}
		if !bst.Search(120) {
			t.Errorf("Failed to delete expected data")
		}
		if !bst.Search(110) {
			t.Errorf("Failed to delete expected data")
		}
		if !bst.Search(140) {
			t.Errorf("Failed to delete expected data")
		}
		if !bst.Search(105) {
			t.Errorf("Failed to delete expected data")
		}
		if !bst.Search(115) {
			t.Errorf("Failed to delete expected data")
		}
	})
}
