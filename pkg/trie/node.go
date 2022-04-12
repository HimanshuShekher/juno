package trie

import (
	"fmt"
	"math"
)

// encoding represents the encoding of a node in a binary tree.
type encoding struct {
	length, path, bottom int
}

// node represents a node in a binary tree.
type node struct {
	encoding
	// TODO: Add node hash field.
	// hash *big.Int
	next []node
}

// newNode initialises a node with two null links.
func newNode() node {
	return node{next: make([]node, 2)}
}

// updateHash updates the node hash according to the directive outlined
// in the specification¹.
//
// ¹ https://docs.starknet.io/docs/State/starknet-state#specifications
func (n *node) updateHash() {
	if n.length == 0 {
		// TODO: Set n.hash = n.bottom.
		// DEBUG.
		fmt.Printf("hash = %v\n", n.bottom)
	} else {
		// TODO: Set n.hash = h(n.bottom,n.path) + n.length
		// DEBUG.
		fmt.Printf("hash = h(%v,%v) + %v\n", n.bottom, n.path, n.length)
	}
}

// encode encodes the (parent) node n according to the directive
// outlined in the specification¹.
//
// ¹ https://docs.starknet.io/docs/State/starknet-state#specifications
func (n *node) encode() {
	left, right := n.next[0], n.next[1]

	switch {
	case left.isEmpty() && right.isEmpty():
		n.encoding = encoding{0, 0, 0}
	case !left.isEmpty() && right.isEmpty():
		n.encoding = encoding{left.length + 1, left.path, left.bottom}
	case left.isEmpty() && !right.isEmpty():
		n.encoding = encoding{
			right.length + 1,
			right.path + int(math.Pow(2, float64(right.length))),
			right.bottom,
		}
	default:
		// TODO: Placeholder. Encode the bottom field of this node as
		// h(left.hash, right.hash).
		n.encoding = encoding{0, 0, 251}
	}
}

// isEmpty returns true if next is a null link. This corresponds to an
// empty node in the specification¹.
//
// ¹ https://docs.starknet.io/docs/State/starknet-state#specifications
func (n *node) isEmpty() bool {
	return n.next == nil
}
