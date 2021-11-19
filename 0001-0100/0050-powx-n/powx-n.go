package main

import "fmt"

/*

50. Pow(x, n)

https://leetcode.com/problems/powx-n/

*/

func myPowHelper(x float64, n int, memory map[int]float64) float64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	if value, ok := memory[n]; ok {
		return value
	}

	result := myPowHelper(x, n/2, memory) * myPowHelper(x, n-n/2, memory)

	memory[n] = result

	return result
}

func myPow(x float64, n int) float64 {
	memory := make(map[int]float64)

	if n >= 0 {
		return myPowHelper(x, n, memory)
	}

	return 1 / myPowHelper(x, -n, memory)
}

func main() {
	fmt.Println(myPow(2, 10))  // 1024
	fmt.Println(myPow(2.1, 3)) // 9.261000000000001
	fmt.Println(myPow(2, -2))  // 0.25
}
