package xp

import (
	"crypto/rand"

	"golang.org/x/crypto/curve25519"
)

func P() (private, public []byte, err error) {
	private = make([]byte, curve25519.ScalarSize)
	_, err = rand.Read(private)
	if err != nil {
		return
	}
	public, err = X(private, nil)
	return
}

func X(scalar, point []byte) (product []byte, err error) {
	if point == nil {
		point = curve25519.Basepoint
	}
	return curve25519.X25519(scalar, point)
}
