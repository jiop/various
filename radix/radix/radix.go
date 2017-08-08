package radix

// original : https://github.com/armon/go-radix

import (
	"sort"
	"strings"
)

type edge struct {
	label byte
	node  *node
}

type edges []edge

func (e edges) Len() int {
	return len(e)
}

func (e edges) Less(i, j int) bool {
	return e[i].label < e[j].label
}

func (e edges) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func (e edges) Sort() {
	sort.Sort(e)
}

type leafNode struct {
	key string
	val interface{}
}

type node struct {
	leaf   *leafNode
	prefix string
	edges  edges
}

func (n *node) isLeaf() bool {
	return n.leaf != nil
}

func (n *node) getEdge(label byte) *node {
	num := len(n.edges)
	idx := sort.Search(num, func(i int) bool {
		return n.edges[i].label >= label
	})
	if idx < num && n.edges[idx].label == label {
		return n.edges[idx].node
	}
	return nil
}

func (n *node) addEdge(e edge) {
	n.edges = append(n.edges, e)
	n.edges.Sort()
}

func (n *node) replaceEdge(e edge) {
	num := len(n.edges)
	idx := sort.Search(num, func(i int) bool {
		return n.edges[i].label >= e.label
	})
	if idx < num && n.edges[idx].label == e.label {
		n.edges[idx].node = e.node
		return
	}
	panic("replacing missing edge")
}

type Tree struct {
	root *node
	size int
}

func New() *Tree {
	return &Tree{root: &node{}}
}

func NewFromMap(init map[string]interface{}) *Tree {
	tree := &Tree{root: &node{}}
	for k, v := range init {
		tree.Insert(k, v)
	}
	return tree
}

func (t *Tree) Len() int {
	return t.size
}

func longestPrefix(a, b string) int {
	max := len(a)
	if l := len(b); l < max {
		max = l
	}
	for i := 0; i < max; i++ {
		if a[i] != b[i] {
			return i
		}
	}
	return max
}

func (t *Tree) Insert(s string, v interface{}) (interface{}, bool) {
	var parent *node
	search := s
	n := t.root
	for {
		if len(search) == 0 {
			if n.isLeaf() {
				var old interface{}
				old, n.leaf.val = n.leaf.val, v
				return old, true
			}
			n.leaf = &leafNode{key: s, val: v}
			t.size++
			return nil, false
		}

		parent = n
		n = n.getEdge(search[0])

		if n == nil {
			parent.addEdge(edge{
				label: search[0],
				node: &node{
					leaf:   &leafNode{key: s, val: v},
					prefix: search,
				},
			})
			t.size++
			return nil, false
		}

		commonPrefix := longestPrefix(search, n.prefix)
		if commonPrefix == len(n.prefix) {
			search = search[commonPrefix:]
			continue
		}

		t.size++
		child := &node{prefix: search[:commonPrefix]}
		parent.replaceEdge(edge{label: search[0], node: child})
		child.addEdge(edge{label: n.prefix[commonPrefix], node: n})
		n.prefix = n.prefix[commonPrefix:]

		leaf := &leafNode{key: s, val: v}

		search = search[commonPrefix:]
		if len(search) == 0 {
			child.leaf = leaf
			return nil, false
		}

		child.addEdge(edge{label: search[0], node: &node{leaf: leaf, prefix: search}})
		return nil, false
	}
}

func (t *Tree) Get(s string) (interface{}, bool) {
	n := t.root
	search := s
	for {
		if len(search) == 0 {
			if n.isLeaf() {
				return n.leaf.val, true
			}
			break
		}

		if n = n.getEdge(search[0]); n == nil {
			break
		}

		if !strings.HasPrefix(search, n.prefix) {
			break
		}

		search = search[len(n.prefix):]
	}
	return nil, false
}

func (t *Tree) Minimum() (string, interface{}, bool) {
	n := t.root
	for {
		if n.isLeaf() {
			return n.leaf.key, n.leaf.val, true
		}
		if len(n.edges) <= 0 {
			break
		}
		n = n.edges[0].node
	}
	return "", nil, false
}

func (t *Tree) Maximum() (string, interface{}, bool) {
	n := t.root
	for {
		if num := len(n.edges); num > 0 {
			n = n.edges[num-1].node
			continue
		}
		if n.isLeaf() {
			return n.leaf.key, n.leaf.val, true
		}
	}
}

type walkFn func(string, interface{}) bool

func (t *Tree) Walk(fn walkFn) {
	recursiveWalk(t.root, fn)
}

func recursiveWalk(n *node, fn walkFn) bool {
	if n.leaf != nil && fn(n.leaf.key, n.leaf.val) {
		return true
	}
	for _, e := range n.edges {
		if recursiveWalk(e.node, fn) {
			return true
		}
	}
	return false
}

func (t *Tree) ToMap() map[string]interface{} {
	out := make(map[string]interface{}, t.size)
	t.Walk(func(k string, v interface{}) bool {
		out[k] = v
		return false
	})
	return out
}

func (n *node) delEdge(label byte) {
	num := len(n.edges)
	idx := sort.Search(num, func(i int) bool {
		return n.edges[i].label >= label
	})
	if idx < num && n.edges[idx].label == label {
		copy(n.edges[idx:], n.edges[idx+1:])
		n.edges[len(n.edges)-1] = edge{}
		n.edges = n.edges[:len(n.edges)-1]
	}
}

func (n *node) mergeChild() {
	e := n.edges[0]
	child := e.node
	n.prefix = n.prefix + child.prefix
	n.leaf = child.leaf
	n.edges = child.edges
}

func (t *Tree) Delete(s string) (interface{}, bool) {
	var parent *node
	var label byte

	n := t.root
	search := s

	for {
		if len(search) == 0 {
			if !n.isLeaf() {
				break
			}

			var leaf *leafNode
			leaf, n.leaf = n.leaf, nil
			t.size--

			if parent != nil && len(n.edges) == 0 {
				parent.delEdge(label)
			}

			if n != t.root && len(n.edges) == 1 {
				n.mergeChild()
			}

			if parent != nil && parent != t.root && len(parent.edges) == 1 && !parent.isLeaf() {
				parent.mergeChild()
			}

			return leaf.val, true
		}

		parent = n
		label = search[0]
		n = n.getEdge(label)
		if n == nil {
			break
		}

		if !strings.HasPrefix(search, n.prefix) {
			break
		}
		search = search[len(n.prefix):]
	}

	return nil, false
}
