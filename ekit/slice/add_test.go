package slice

import "fmt"

/**
 * @author: biao
 * @date: 2025/12/18 上午10:44
 * @description:
 */

func ExampleAdd() {
	res, _ := Add[int]([]int{1, 2, 3, 4}, 233, 2)
	fmt.Println(res)
	_, err := Add[int]([]int{1, 2, 3, 4}, 233, -1)
	fmt.Println(err)
	// Output:
	// [1 2 233 3 4]
	// ekit: 下标超出范围，长度 4, 下标 -1
}
