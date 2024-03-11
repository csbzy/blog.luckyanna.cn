package alg

// // 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
func permute(nums []int) [][]int {
	var result [][]int
	backtrack(nums, []int{}, &result)
	return result
}

func backtrack(nums []int, path []int, result *[][]int) {
	if len(path) == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, path)
		*result = append(*result, temp)
	}

	for _, e := range nums {
		if contains(path, e) {
			continue
		}

		path = append(path, e)
		backtrack(nums, path, result)
		path = path[:len(path)-1]
	}
}

func contains(list []int, want int) bool {
	for _, e := range list {
		if e == want {
			return true
		}
	}
	return false
}

// 	func permute(nums []int) [][]int {
// 		var result [][]int
// 		backtrace(nums, []int{}, &result)
// 		return result
// 	}

// 	func backtrace(nums []int, path []int, result *[][]int) {
// 		if len(path) == len(nums) {
// 			// Make a copy of the path and append it to the result
// 			temp := make([]int, len(path))
// 			copy(temp, path)
// 			*result = append(*result, temp)
// 			return
// 		}

// 		for i := 0; i < len(nums); i++ {
// 			if contains(path, nums[i]) {
// 				continue
// 			}
// 			path = append(path, nums[i])
// 			backtrace(nums, path, result)
// 			path = path[:len(path)-1]
// 		}
// 	}

// 	func contains(nums []int, num int) bool {
// 		for _, n := range nums {
// 			if n == num {
// 				return true
// 			}
// 		}
// 		return false
// 	}
