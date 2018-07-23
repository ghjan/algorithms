package main

import (
	"fmt"
	"sort"
)

/*
golang中sort包用法
https://www.cnblogs.com/msnsj/p/4242578.html
 */
//定义interface{},并实现sort.Interface接口的三个方法
type IntSlice []int

func (c IntSlice) Len() int {
	return len(c)
}
func (c IntSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c IntSlice) Less(i, j int) bool {
	return c[i] < c[j]
}

func GuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}

func TestSearch() {
	a := []int{1, 2, 3, 4, 5}
	b := sort.Search(len(a), func(i int) bool { return a[i] >= 30 })
	fmt.Println(b) //5，查找不到，返回a slice的长度５，而不是-1
	c := sort.Search(len(a), func(i int) bool { return a[i] <= 3 })
	fmt.Println(c) //0，利用二分法进行查找，返回符合条件的最左边数值的index，即为０
	d := sort.Search(len(a), func(i int) bool { return a[i] == 3 })
	fmt.Println(d)
}

func main() {
	a := IntSlice{1, 3, 5, 7, 2}
	b := []float64{1.1, 2.3, 5.3, 3.4}
	c := []int{1, 3, 5, 4, 2}
	fmt.Println(sort.IsSorted(a)) //false
	if !sort.IsSorted(a) {
		sort.Sort(a)
	}

	if !sort.Float64sAreSorted(b) {
		sort.Float64s(b)
	}
	if !sort.IntsAreSorted(c) {
		sort.Ints(c)
	}
	fmt.Println(a) //[1 2 3 5 7]
	fmt.Println(b) //[1.1 2.3 3.4 5.3]
	fmt.Println(c) // [1 2 3 4 5]
	GuessingGame()
	TestSearch()
}
