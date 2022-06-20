package search

import (
	"testing"

	"github.com/qshuai/go-dsa/utils"
)

func TestAcTrieNode_Match(t *testing.T) {
	root := NewAcTrieTree()
	root.AddEntry("abcd", "bc", "bcd", "c", "yy")
	keywords := root.Match("xxabcdef")
	want := []string{"abcd", "bc", "bcd", "c"}
	if !utils.ElementEqual[string](keywords, want) {
		t.Errorf("Match() = %v want %v", keywords, want)
	}
}
