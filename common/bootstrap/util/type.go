package util

import "fmt"

func SafeToInt32(x int) (int32, error) {
	// 检查值的范围
	if x < -(1<<31) || x > (1<<31)-1 {
		return 0, fmt.Errorf("int value out of range for int32")
	}

	// 使用类型断言进行转换
	return int32(x), nil
}
