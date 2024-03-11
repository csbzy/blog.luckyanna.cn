package alg

import "fmt"

// 给你一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 判断你是否能够到达最后一个下标，如果可以，返回 true ；否则，返回 false 。
func canJump1(nums []int) bool {
	// dp[i]= dp[j]+(nums[j]> i-j)
	if len(nums) == 0 {
		return false
	}

	if len(nums) == 1 {
		return true
	}
	dp := make([]bool, len(nums))

	dp[0] = true
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && nums[j] >= i-j {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(nums)-1]
}

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return true
	}

	maxPosition := 0
	fmt.Println("-----")
	for i, num := range nums {
		if i == len(nums)-1 {
			continue
		}

		if maxPosition < i {
			return false
		}
		if i+num > maxPosition {
			maxPosition = i + num
		}

		if maxPosition >= len(nums)-1 {
			return true
		}
	}
	return false

	// n := len(nums)
	// if len(nums) == 0 {
	// 	return false
	// }

	// if len(nums) == 1 {
	// 	return true
	// }

	// maxPosition := nums[0]
	// for i, e := range nums {

	// 	if i == n-1 {
	// 		return false
	// 	}

	// 	if i > maxPosition {
	// 		return false
	// 	}

	// 	maxPosition = max(maxPosition, i+e)

	// 	if maxPosition >= len(nums)-1 {
	// 		return true
	// 	}
	// }
	// return false
}
