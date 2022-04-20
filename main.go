package main

import "fmt"

var myData struct {
	Name string
	Data []int
}

type MyData struct {
	Name string
	Data []int
}

func main() {
	// n := 123
	// p := &n
	// q := &p
	// m := 10000
	// p2 := &m
	// q2 := &p2
	// fmt.Printf("p value:%d, address%p\n", *p, p)
	// fmt.Printf("q value:%d, address%p\n", **q, *q)
	// fmt.Printf("p2 value:%d, address%p\n", *p2, p2)
	// fmt.Printf("q2 value:%d, address%p\n", **q2, *q2)
	// pb := p
	// p = p2
	// p2 = pb
	// fmt.Printf("p value:%d, address%p\n", *p, p)
	// fmt.Printf("p2 value:%d, address%p\n", *p2, p2)

	// run()
	// sl()

	// myData.Name = "Taro"
	// myData.Data = []int{10, 20, 30}
	// fmt.Println(myData)

	// taro := MyData{"taro", []int{10, 20, 30}}
	// hanako := MyData{
	// 	Name: "hanako",
	// 	Data: []int{90, 80, 60},
	// }
	// fmt.Println(taro)
	// rev(&taro)
	// fmt.Println(taro)
	// fmt.Println(hanako)

	taro := new(MyData)
	fmt.Println(taro)
	taro.Name = "Taro"
	taro.Data = make([]int, 5)
	fmt.Println(taro)
}

func run() {
	n := 123
	fmt.Printf("value:%d.\n", n)
	change1(n)
	fmt.Printf("value:%d.\n", n)
	change2(&n)
	fmt.Printf("value:%d.\n", n)
}

func change1(n int) {
	n *= 2
}

func change2(n *int) {
	*n *= 2
}

func sl() {
	ar := []int{10, 20, 30}
	fmt.Println(ar)
	initial(&ar)
	fmt.Println(ar)
}

func initial(ar *[]int) {
	for i := 0; i < len(*ar); i++ {
		(*ar)[i] = 0
	}
}

func rev(md *MyData) {
	od := (*md).Data
	nd := []int{}
	for i := len(od) - 1; i >= 0; i-- {
		nd = append(nd, od[i])
	}
	md.Data = nd
}
