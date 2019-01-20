package slices

import (
	"fmt"
)

// 指针传值才会改变原来的
func appendf(a *[]int, i ...int) {
	fmt.Printf("%p, %v\n", &a, a)
	*a = append(*a, i...)
	fmt.Printf("%p, %v\n", &a, a)
}

// 不返回
func appendfs(a []int, i ...int) {
	fmt.Printf("%p, %v\n", &a, a)
	a = append(a, i...)
	fmt.Printf("%p, %v\n", &a, a)

}

// 返回
func appendfa(a []int, i ...int) []int {
	fmt.Printf("%p, %v\n", &a, a)
	a = append(a, i...)
	fmt.Printf("%p, %v\n", &a, a)
	return a
}

func cut(a []int, i, j int) []int {
	a = append(a[:i], a[j:]...)
	return a
}

func delete(a []int, i int) []int {
	a = append(a[:i], a[i+1:]...)
	return a
}

func popshift(a []int) ([]int, int) {
	i, s := a[0], a[1:]
	return s, i
}
func popback(a []int) ([]int, int) {
	i, s := a[len(a)-1], a[:len(a)-1]
	return s, i
}

func tricks(a []int, i int) []int {
	fmt.Printf("a:%v, p:%p\n", a, &a)
	fmt.Printf("len(a):%v, cap(a):%v\n", len(a), cap(a))
	b := a[:0]
	fmt.Printf("b:%v, p:%p\n", b, &b)
	fmt.Printf("len(b):%v, cap(b):%v\n", len(b), cap(b))
	for _, data := range a {
		if data == i {
			continue
		}
		b = append(b, data)
	}
	fmt.Printf("a:%v, p:%p\n", a, &a)
	fmt.Printf("len(a):%v, cap(a):%v\n", len(a), cap(a))
	fmt.Printf("b:%v, p:%p\n", b, &b)
	fmt.Printf("len(b):%v, cap(b):%v\n", len(b), cap(b))
	return b
}

func reverse(a []int) []int {
	for l, r := 0, len(a)-1; l < r; l, r = l+1, r-1 {
		a[l], a[r] = a[r], a[l]
	}
	return a
}
