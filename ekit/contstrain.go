package ekit

/**
 * @author: biao
 * @date: 2025/12/18 上午11:06
 * @description: 约束泛型类型为数字类型
 */

type RealNumber interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
}

type Number interface {
	RealNumber | ~complex64 | ~complex128
}
