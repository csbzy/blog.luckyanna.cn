package alg

// 45. 跳跃游戏 II
// 给定一个长度为 n 的 0 索引整数数组 nums。初始位置为 nums[0]。

// 每个元素 nums[i] 表示从索引 i 向前跳转的最大长度。换句话说，如果你在 nums[i] 处，你可以跳转到任意 nums[i + j] 处:

// 0 <= j <= nums[i]
// i + j < n
// 返回到达 nums[n - 1] 的最小跳跃次数。生成的测试用例可以到达 nums[n - 1]。
// func jump(nums []int) int {
// 	// dp[n-1] = min(dp(n))
// 	if len(nums) == 0 {
// 		return 0
// 	}

// 	if len(nums) == 1 {
// 		if nums[0] != 0 {
// 			return 0
// 		}

// 		return 0

// 	}
// 	curIndex := len(nums) - 1
// 	loop := 0

// 	fmt.Println("curindex", curIndex)
// 	for curIndex > 0 {
// 		curIndex = findMax(nums, curIndex)
// 		if curIndex >= 0 {
// 			loop++
// 		}
// 	}
// 	return loop
// }

// func findMax(nums []int, right int) int {
// 	if right == 0 {
// 		return -1
// 	}
// 	max := 0
// 	maxIndex := 0
// 	index := 0
// 	isFirst := true
// 	for index < right {
// 		fmt.Println(nums[index] > max, (nums[index] >= right-index), nums[index], max, right)
// 		if nums[index] >= max && (nums[index] >= right-index) {
// 			fmt.Println("maxindex", maxIndex, "index", index)
// 			if isFirst {
// 				max = nums[index]
// 				maxIndex = index
// 				fmt.Println(maxIndex, index, max)
// 				isFirst = false
// 			} else if index <= maxIndex {
// 				max = nums[index]
// 				maxIndex = index
// 				fmt.Println("min", maxIndex, index, max)
// 			}

// 			fmt.Println("222maxindex", maxIndex)

// 		}
// 		index++
// 	}
// 	fmt.Println(maxIndex)
// 	return maxIndex
// }

// func jump(nums []int) int {
//     n := len(nums)
//     if n <= 1 {
//         return 0
//     }

//     maxPosition := 0
//     end := 0
//     steps := 0

//     for i := 0; i < n-1; i++ {
//         maxPosition = max(maxPosition, i+nums[i])

//         if i == end {
//             end = maxPosition
//             steps++
//         }
//     }

//     return steps
// }

// func max(a, b int) int {
//     if a > b {
//         return a
//     }
//     return b
// }

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 贪心算法
// 在这段代码中，我们看到了一个名为`jump`的函数，它实现了贪心算法来解决一个问题。贪心算法是一种在每个步骤中都选择当前最优解的策略，以期望最终得到全局最优解的算法。

// 在这个特定的问题中，我们要求通过跳跃来到达数组的最后一个元素，数组中的每个元素表示从当前位置最多可以跳跃的步数。我们的目标是找到到达最后一个元素所需的最小跳跃次数。

// 让我们来详细解释一下这个贪心算法的实现：

// 1. 首先，我们检查数组的长度`n`。如果数组长度小于等于1，意味着数组中只有一个元素或者没有元素，无需跳跃，直接返回0。

// 2. 我们定义三个变量：
//    - `maxPosition`：表示当前能够到达的最远位置。
//    - `end`：表示当前跳跃的边界，即在这个边界内的元素都可以通过一次跳跃到达。
//    - `steps`：表示已经跳跃的次数。

// 3. 接下来，我们使用一个循环来遍历数组中的元素，从索引0开始，直到倒数第二个元素。在循环中，我们执行以下操作：
//    - 更新`maxPosition`，将其设置为当前位置加上当前位置能够跳跃的最大步数（即`i+nums[i]`）和`maxPosition`中的较大值。
//    - 如果当前位置`i`等于`end`，意味着我们已经到达了上一次跳跃的边界，需要更新边界为`maxPosition`，并且跳跃次数加1。

// 4. 循环结束后，我们返回跳跃次数`steps`作为结果。

// 这个贪心算法的思想是，我们每次都选择能够跳跃的最远位置，以尽可能少的跳跃次数到达目标。通过不断更新`maxPosition`和`end`，我们可以保证每次跳跃都是当前最优解，最终得到到达最后一个元素所需的最小跳跃次数。

// 希望这个解释能够帮助你理解这段代码中的贪心算法实现。如果你有任何进一步的问题，请随时提问！
// func jump(nums []int) int {
// 	n := len(nums)
// 	if n <= 1 {
// 		return 0
// 	}

// 	maxPosition := nums[	]
// 	end := 0
// 	steps := 0

// 	for i := 0; i < n-1; i++ {
// 		maxPosition = max(maxPosition, i+nums[i])

// 		if i == end {
// 			end = maxPosition
// 			steps++
// 		}
// 	}

// 	return steps
// }

func jump(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	maxPosition := 0
	end := 0
	steps := 0
	for i, e := range nums {
		if i == n-1 {
			break
		}
		maxPosition = max(maxPosition, i+e)
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}
