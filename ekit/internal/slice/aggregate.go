package slice

import "github.com/bbbbbbbbiao/GoGenericToolLibrary/ekit"

/**
 * @author: biao
 * @date: 2025/12/18 上午11:05
 * @description: 求切片的和、最大值、最小值
 */

func Sum[T ekit.Number](slice []T) T {
	var sum T
	for _, val := range slice {
		sum += val
	}
	return sum
}

func Max[T ekit.RealNumber](slice []T) T {
	max := slice[0]

	for _, val := range slice {
		if val > max {
			max = val
		}
	}
	return max
}

func Min[T ekit.RealNumber](slice []T) T {
	min := slice[0]

	for _, val := range slice {
		if val < min {
			min = val
		}
	}
	return min
}
