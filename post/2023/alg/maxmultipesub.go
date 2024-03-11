package alg

import "context"

// 152. 乘积最大子数组

// 给你一个整数数组 nums ，请你找出数组中乘积最大的非空连续子数组（该子数组中至少包含一个数字）
// ，并返回该子数组所对应的乘积。

// 测试用例的答案是一个 32-位 整数。

// 子数组 是数组的连续子序列。

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	var (
		maxProduct = nums[0]
		minProduct = nums[0]
		result     = nums[0]
	)
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxProduct, minProduct = minProduct, maxProduct
		}

		maxProduct = max(nums[i], maxProduct*nums[i])
		minProduct = min(nums[i], minProduct*nums[i])
		if maxProduct > result {
			result = maxProduct
		}
	}
	return result
}

func maxProduct1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxProduct := nums[0]
	minProduct := nums[0]
	result := nums[0]

	for i := 1; i < len(nums); i++ {

		// maxProduct表示前n个数,连续相乘后最大的结果>=0
		// minProduct 表示前n个数,连续相乘后最小的结果<0或等于nums[i]
		if nums[i] < 0 {
			maxProduct, minProduct = minProduct, maxProduct
		}

		maxProduct = max(nums[i], maxProduct*nums[i])
		minProduct = min(nums[i], minProduct*nums[i])

		result = max(result, maxProduct)
		context.WithValue(parent Context, key any, val any)
	}

	return result
}

// func max(a, b int) int {
//     if a > b {
//         return a
//     }
//     return b
// }

// func min(a, b int) int {
//     if a < b {
//         return a
//     }
//     return b
// }

// func maxProduct(nums []int) int {
// 	if len(nums) < 1 {
// 		return 0
// 	}

// 	if len(nums) < 2 {
// 		return nums[0]
// 	}

// 	maxMutiple := nums[0]
// 	conMutiple := nums[0]
// 	isContinue := true
// 	for i := 1; i <= len(nums)-1; i++ {
// 		fmt.Println("i:", i, "data:", nums[i])
// 		if nums[i] == 0 {
// 			isContinue = false
// 			conMutiple = 0
// 			maxMutiple = max(maxMutiple, 0)
// 			continue
// 		}

// 		if isContinue {
// 			conMutiple = conMutiple * nums[i]
// 			maxMutiple = max(max3(maxMutiple, conMutiple, nums[i]), nums[i]*(nums[i-1])*(nums[i+1]))
// 		} else {
// 			maxMutiple = max(maxMutiple, nums[i])
// 			conMutiple = nums[i]
// 		}

// 		isContinue = true
// 	}
// 	fmt.Println("max cur", maxMutiple)
// 	return maxMutiple
// }

// func max3(a, b, c int) int {
// 	fmt.Println("max3", a, b, c, max(a, max(b, c)))
// 	return max(a, max(b, c))
// }
