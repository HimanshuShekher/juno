package trie

// isize gives the bit size of an int (which is either 32 or 64
// depending on the system).
const isize = 32 << (^uint(0) >> 63)

// bit extracts the i-th bit of a (zero-indexed).
func bit(a, i int) int {
	return (a >> (isize - i - 2)) & 1
}
