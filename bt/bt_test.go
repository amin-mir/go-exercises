package bt

import (
	"reflect"
	"testing"
)

func TestNode_Insert(t *testing.T) {
	tests := []struct {
		name     string
		root     int
		vals     []int
		expected []int
	}{
		{
			name:     "should work for {10,1,40,7,50,5}",
			root:     10,
			vals:     []int{1, 40, 7, 50, 5},
			expected: []int{1, 5, 7, 10, 40, 50},
		},
		{
			name:     "should work for {10,5,40,50,1,7}",
			root:     10,
			vals:     []int{5, 40, 50, 1, 7},
			expected: []int{1, 5, 7, 10, 40, 50},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := New(tt.root)
			for _, v := range tt.vals {
				root.Insert(v)
			}

			sorted := root.Sorted()
			if !reflect.DeepEqual(tt.expected, sorted) {
				t.Fatalf("expected: %v, got %v", tt.expected, sorted)
			}
		})
	}
}

func TestNode_Search(t *testing.T) {
	root := New(10)
	vals := []int{5, 40, 50, 1, 7}
	for _, v := range vals {
		root.Insert(v)
	}

	if !root.Search(7) {
		t.Fatalf("Search couldn't find %d", 7)
	}

	if root.Search(11) {
		t.Fatalf("Seach shouldn't find %d", 11)
	}

	if !root.Search(50) {
		t.Fatalf("Search couldn't find %d", 50)
	}
}

func TestNode_Delete(t *testing.T) {
	t.Run("should delete leaf nodes", func(t *testing.T) {
		root := New(50)
		vals := []int{30, 70, 20, 40, 60, 80}
		for _, v := range vals {
			root.Insert(v)
		}

		root.Delete(20)

		sorted := root.Sorted()
		expected := []int{30, 40, 50, 60, 70, 80}
		if !reflect.DeepEqual(expected, sorted) {
			t.Fatalf("expected: %v, got %v", expected, sorted)
		}
	})

	t.Run("should delete node with one child", func(t *testing.T) {
		root := New(50)
		vals := []int{30, 70, 40, 60, 80}
		for _, v := range vals {
			root.Insert(v)
		}

		root.Delete(30)

		sorted := root.Sorted()
		expected := []int{40, 50, 60, 70, 80}
		if !reflect.DeepEqual(expected, sorted) {
			t.Fatalf("expected: %v, got %v", expected, sorted)
		}
	})

	t.Run("should delete node with two children", func(t *testing.T) {
		root := New(50)
		vals := []int{40, 70, 60, 80}
		for _, v := range vals {
			root.Insert(v)
		}

		root.Delete(50)

		sorted := root.Sorted()
		expected := []int{40, 60, 70, 80}
		if !reflect.DeepEqual(expected, sorted) {
			t.Fatalf("expected: %v, got %v", expected, sorted)
		}
	})
}
