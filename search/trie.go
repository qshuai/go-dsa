package search

import (
	"container/list"
	"math"
	"strings"
)

const (
	startPoint       = 'a'
	characterSetSize = 26
)

// TrieNode is a node represent of trie tree. 只支持26个小写的英文字母
type TrieNode struct {
	val      byte                        // math.MaxUint8 represents invalid value, that is the root node
	children [characterSetSize]*TrieNode // 如果元素值为nil代表该元素不存在
	leaf     bool
}

// AddEntry adds an entry to tire tree
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
					children: [characterSetSize]*TrieNode{},
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
		children: [characterSetSize]*TrieNode{},
		leaf:     false,
	}
}

// AcTrieNode AC自动机
type AcTrieNode struct {
	val      byte                          // math.MaxUint8 represents invalid value, that is the root node
	children [characterSetSize]*AcTrieNode // 如果元素值为nil代表该元素不存在
	leaf     bool
	keyword  string // 当leaf为true时，此字段会冗余完整的关键词

	failPtr *AcTrieNode
}

// AddEntry adds an entry to tire tree
func (t *AcTrieNode) AddEntry(keywords ...string) {
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
				travel.children[char] = &AcTrieNode{
					val:      char,
					children: [characterSetSize]*AcTrieNode{},
				}
				if i == len(keyword)-1 {
					travel.children[char].leaf = true
					travel.children[char].keyword = keyword
				}
			}

			if !travel.children[char].leaf && i == len(keyword)-1 {
				travel.children[char].leaf = true
				travel.children[char].keyword = keyword
			}
			travel = travel.children[char]
		}
	}
}

func (t *AcTrieNode) Match(text string) []string {
	t.buildFailurePtr()

	p := t
	var ret []string
	for i := 0; i < len(text); i++ {
		idx := text[i] - startPoint
		for p.children[idx] == nil && p != t {
			p = p.failPtr
		}

		p = p.children[idx]
		if p == nil {
			p = t
		}

		tmp := p
		for tmp != t {
			if tmp.leaf {
				ret = append(ret, tmp.keyword)
			}

			tmp = tmp.failPtr
		}
	}

	return ret
}

// buildFailurePtr 构建AC自动机的失败指针
func (t *AcTrieNode) buildFailurePtr() {
	queue := list.New()
	queue.PushBack(t)
	for queue.Len() > 0 {
		ele := queue.Front()
		queue.Remove(ele)
		p := ele.Value.(*AcTrieNode)

		for i := 0; i < characterSetSize; i++ {
			pc := p.children[i]
			if pc == nil {
				continue
			}
			if p == t {
				pc.failPtr = t
			} else {
				q := p.failPtr
				for q != nil {
					qc := q.children[pc.val]
					if qc != nil {
						pc.failPtr = qc
						break
					}
					q = q.failPtr
				}
				if q == nil {
					pc.failPtr = t
				}
			}

			queue.PushBack(pc)
		}
	}
}

func NewAcTrieTree() *AcTrieNode {
	return &AcTrieNode{
		val:      math.MaxUint8,
		children: [characterSetSize]*AcTrieNode{},
		leaf:     false,
		failPtr:  nil,
	}
}
