package block

import (
	"crypto/ecdsa"
	"testing"
)

func newKey() *ecdsa.PrivateKey {
	k, _ := GenerateKey()
	return k
}

func TestBlockValid(t *testing.T) {
	a, b := GenKeys(10)
	x := GenTestTransactions(a, 10)
	p := NewPetition(GenTestSignatures(a, x), x)

	y, err := NewBlock(nil, p, x, b)
	if err != nil {
		t.Error(err)
	}

	if y == nil {
		t.Fail()
	}

}
