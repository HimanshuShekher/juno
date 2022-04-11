package trie

// encoding represents the encoding of a node in a binary tree.
type encoding struct {
	length, path, bottom int
}

// node represents a node in a binary tree.
type node struct {
	encoding
	// hash *big.Int
	next []node
}

// newNode initialises a node with two null links.
func newNode() node {
	return node{next: make([]node, 2)}
}

func (n *node) computeHash() {}

func (n *node) updateEncoding() {}

// isNull returns true if next is a null link.
func (n *node) isNull() bool {
	return n.next == nil
}
