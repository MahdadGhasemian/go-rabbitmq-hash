/**
 * Author: Mahdad Ghasemian
 * File: go-rabbitmq-hash_test.go
 */

package rabbitmqHash

import (
	"testing"
)

func TestHash(t *testing.T) {
	p := "test12"
	s := uint32(0x908DC60A)
	expected := "kI3GCqW5JLMJa4iX1lo7X4D6XbYqlLgxIs30+P6tENUV2POR"

	if got := Hash(p, s); got != expected {
		t.Errorf("Hash(%s, %d) = %s, didn't return %s", p, s, got, expected)
	}
}
