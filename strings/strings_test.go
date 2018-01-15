package strings

import (
	"testing"
)

// go test -v -test.run TestCompare
func TestCompare(t *testing.T) {
	i := Compare("a")
	t.Log(i)
}
