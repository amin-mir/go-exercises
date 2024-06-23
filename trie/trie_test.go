package trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	tests := []struct {
		name string
		trie Trie
	}{
		{
			name: "TrieSlice",
			trie: NewTrieSlice(),
		},
		{
			name: "TrieMap",
			trie: NewTrieMap(),
		},
		{
			name: "TrieMapNoAlloc",
			trie: NewTrieMapNoAlloc(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.trie.Insert("the")
			tt.trie.Insert("these")
			tt.trie.Insert("their")
			tt.trie.Insert("any")
			tt.trie.Insert("answer")

			if !tt.trie.Search("the") {
				t.Fatalf("'the' should be found in Trie")
			}
			if tt.trie.Search("them") {
				t.Fatalf("'them' should not be found in Trie")
			}
		})
	}
}

func BenchmarkNodeSlice(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		trie := NewTrieSlice()
		trie.Insert("the")
		trie.Insert("them")
		trie.Insert("any")
		trie.Insert("answer")
		trie.Insert("their")
		trie.Insert("falcon")
		trie.Insert("fall")

		trie.Search("fal")
		trie.Search("fall")
		trie.Search("them")
		trie.Search("theirs")
	}
}

func BenchmarkNodeMap(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		trie := NewTrieMap()
		trie.Insert("the")
		trie.Insert("them")
		trie.Insert("any")
		trie.Insert("answer")
		trie.Insert("their")
		trie.Insert("falcon")
		trie.Insert("fall")

		trie.Search("fal")
		trie.Search("fall")
		trie.Search("them")
		trie.Search("theirs")
	}
}

func BenchmarkNodeMapNoAlloc(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		trie := NewTrieMapNoAlloc()
		trie.Insert("the")
		trie.Insert("them")
		trie.Insert("any")
		trie.Insert("answer")
		trie.Insert("their")
		trie.Insert("falcon")
		trie.Insert("fall")

		trie.Search("fal")
		trie.Search("fall")
		trie.Search("them")
		trie.Search("theirs")
	}
}
