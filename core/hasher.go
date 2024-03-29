package core

import (
	"blockchain/types"
	"crypto/sha256"
)

type Hasher[T any] interface {
	Hash(T) types.Hash
}

type BlockHasher struct{}

func (BlockHasher) Hash(he *Header) types.Hash {
	h := sha256.Sum256(he.Bytes())
	return types.Hash(h)
}
