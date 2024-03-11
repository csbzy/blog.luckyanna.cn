// ###  无重复字符的最长子串
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

// 示例 1:
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 示例 2:

// 输入: s = "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 示例 3:

// 输入: s = "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//      请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

// 提示：

// 0 <= s.length <= 5 * 104
// s 由英文字母、数字、符号和空格组成

package alg

func longestsubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	var (
		left      = 0
		right     = 0
		maxLength = 0
		exists    = map[byte]int{}
	)

	for right < len(s) {
		if idx, ok := exists[s[right]]; ok {
			left = maxIn(left, idx+1)
		}

		exists[s[right]] = right
		if right-left+1 > maxLength {
			maxLength = right - left + 1
		}
		right++
	}
	return maxLength
}

// func longestsubstring(s string) int {
// 	var (
// 		left, right, max = 0, 0, 0
// 		exists           = map[byte]int{}
// 	)

// 	if right < len(s) {
// 		if idx, ok := exists[s[right]]; ok {
// 			left = maxIn(left, idx+1)
// 		}

// 		exists[s[right]] = right

// 		max = maxIn(max, right-left+1)
// 		right++

// 	}
// 	return max
// }

// 滑动窗口
func lengthOfLongestSubstringWindows(s string) int {
	var (
		left, right, max = 0, 0, 0
		exists           = make(map[string]int)
	)

	for right < len(s) {
		if idx, ok := exists[string(s[right])]; ok {
			left = maxIn(left, idx+1)
		}
		exists[string(s[right])] = right

		max = maxIn(max, right-left+1)
		right++

	}
	return max
}

func maxIn(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 暴力搜索遍历解
func lengthOfLongestSubstring(s string) int {
	// cur 当前最长字串
	// exists map出现过的
	var (
		length = len(s)
		max    = make([]string, 0, length)
	)
	for i := 0; i < length; i++ {
		cur := make([]string, 0, length)
		exists := make(map[string]struct{})
		for j := i; j < length; j++ {
			if _, exist := exists[string(s[j])]; exist {
				break
			}

			cur = append(cur, string(s[j]))
			exists[string(s[j])] = struct{}{}
			if len(cur) > length-i {
				break
			}

		}
		if len(cur) > len(max) {
			max = cur
		}
	}

	return len(max)
}
