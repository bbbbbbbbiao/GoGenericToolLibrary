package slice

import "github.com/bbbbbbbbiao/GoGenericToolLibrary/ekit/internal/errs"

/**
 * @author: biao
 * @date: 2025/12/18 上午10:23
 * @description: 向切片中添加元素
 */

func Add[T any](slice []T, element T, index int) ([]T, error) {
	length := len(slice)

	if index < 0 || index > length {
		return nil, errs.NewErrIndexOutOfRange(length, index)
	}

	var ele T
	slice = append(slice, ele)

	for i := length; i > index; i-- {
		slice[i] = slice[i-1]
	}
	slice[index] = element
	return slice, nil
}
