/**
 * Author: Mahdad Ghasemian
 * File: go-rabbitmq-hash.go
 */

package lib

import (
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
)

// Return a new sha256 password according to rabbitmq requirements
func Hash(password string, salt uint32) string {
	// https://www.rabbitmq.com/passwords.html
	var saNum uint32
	if salt == 0 {
		saNum = rand.Uint32()
	} else {
		saNum = salt
	}

	sa := i32tob(saNum)
	s1 := append(sa[:], []byte(password[:])...)
	sum := sha256.Sum256(s1)
	s2 := append(sa[:], []byte(sum[:])...)
	return base64.StdEncoding.EncodeToString(s2[:])
}

func i32tob(val uint32) []byte {
	r := make([]byte, 4)
	for i := uint32(0); i < 4; i++ {
		r[3-i] = byte((val >> (8 * i)) & 0xff)
	}
	return r
}
