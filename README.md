u-merklele
---
<a href="https://godoc.org/github.com/theodesp/u-merklele">
<img src="https://godoc.org/github.com/theodesp/u-merklele?status.svg" alt="GoDoc">
</a>

<a href="https://opensource.org/licenses/MIT" rel="nofollow">
<img src="https://img.shields.io/github/license/mashape/apistatus.svg" alt="License"/>
</a>

<a href="https://travis-ci.org/theodesp/u-merklele" rel="nofollow">
<img src="https://travis-ci.org/theodesp/u-merklele.svg?branch=master" />
</a>

<a href="https://codecov.io/gh/theodesp/u-merklele">
  <img src="https://codecov.io/gh/theodesp/u-merklele/branch/master/graph/badge.svg" />
</a>

Simple reference implementation of Merkle trees for general use.

A **hash tree** or **Merkle tree** is a tree in which every leaf node is labelled with the hash of a data 
block and every non-leaf node is labelled with the cryptographic hash of the labels of its child nodes. 
Hash trees allow efficient and secure verification of the contents of large data structures.

## Installation
```bash
$ go get -u github.com/theodesp/u-merklele
```

## Usage
```go
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
```

You can also run the demo application in the `example` folder

## LICENCE
Copyright © 2017 Theo Despoudis MIT license