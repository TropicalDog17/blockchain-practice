package merkle

import (
	"crypto/sha256"
	"fmt"
)

type MerkleTree struct {
	root *Node
}
type Node struct {
	left  *Node
	right *Node
	data  string
}

func NewNode(data string) Node {
	node := Node{
		left:  nil,
		right: nil,
		data:  data,
	}
	return node
}
func NewLeafNode(txData string) Node {
	h := sha256.New()
	h.Write([]byte(txData))
	hashedData := fmt.Sprintf("%x", h.Sum(nil))

	return NewNode(hashedData)
}
func NewBranchNode(left *Node, right *Node) Node {
	h := sha256.New()
	h.Write([]byte(left.data + right.data))
	hashedData := fmt.Sprintf("%x", h.Sum(nil))

	return NewNode(hashedData)
}

func (mkTree *MerkleTree) Verify(txData string, path []*Node) string {
	h := sha256.New()
	h.Write([]byte(txData))
	currHash := fmt.Sprintf("%x", h.Sum(nil))
	for _, node := range path {
		// Compute hash of two child nodes
		currHash += node.data
		h.Reset()
		h.Write([]byte(currHash))
		// Get the hash of new parent node
		currHash = fmt.Sprintf("%x", h.Sum(nil))
	}
	hashedData := fmt.Sprintf("%x", h.Sum(nil))
	return hashedData
}

// func NewMerkleTree() MerkleTree {
// 	root := NewNode("")

// 	return MerkleTree{
// 		root: &root,
// 	}
// }
