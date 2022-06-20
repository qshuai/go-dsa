package search

import (
	"testing"

	"github.com/qshuai/go-dsa/utils"
)

func TestTrieNode_Search(t *testing.T) {
	root := NewTrieTree()
	root.AddEntry("abcd", "bc", "bcd", "c", "yy")
	ret1 := root.Search("a")
	want := []string{"abcd"}
	if !utils.ElementEqual[string](ret1, want) {
		t.Errorf("Search() = %v, want: %v", ret1, want)
	}

	ret2 := root.Search("bc")
	want = []string{"bc", "bcd"}
	if !utils.ElementEqual(ret2, want) {
		t.Errorf("Search() = %v, want %v", ret2, want)
	}

	ret3 := root.Search("c")
	want = []string{"c"}
	if !utils.ElementEqual(ret3, want) {
		t.Errorf("Search() = %v, want %v", ret3, want)
	}

	ret4 := root.Search("xx")
	want = []string{}
	if !utils.ElementEqual(ret4, want) {
		t.Errorf("Search() = %v, want %v", ret4, want)
	}
}

func TestAcTrieNode_Match(t *testing.T) {
	root := NewAcTrieTree()
	root.AddEntry("abcd", "bc", "bcd", "c", "yy")
	keywords := root.Match("xxabcdef")
	want := []string{"abcd", "bc", "bcd", "c"}
	if !utils.ElementEqual[string](keywords, want) {
		t.Errorf("Match() = %v, want %v", keywords, want)
	}

	root = NewAcTrieTree()
	root.AddEntry("hers", "he", "his", "she")
	keywords = root.Match("ahishers")
	want = []string{"his", "she", "hers", "he"}
	if !utils.ElementEqual[string](keywords, want) {
		t.Errorf("Match() = %v, want %v", keywords, want)
	}
}
