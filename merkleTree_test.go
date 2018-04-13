package umerklele

import (
	"crypto/sha1"
	"crypto/sha256"
	"hash"
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type MerkleTestSuite struct{}

var _ = Suite(&MerkleTestSuite{})

func (s *MerkleTestSuite) TestNew(c *C) {
	testCases := []struct {
		inputHash    hash.Hash
		expectedHash hash.Hash
	}{
		{
			nil,
			sha256.New(),
		},
		{
			sha1.New(),
			sha1.New(),
		},
	}

	for _, tc := range testCases {
		c.Assert(New(tc.inputHash).hashFunc, DeepEquals, tc.expectedHash)
	}
}
