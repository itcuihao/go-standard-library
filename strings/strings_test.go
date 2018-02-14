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
	s := "ab,ab,ab,ab"
	t.Log(Split(s, ","))
	t.Log(Split(s, ""))
	t.Log(Split(s, "."))
}

func TestSplitN(t *testing.T) {
	s := "ab,ab,ab,ab"
	for i := -2; i < 6; i++ {
		t.Logf("n=%d;out:%v", i, SplitN(s, ",", i))
	}
}
func TestSplitAfter(t *testing.T) {
	s := "ab,ab,ab,ab"
	t.Log(SplitAfter(s, ","))
	t.Log(SplitAfter(s, ""))
	t.Log(SplitAfter(s, "."))
}

func TestSplitAfterN(t *testing.T) {
	s := "ab,ab,ab,ab"
	for i := -2; i < 6; i++ {
		t.Logf("n=%d;out:%v", i, SplitAfterN(s, ",", i))
	}
}

func TestTitle(t *testing.T) {
	s := "ab b abb"
	t.Log(Title(s))
}

func TestToTitle(t *testing.T) {
	s := "ab ab ab"
	t.Log(ToTitle(s))
}
func TestToLower(t *testing.T) {
	s := "ab AB ab"
	t.Log(ToLower(s))
}
func TestToUpper(t *testing.T) {
	s := "ab AB ab"
	t.Log(ToUpper(s))
}
func TestTrim(t *testing.T) {
	s := "ab aaABa ba"
	t.Log(Trim(s, "a"))
}
func TestTrimLeft(t *testing.T) {
	s := "ab aaABa ba"
	t.Log(TrimLeft(s, "a"))
}

func TestTrimRight(t *testing.T) {
	s := "ab aaABa ba"
	t.Log(TrimRight(s, "a"))
}
func TestTrimSpace(t *testing.T) {
	s := "\r\nab aaABa ba\n\r"
	t.Log(s)
	t.Log(TrimSpace(s))
}

func TestTrimPrefix(t *testing.T) {
	s := "Abcd"
	t.Log(TrimPrefix(s, "A"))
}
