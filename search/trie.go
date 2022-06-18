package search

import (
	"math"
	"strings"
)

const startPoint = 'a'

// TrieNode is a node represent of trie tree. 只支持26个小写的英文字母
type TrieNode struct {
	val      byte          // math.MaxUint8 represents invalid value, that is the root node
	children [26]*TrieNode // 如果元素值为nil代表该元素不存在
	leaf     bool
}

// AddEntry adds a entry to tire tree
func (t *TrieNode) AddEntry(keywords ...string) {
	if len(keywords) <= 0 {
		return
	}

	for _, keyword := range keywords {
		keyword = strings.TrimSpace(keyword)
		if len(keyword) <= 0 {
			continue
		}

		travel := t
		for i := 0; i < len(keyword); i++ {
			char := keyword[i] - startPoint
			if travel.children[char] == nil {
				travel.children[char] = &TrieNode{
					val:      char,
					children: [26]*TrieNode{},
					leaf:     i == len(keyword)-1,
				}
			}

			if !travel.children[char].leaf && i == len(keyword)-1 {
				travel.children[char].leaf = true
			}
			travel = travel.children[char]
		}
	}
}

// Search returns all completed entry with the specified prefix;
func (t *TrieNode) Search(prefix string) []string {
	prefix = strings.TrimSpace(prefix)
	if len(prefix) <= 0 {
		return nil
	}

	travel := t
	var self bool
	for i := 0; i < len(prefix); i++ {
		char := prefix[i] - startPoint
		if travel.children[char] == nil {
			return nil
		}

		if i == len(prefix)-1 && travel.children[char].leaf {
			self = true
		}

		travel = travel.children[char]
	}
	if travel == nil {
		if self {
			return []string{prefix}
		}

		return nil
	}

	// 遍历节点获取，所有可能的字符串
	var ret []string
	for i := 0; i < len(travel.children); i++ {
		trieDfs(travel.children[i], &([]byte{}), &ret)
	}

	for idx, entry := range ret {
		ret[idx] = prefix + entry
	}
	if self {
		ret = append(ret, prefix)
	}
	return ret
}

// trieDfs travels all entries after this node
func trieDfs(t *TrieNode, temp *[]byte, collector *[]string) {
	if t == nil {
		return
	}

	*temp = append(*temp, t.val+startPoint)
	if t.leaf {
		*collector = append(*collector, string(*temp))
		*temp = (*temp)[:(len(*temp) - 1)]
	}

	for i := 0; i < len(t.children); i++ {
		trieDfs(t.children[i], temp, collector)
	}
}

func NewTrieTree() *TrieNode {
	return &TrieNode{
		val:      math.MaxUint8,
		children: [26]*TrieNode{},
		leaf:     false,
	}
}
