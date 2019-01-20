package slices

import (
	"fmt"
	"testing"
)

var a = []int{1, 2, 3, 4, 5}

func TestAppendo(t *testing.T) {

	// old a:[1 2 3 4 5], p:0x5f4950
	// 0xc04206a030, &[1 2 3 4 5]
	// 0xc04206a030, &[1 2 3 4 5 6]
	// new a:[1 2 3 4 5 6], p:0x5f4950
	fmt.Printf("old a:%v, p:%p\n", a, &a)
	appendf(&a, 6)
	fmt.Printf("new a:%v, p:%p\n", a, &a)

}
func TestAppends(t *testing.T) {

	// old a:[1 2 3 4 5], p:0x5f3950
	// 0xc042046420, [1 2 3 4 5]
	// 0xc042046420, [1 2 3 4 5 6]
	// new a:[1 2 3 4 5], p:0x5f3950
	// 两个的内存地址不同
	fmt.Printf("old a:%v, p:%p\n", a, &a)
	appendfs(a, 6)
	fmt.Printf("new a:%v, p:%p\n", a, &a)

}
func TestAppenda(t *testing.T) {

	// old a:[1 2 3 4 5], p:0x5f3950
	// 0xc042046420, [1 2 3 4 5]
	// 0xc042046420, [1 2 3 4 5 6]
	// new a:[1 2 3 4 5 6], p:0x5f3950
	fmt.Printf("old a:%v, p:%p\n", a, &a)
	a = appendfa(a, 6)
	fmt.Printf("new a:%v, p:%p\n", a, &a)
}

func TestCut(t *testing.T) {
	a = cut(a, 1, 2)
	fmt.Println(a)
}
func TestDelete(t *testing.T) {
	a = delete(a, 1)
	fmt.Println(a)
}
func TestPops(t *testing.T) {
	a, i := popshift(a)
	fmt.Println(i, a)
}
func TestPope(t *testing.T) {
	a, i := popback(a)
	fmt.Println(i, a)
}
func TestTrick(t *testing.T) {
	b := tricks(a, 1)
	fmt.Printf("b:%v, p:%p\n", b, &b)
}
func TestReverse(t *testing.T) {
	b := reverse(a)
	fmt.Printf("b:%v, p:%p\n", b, &b)
}
