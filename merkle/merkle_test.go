package merkle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerkleRootInit(t *testing.T) {
	// Leaf Nodes
	n1 := NewLeafNode("n1")
	n2 := NewLeafNode("n2")
	n3 := NewLeafNode("n3")
	n4 := NewLeafNode("n4")
	assert.Equal(t, n1.data, "676b8bb84ce7267dd520deca4811c8f10a53e636352f06987f42fe425acedd80")
	// Level 2
	n12 := NewBranchNode(&n1, &n2)
	n34 := NewBranchNode(&n3, &n4)

	root := NewBranchNode(&n12, &n34)
	assert.Equal(t, root.data, "57479c722644e4f40ae74170319142d8c95b480434a23d1f96de9dee0a846ded")
	mkTree := MerkleTree{root: &root}
	path := []*Node{&n2, &n34}
	assert.Equal(t, mkTree.Verify("n1", path), root.data)
}
