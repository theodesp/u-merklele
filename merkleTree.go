package umerklele

import (
	"crypto/sha256"
	"hash"
)

// MerkleTree represents a binary Merkle Tree. This consists of two child nodes, and a
// hash representing those two child nodes. The children can either be leaf
// nodes that contain data blocks or be Merkle Trees.
type MerkleTree struct {
	leftTree, rightTree *MerkleTree

	// The hash function used to calculate the hash values
	hashFunc hash.Hash

	// The hash value of this tree node
	hashCode []byte
}

// New creates a Merkle tree of using the specified hash digest algorithm.
func New(h hash.Hash) *MerkleTree {
	if h == nil {
		h = sha256.New()
	}
	return &MerkleTree{
		nil, nil, h, nil,
	}
}

// NewLeaf creates a Merkle tree of using the specified hash digest algorithm and data
func NewLeaf(data []byte, h hash.Hash) *MerkleTree {
	if h == nil {
		h = sha256.New()
	}

	mt := &MerkleTree{
		nil, nil, h, nil,
	}

	mt.Hash(data)
	return mt
}

// Hash assigns a Payload for a specific Node
func (mt *MerkleTree) Hash(data []byte) *MerkleTree {
	if data == nil {
		return mt
	}

	mt.hashFunc.Write(data)
	mt.hashCode = mt.hashFunc.Sum(nil)
	mt.hashFunc.Reset()

	return mt
}

// LeftTree returns the Left Child Tree if there is one
func (mt MerkleTree) LeftTree() *MerkleTree {
	return mt.leftTree
}

// RightTree returns the Right Child Tree if there is one
func (mt MerkleTree) RightTree() *MerkleTree {
	return mt.rightTree
}

// HashCode Returns the current hash value associated with this root note of this Merkle Tree
func (mt *MerkleTree) HashCode() []byte {
	return mt.hashCode
}

// IsLeaf indicates this Tree is a Leaf (has no children)
func (mt MerkleTree) IsLeaf() bool {
	return len(mt.hashCode) != 0 && mt.leftTree == nil && mt.rightTree == nil
}

// Merge two child subtrees to this Merkle Tree.
func (mt *MerkleTree) Merge(lt, rt *MerkleTree) *MerkleTree {
	if lt == nil || rt == nil {
		return mt
	}

	mt.leftTree = lt
	mt.rightTree = rt
	mt.updateHashCode()

	return mt
}

// Height returns the current tree Height.
func (mt *MerkleTree) Height() uint32 {
	if mt.IsLeaf() {
		return uint32(1)
	}

	var queue = []*MerkleTree{mt}
	var height = uint32(0)

	// Level order traversal
	for len(queue) > 0 {

		var size = len(queue)

		for size > 0 {
			var tempNode = queue[0]
			queue = queue[1:]

			if tempNode.LeftTree() != nil {
				queue = append(queue, tempNode.LeftTree())
			}

			if tempNode.RightTree() != nil {
				queue = append(queue, tempNode.RightTree())
			}

			size--
		}
		height++
	}

	return height
}

// Do calls f for each entry in the Merkele Tree. It does a
// level traversal on each Node.
func (mt *MerkleTree) Do(f func(mt *MerkleTree)) {
	h := mt.Height()

	for i := uint32(0) + 1; i <= h; i++ {
		mt.doAt(f, i)
	}
}

func (mt *MerkleTree) doAt(f func(mt *MerkleTree), level uint32) {
	if mt == nil {
		return
	}

	if level == 1 {
		f(mt)
	} else if level > 1 {
		mt.LeftTree().Do(f)
		mt.RightTree().Do(f)
	}
}

func (mt *MerkleTree) updateHashCode() {
	if !mt.IsLeaf() {
		mt.hashFunc.Write(mt.LeftTree().HashCode())
		mt.hashFunc.Write(mt.RightTree().HashCode())

		mt.hashCode = mt.hashFunc.Sum(nil)
		mt.hashFunc.Reset()
	}
}