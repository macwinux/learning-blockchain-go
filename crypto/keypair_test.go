package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeypair_Sign_Verify_Valid(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()
	msg := []byte("Hello World")

	sig, err := privKey.Sign(msg)

	assert.Nil(t, err)
	assert.True(t, sig.Verify(pubKey, msg))
}

func TestKeypair_Sign_Verify_Invalid(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()
	msg := []byte("Hello World")

	sig, err := privKey.Sign(msg)
	oPrivKey := GeneratePrivateKey()
	oPubKey := oPrivKey.PublicKey()
	assert.Nil(t, err)
	assert.False(t, sig.Verify(oPubKey, msg))
	assert.False(t, sig.Verify(pubKey, []byte("xsdfadsfe")))
}
