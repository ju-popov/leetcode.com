package main

import "fmt"

const (
	maxInt32 = 2147483647
	minInt32 = -2147483648
)

func abs(value int) int {
	if value >= 0 {
		return value
	}

	return -value
}

func addResult(result int32, value int32) int32 {
	if value > 0 {
		if maxInt32-result < value {
			return maxInt32
		}
	} else {
		if minInt32-result > value {
			return minInt32
		}
	}

	return result + value
}

func divideHelper(dividend int, positiveDivisor int) (int32, int) {
	var result int32

	if abs(dividend) < positiveDivisor {
		return 0, dividend
	}

	if abs(dividend) >= positiveDivisor+positiveDivisor {
		var value int32
		value, dividend = divideHelper(dividend, positiveDivisor+positiveDivisor)
		result = addResult(result, value)
		result = addResult(result, value)
	}

	if abs(dividend) >= positiveDivisor {
		if dividend > 0 {
			dividend -= positiveDivisor
			result = addResult(result, 1)
		} else if dividend < 0 {
			dividend += positiveDivisor
			result = addResult(result, -1)
		}
	}

	return result, dividend
}

func divide(dividend int, divisor int) int {
	if divisor < 0 {
		dividend = -dividend
		divisor = -divisor
	}

	result, _ := divideHelper(dividend, divisor)

	return int(result)
}

func main() {
	fmt.Println(divide(10, 3))
	fmt.Println(divide(72, -3))
	fmt.Println(divide(0, 1))
	fmt.Println(divide(1, 1))
	fmt.Println(divide(maxInt32, 1))
	fmt.Println(divide(maxInt32+1, 1))
	fmt.Println(divide(minInt32, 1))
	fmt.Println(divide(minInt32-1, 1))
}
