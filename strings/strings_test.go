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
func TestLastIndex(t *testing.T) {
	s := "abcabcabc"
	t.Log(LastIndex(s, "abc"))
}
func TestRepeat(t *testing.T) {
	s := "abc"
	t.Log(Repeat(s, 2))
}
func TestReplace(t *testing.T) {
	s := "abcabcabc"
	t.Log(Replace(s, "ab", "c", 0))
}

func TestSplit(t *testing.T) {
	s := "abcabcabc"
	t.Log(Split(s, "b"))
}

func TestSplitN(t *testing.T) {
	s := "ab,ab,ab,ab"
	for i := -2; i < 6; i++ {
		t.Logf("n=%d;out:%v", i, SplitN(s, ",", i))
	}
}
