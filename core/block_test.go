package core

import (
	"blockchain/crypto"
	"blockchain/types"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func RandomBlock(height uint32) *Block {
	header := &Header{
		Version:       1,
		PrevBlockHash: types.RandomHash(),
		Height:        height,
		Timestamp:     uint64(time.Now().UnixNano()),
	}
	tx := Transaction{
		Data: []byte("foo"),
	}

	return NewBlock(header, []Transaction{tx})
}

func TestSignBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := RandomBlock(0)
	assert.Nil(t, b.Sign(privKey))
	assert.Nil(t, b.Verify())
	oPrivKey := crypto.GeneratePrivateKey()
	b.Validator = oPrivKey.PublicKey()
	assert.NotNil(t, b.Verify())

}
