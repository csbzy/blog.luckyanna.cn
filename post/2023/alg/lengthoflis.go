// 300. 最长递增子序列

// 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

// 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。
// 例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

package alg

func lengthOfLIS(nums []int) int {
	// dp[i] = if nums[i] > nums[j] => max(dp[i],dp[j]+1)
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	maxLength := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLength = max(dp[i], maxLength)
	}
	return maxLength

	// tails := []int{}
	// for _, n := range nums {
	// 	i := sort.Search(len(tails), func(i int) bool {
	// 		return tails[i] >= n
	// 	})
	// 	if i == len(tails) {
	// 		tails = append(tails, n)
	// 	} else {
	// 		tails[i] = n
	// 	}
	// }
	// return len(tails)
}

// func lengthOfLIS(nums []int) int {
// 	tails := []int{}
// 	for _, num := range nums {
// 		i := sort.Search(len(tails), func(i int) bool { return tails[i] >= num })
// 		if i == len(tails) {
// 			tails = append(tails, num)
// 		} else {
// 			tails[i] = num
// 		}
// 	}

// 	return len(tails)
// }

// func lengthOfLIS(nums []int) int {
// 	// dp[i] = max(dp[i-1],dp[i-1]+1)

// 	if len(nums) <= 1 {
// 		return len(nums)
// 	}

// 	var (
// 		maxLength = 0
// 		dp        = make([]int, len(nums))
// 	)

// 	dp[0] = 1
// 	for i := 0; i < len(nums); i++ {
// 		dp[i] = 1
// 		for j := 0; j < i; j++ {
// 			if nums[i] > nums[j] {
// 				// 这是因为`dp[i]`可能已经在之前的循环中被更新过，
// 				// 我们需要保证`dp[i]`始终表示以`nums[i]`结尾的最长递增子序列的长度。
// 				// 所以，`dp[i] = max(dp[i], dp[j]+1)`这行代码的目的是更新`dp[i]`，
// 				// 使其始终表示以`nums[i]`结尾的最长递增子序列的长度。
// 				dp[i] = max(dp[i], dp[j]+1)
// 			}
// 		}
// 		if dp[i] > maxLength {
// 			maxLength = dp[i]
// 		}
// 	}
// 	return maxLength
// }

// func lengthOfLIS1(nums []int) int {
// 	if len(nums) <= 1 {
// 		return len(nums)
// 	}

// 	// dp记录每个位置的最大递增子序列
// 	dp := make([]int, len(nums))
// 	dp[0] = 1
// 	maxLen := 1

// 	for i := 0; i < len(nums); i++ {
// 		maxval := 0
// 		for j := 0; j < i; j++ {
// 			if nums[i] > nums[j] {
// 				maxval = max(maxval, dp[j])
// 			}
// 		}
// 		dp[i] = maxval + 1
// 		maxLen = max(maxLen, dp[i])
// 	}
// 	return maxLen
// }

// func lengthOfLIS(nums []int) int {
// 	if len(nums) <= 1 {
// 		return len(nums)
// 	}

// 	dp := make([]int, len(nums))
// 	dp[0] = 1
// 	maxLen := 1

// 	for i := 1; i < len(nums); i++ {
// 		maxVal := 0
// 		for j := 0; j < i; j++ {
// 			if nums[i] > nums[j] {
// 				maxVal = max(maxVal, dp[j])
// 			}
// 		}
// 		dp[i] = maxVal + 1
// 		maxLen = max(maxLen, dp[i])
// 	}

// 	return maxLen
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func dfsLIS(nums []int, end int) []int {
// if end == 1 {
//     if nums[0] < nums[1] {
//         return []int{nums[0], nums[1]}
//     } else {
//         return []int{nums[1]}
//     }
// }

// s := dfsLIS(nums, end-1)
// if s == nil {
//     return nil
// }

// if nums[end] > s[len(s)-1] {
//     return append(s, nums[end])
// }

// if len(s) <= 2 {
//     if nums[end] < s[0] {
//         return []int{nums[end]}
//     } else {
//         return s
//     }
// }

// return s
// }

// func dfsLIS(nums []int, end int) []int {
//     if end == 1 {
//         if nums[0] < nums[1] {
//             return []int{nums[0], nums[1]}
//         } else {
//             return []int{nums[1]}
//         }
//     }

//     s := dfsLIS(nums, end-1)
//     fmt.Println("end", end, "value", nums[end], "s:", s)
//     if s == nil {
//         return nil
//     }

//     if nums[end] > s[len(s)-1] {
//         return append(s, nums[end])
//     }

//     if len(s) <= 2 {
//         if nums[end] < s[0] {
//             return []int{nums[end]}
//         } else {
//             return s
//         }
//     }

//     return s
// }

// func lengthOfLIS(nums []int) int {
// 	if len(nums) <= 1 {
// 		return len(nums)
// 	}

// 	l := dfsLIS(nums, len(nums)-1)
// 	fmt.Println(l)
// 	return len(l)
// }

// func dfsLIS(nums []int, end int) []int {
// 	if end == 1 {
// 		if nums[0] < nums[1] {
// 			return []int{nums[0], nums[1]}
// 		} else {
// 			return []int{nums[1]}
// 		}
// 	}

// 	s := dfsLIS(nums, end-1)
// 	fmt.Println("end", end, "value", nums[end], "s:", s)
// 	if s == nil {
// 		return nil
// 	}

// 	if nums[end] > s[len(s)-1] {
// 		return append(s, nums[end])
// 	}

// 	if len(s) <= 2 {
// 		if nums[end] < s[0] {
// 			return []int{nums[end]}
// 		} else if nums[end] > s[0] {
// 			return append(s[:len(s)-1], nums[end])
// 		}
// 	}

// 	// if nums[end] < s[len(s)-1] {
// 	// 	if len(s) <= 2 {
// 	//
// 	// 	}
// 	// 	// if nums[end] < s[0] {
// 	// 	// 	return []int{nums[end]}
// 	// 	// }else{
// 	// 	//     return []int{}
// 	// 	// }
// 	// 	fmt.Println("return:", append(s[:len(s)-1], nums[end]))

// 	// 	return append(s[:len(s)-1], nums[end])
// 	// } else
// 	return s
// }
