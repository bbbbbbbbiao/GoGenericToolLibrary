package slice

import slice2 "github.com/bbbbbbbbiao/GoGenericToolLibrary/ekit/internal/slice"

/**
 * @author: biao
 * @date: 2025/12/18 上午10:39
 * @description:
 */

func Add[T any](slice []T, element T, index int) ([]T, error) {
	res, err := slice2.Add(slice, element, index)
	return res, err
}
