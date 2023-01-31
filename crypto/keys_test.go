package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	privKey := GeneratePrivateKey()
	assert.Equal(t, len(privKey.Bytes()), privKetLen)
	pubKey := privKey.Public()
	assert.Equal(t, len(pubKey.Bytes()), pubKeyLen)
}

func TestNewPrivateKeyFromString(t *testing.T) {
	var (
		seed       = "63088a6abef92633fc563344709c11e082196f67daa8efdfee3cb6de5b41a355"
		privKey    = NewPrivateKeyFromString(seed)
		addressStr = "15b1b2e79f06823edc66f7db4e5818e5688b1cc5"
	)

	assert.Equal(t, privKetLen, len(privKey.Bytes()))
	address := privKey.Public().Address()
	assert.Equal(t, addressStr, address.String())
}

func TestPrivateKeySign(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()
	msg := []byte("foo bar baz")

	sig := privKey.Sign(msg)
	assert.True(t, sig.Verify(pubKey, msg))

	// Test with invalide msg
	assert.False(t, sig.Verify(pubKey, []byte("foo")))

	// Test with invalide pubkey
	invalidePrivKey := GeneratePrivateKey()
	invalidePubKey := invalidePrivKey.Public()
	assert.False(t, sig.Verify(invalidePubKey, msg))
}

func TestPublicKeyToAddress(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.Public()
	address := pubKey.Address()
	assert.Equal(t, addressLen, len(address.Bytes()))
	fmt.Println(address)
}
