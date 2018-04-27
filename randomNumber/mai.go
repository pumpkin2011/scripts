package main

import "fmt"

// nx = (a * (n-1)x + b) % m
// a b m 决定了随机数的质量
// 高性能规则：
// a = 4p + 1
// b = 2q + 1
// p q 为正整数
// m 影响序列周期长短，越大越好
// a b值越大越均匀
// a m互质，质量比较好
const (
	a = 111
	b = 313
	m = 10
)

func randomNumber(prev int) int {
	return (a*prev + b) % m
}

func main() {
	mm := make(map[int]int)
	num := 10
	for j := 0; j < 100; j++ {
		mm = make(map[int]int)
		for i := 0; i < 999; i++ {
			num = randomNumber(num)
			mm[num]++
		}
		fmt.Println(mm)
	}
}
