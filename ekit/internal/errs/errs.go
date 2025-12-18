package errs

import "fmt"

/**
 * @author: biao
 * @date: 2025/12/18 上午10:28
 * @description:
 */

// 索引超出范围错误
func NewErrIndexOutOfRange(length int, index int) error {
	return fmt.Errorf("ekit: 下标超出范围，长度 %d, 下标 %d", length, index)
}
