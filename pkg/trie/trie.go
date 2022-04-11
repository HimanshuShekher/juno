// Package trie implements a binary trie. See the following¹ for 
// details. 
// 
// ¹ https://docs.starknet.io/docs/State/starknet-state/#example-trie
package trie

import "errors"

// Trie represents a binary trie.
type Trie struct {
	root node
}

// ErrNotFound indicates that the queried item does not exist in the
// store.
var ErrNotFound = errors.New("item not found")

// New constructs a new binary trie.
func New() Trie {
	return Trie{}
}

func (t *Trie) get(n node, key, d int) (node, error) {
	if n.isNull() {
		return node{}, ErrNotFound
	}

	if d == isize-1 {
		return n, nil
	}

	b := bit(key, d)
	return t.get(n.next[b], key, d+1)
}

func (t *Trie) put(n node, key, val, d int) node {
	if n.isNull() {
		n = newNode()
	}

	if d == isize-1 {
		n.bottom = val
		// XXX: For now the default values are good enough but with the
		// switch to big.Ints, the encoding has to be explicit here.
		// TODO: Compute hash.
		return n
	}

	b := bit(key, d)
	n.next[b] = t.put(n.next[b], key, val, d+1)

	// TODO: Update the encoding of n.
	// TODO: Compute hash of n.
	return n
}

// Get retrieves a value from the trie with the corresponding key.
func (t *Trie) Get(key int) (int, error) {
	n, err := t.get(t.root, key, 0)
	if err != nil {
		return 0, err
	}
	return n.bottom, nil
}

// Put inserts a key-value pair in the trie.
func (t *Trie) Put(key, val int) {
	t.root = t.put(t.root, key, val, 0)
}
