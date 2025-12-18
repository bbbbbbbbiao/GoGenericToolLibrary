package slice

import "github.com/bbbbbbbbiao/GoGenericToolLibrary/ekit/internal/errs"

/**
 * @author: biao
 * @date: 2025/12/18 上午10:58
 * @description: 删除切片中的元素
 */

func Delete[T any](slice []T, index int) ([]T, T, error) {
	length := len(slice)
	if index < 0 || index >= length {
		var ele T
		return nil, ele, errs.NewErrIndexOutOfRange(length, index)
	}

	ele := slice[index]

	for i := index; i < length-1; i++ {
		slice[i] = slice[i+1]
	}
	slice = slice[:length-1]
	return slice, ele, nil
}
