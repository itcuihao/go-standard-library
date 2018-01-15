package strings

import (
	"testing"
)

// go test -v -test.run TestCompare
func TestCompare(t *testing.T) {
	i := Compare("a")
	t.Log(i)
}

// go test -v -test.run TestJoin
func TestJoin(t *testing.T) {
	s := []string{"I", "am", "hao"}
	t.Log(Join(s, " "))
}
