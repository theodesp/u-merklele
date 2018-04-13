package main

import (
	"crypto/sha256"
	"fmt"
	"u-merklele"
)

func main() {
	runExample1()
}

func runExample1()  {
	// Create some leaf trees
	block1 := []byte{
		byte(1), byte(10), byte(12), byte(20), byte(90), byte(45), byte(23), byte(67),
	}

	// Create some leaf trees
	block2 := []byte{
		byte(10), byte(45), byte(22), byte(26), byte(78), byte(33), byte(67), byte(22),
	}

	l1 := umerklele.NewLeaf(block1, sha256.New())
	l2 := umerklele.NewLeaf(block2, sha256.New())
	fmt.Printf("Leaf 1 Digest is: %x\n", l1.HashCode())
	fmt.Printf("Leaf 2 Digest is: %x\n", l2.HashCode())

	m1 := umerklele.New(sha256.New())
	m1.Merge(l1, l2)

	fmt.Printf("Merkle Digest of both Leafs is: %x\n", m1.HashCode())

	m2 := umerklele.New(sha256.New())
	m2.Merge(l2, l1)

	fmt.Printf("Merkle Digest of the extended tree is: %x\n", m2.HashCode())

	m2.Do(func(mt *umerklele.MerkleTree) {
		fmt.Printf("[Merkele Tree: hashCode=%x], isLeaf=%t\n", mt.HashCode(), mt.IsLeaf())
	})
}