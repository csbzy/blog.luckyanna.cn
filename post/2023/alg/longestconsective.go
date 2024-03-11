// 128. 最长连续序列

// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
package alg

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}

	var (
		maxLength = 0
		exists    = map[int]struct{}{}
	)

	for _, e := range nums {
		exists[e] = struct{}{}
	}

	for i := 0; i < len(nums); i++ {
		cur := nums[i]

		curMaxLength := 0
		// 从每一个连续序列的起始值开始计算
		if _, exist := (exists[cur-1]); exist {
			continue
		}

		exist := true
		for exist {
			cur++
			curMaxLength++
			_, exist = exists[cur]
		}

		if maxLength < curMaxLength {
			maxLength = curMaxLength
		}

	}
	return maxLength
}

// func longestConsecutive(nums []int) int {
// 	m := map[int]int{}
// 	for _, e := range nums {
// 		m[e] = e
// 	}

// 	maxLength := 0
// 	curMax := 0
// 	for _, e := range m {
// 		if _, ok := m[e-1]; ok {
// 			continue
// 		}
// 		curE := e
// 		ok := true
// 		for ok {
// 			curMax++
// 			curE++
// 			_, ok = m[curE]
// 		}
// 		if curMax > maxLength {
// 			maxLength = curMax
// 		}
// 		curMax = 0
// 	}
// 	return maxLength
// }
