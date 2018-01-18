package strings

import (
	"fmt"
	"strings"
	"unicode"
)

// Compare 比较字符串
func Compare(s string) int {
	fmt.Println("---Compare---")
	fmt.Println(strings.Compare("A", "B"))             // A < B
	fmt.Println(strings.Compare("B", "A"))             // B > A
	fmt.Println(strings.Compare("Japan", "Australia")) // J > A
	fmt.Println(strings.Compare("Australia", "Japan")) // A < J
	fmt.Println(strings.Compare("Germany", "Germany")) // G == G
	fmt.Println(strings.Compare("Germany", "GERMANY")) // GERMANY > Germany
	fmt.Println(strings.Compare("", ""))
	fmt.Println(strings.Compare("", " ")) // Space is less

	return strings.Compare("abc", s)
}

// Contains 包含字符串
func Contains() {
	fmt.Println("---Contains---")
	fmt.Println(strings.Contains("Australia", "Aus")) // Any part of string
	fmt.Println(strings.Contains("Australia", "Australian"))
	fmt.Println(strings.Contains("Japan", "JAPAN")) // Case sensitive
	fmt.Println(strings.Contains("Japan", "JAP"))   // Case sensitive
	fmt.Println(strings.Contains("Become inspired to travel to Australia.", "Australia"))
	fmt.Println(strings.Contains("", ""))
	fmt.Println(strings.Contains("  ", " ")) // space also consider as string
	fmt.Println(strings.Contains("12554", "1"))
}

// ContainsAny 字符串
func ContainsAny() {
	fmt.Println("---ContainsAny---")
	fmt.Println(strings.ContainsAny("Australia", "a"))
	fmt.Println(strings.ContainsAny("Australia", "r & a"))
	fmt.Println(strings.ContainsAny("JAPAN", "j"))
	fmt.Println(strings.ContainsAny("JAPAN", "J"))
	fmt.Println(strings.ContainsAny("JAPAN", "JAPAN"))
	fmt.Println(strings.ContainsAny("JAPAN", "japan"))
	fmt.Println(strings.ContainsAny("Shell-12541", "1"))
	//  Contains vs ContainsAny
	fmt.Println(strings.ContainsAny("Shell-12541", "1-2")) // true
	fmt.Println(strings.Contains("Shell-12541", "1-2"))    // false
}

// Fields 按空格切分
func Fields() {

	fmt.Println("---Fields---")
	testString := "Australia is a country and continent surrounded by the Indian and Pacific oceans."
	testArray := strings.Fields(testString)
	for _, v := range testArray {
		fmt.Println(v)
	}
}

// FieldsFunc 按func要求切分
func FieldsFunc() {
	fmt.Println("---FieldsFunc---")
	x := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	strArray := strings.FieldsFunc(`Australia major cities – Sydney, Brisbane,
		Melbourne, Perth, Adelaide – are coastal`, x)
	for _, v := range strArray {
		fmt.Println(v)
	}

	fmt.Println("\n*****************Split by number*******************\n")

	y := func(c rune) bool {
		return unicode.IsNumber(c)
	}
	testff := strings.FieldsFunc(`1 Sydney Opera House.2 Great Barrier Reef.3 Uluru-Kata Tjuta National Park.4 Sydney Harbour Bridge.5 Blue Mountains National Park.6 Melbourne.7 Bondi Beach`, y)
	for _, w := range testff {
		fmt.Println(w)
	}
}

// HasPrefix 字符串头包含
func HasPrefix() {
	fmt.Println("---HasPrefix---")
	fmt.Println(strings.HasPrefix("Australia", "Aus"))
	fmt.Println(strings.HasPrefix("Australia", "aus"))
	fmt.Println(strings.HasPrefix("Australia", "Jap"))
	fmt.Println(strings.HasPrefix("Australia", ""))
}

// HasSuffix 字符串尾包含
func HasSuffix() {
	fmt.Println("---HasSuffix---")
	fmt.Println(strings.HasSuffix("Australia", "lia"))
	fmt.Println(strings.HasSuffix("Australia", "A"))
	fmt.Println(strings.HasSuffix("Australia", "LIA"))
	fmt.Println(strings.HasSuffix("123456", "456"))
	fmt.Println(strings.HasSuffix("Australia", ""))
}

// Index 索引
func Index() {
	fmt.Println("---Index---")
	fmt.Println(strings.Index("Australia", "Aus"))
	fmt.Println(strings.Index("Australia", "aus"))
	fmt.Println(strings.Index("Australia", "A"))
	fmt.Println(strings.Index("Australia", "a"))
	fmt.Println(strings.Index("Australia", "Jap"))
	fmt.Println(strings.Index("Japan-124", "-"))
	fmt.Println(strings.Index("Japan-124", ""))
}

// IndexAny 索引
func IndexAny() {
	fmt.Println("---IndexAny---")
	fmt.Println(strings.IndexAny("australia", "jupn")) // a position
	fmt.Println(strings.IndexAny("japan", "jpen"))     // p position
	fmt.Println(strings.IndexAny("mobile", "one"))     // o position
	fmt.Println(strings.IndexAny("123456789", "4"))    // 4 position
	fmt.Println(strings.IndexAny("123456789", "0"))    // 0 position
}

// IndexByte 索引
func IndexByte() {
	fmt.Println("---IndexByte---")
	var s, t, u byte
	t = 'l'
	fmt.Println(strings.IndexByte("australia", t))
	fmt.Println(strings.IndexByte("LONDON", t))
	fmt.Println(strings.IndexByte("JAPAN", t))

	s = 1
	fmt.Println(strings.IndexByte("5221-JAPAN", s))

	u = '1'
	fmt.Println(strings.IndexByte("5221-JAPAN", u))
}

// IndexRune 索引
func IndexRune() {
	fmt.Println("---IndexRune---")
	var rs, rt, ru rune
	rt = 'l'
	fmt.Println(strings.IndexRune("australia", rt))
	fmt.Println(strings.IndexRune("LONDON", rt))
	fmt.Println(strings.IndexRune("JAPAN", rt))

	rs = 1
	fmt.Println(strings.IndexRune("5221-JAPAN", rs))

	ru = '1'
	fmt.Println(strings.IndexRune("5221-JAPAN", ru))
}

// Join 连接字符串
func Join(s []string, sep string) string {
	return strings.Join(s, sep)
}

// LastIndex 返回字符串中最后匹配的下标，否则返回-1
func LastIndex(s string, substr string) int {
	return strings.LastIndex(s, substr)
}

// Repeat 字符串重复
func Repeat(s string, count int) string {
	return strings.Repeat(s, count)
}
