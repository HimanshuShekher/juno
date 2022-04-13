package trie

import (
	"fmt"
	"math"
	"math/big"

	"github.com/NethermindEth/juno/pkg/crypto/pedersen"
)

// encoding represents the encoding of a node in a binary tree.
type encoding struct {
	length       uint8
	path, bottom *big.Int
}

// node represents a node in a binary tree.
type node struct {
	encoding
	hash *big.Int
	next []node
}

// newNode initialises a node with two null links.
func newNode() node {
	return node{next: make([]node, 2)}
}

// DEBUG.
// String makes encoding satisfy the Stringer interface.
func (e *encoding) String() string {
	return fmt.Sprintf("(%d, %d, %x)", e.length, e.path, e.bottom)
}

// updateHash updates the node hash according to the directive outlined
// in the specification¹.
//
// ¹ https://docs.starknet.io/docs/State/starknet-state#specifications
func (n *node) updateHash() {
	if n.length == 0 {
		n.hash = new(big.Int).Set(n.bottom)

		// DEBUG.
		fmt.Printf("hash = %x\n\n", n.hash)
	} else {
		h, _ := pedersen.Digest(n.bottom, n.path)
		n.hash = h.Add(h, new(big.Int).SetUint64(uint64(n.length)))

		// DEBUG.
		fmt.Printf("hash = %x\n\n", n.hash)
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
		n.encoding = encoding{0, new(big.Int), new(big.Int)}
	case !left.isEmpty() && right.isEmpty():
		n.encoding = encoding{
			left.length + 1, left.path, new(big.Int).Set(left.bottom),
		}
	case left.isEmpty() && !right.isEmpty():
		addend := new(big.Int).SetUint64(uint64(math.Pow(2, float64(right.length))))
		n.encoding = encoding{
			right.length + 1,
			right.path.Add(right.path, addend),
			new(big.Int).Set(right.bottom),
		}
	default:
		h, _ := pedersen.Digest(left.hash, right.hash)
		n.encoding = encoding{0, new(big.Int), h}
	}
}

// isEmpty returns true if n is an empty node i.e. (0,0,0)¹.
//
// ¹ https://docs.starknet.io/docs/State/starknet-state#specifications
func (n *node) isEmpty() bool {
	return n.next == nil
}
