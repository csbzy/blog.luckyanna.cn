package alg

import "fmt"

// 139. 单词拆分
// 给你一个字符串 s 和一个字符串列表 wordDict 作为字典。请你判断是否可以利用字典中出现的单词拼接出 s 。
// 这个问题也可以使用动态规划来解决。我们可以定义一个布尔数组 `dp`，其中 `dp[i]`
// 表示字符串 `s` 的前 `i` 个字符是否可以用 `wordDict` 中的单词表示。

// 我们初始化 `dp[0]` 为 `true`，表示空字符串可以被表示。
// 然后，我们从左到右计算所有的 `dp[i]`。
// 对于每个 `i`，我们检查所有以 `i` 结尾的子字符串，如果这个子字符串在 `wordDict` 中，
// 并且剩余的部分可以被表示（即 `dp[j]` 为 `true`），那么 `dp[i]` 就可以被表示。
func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	out := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelNodes := []int{}
		length := len(queue)
		for i := 0; i < length; i++ {
			node := queue[0]
			fmt.Println(i, length, queue, node.Left, node.Right)
			levelNodes = append(levelNodes, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			if len(queue) == 1 {
				fmt.Println("last one break")
				queue = []*TreeNode{}
				break
			} else {
				queue = queue[1:]
			}
		}

		fmt.Println("----", queue, len(queue))
		out = append(out, levelNodes)
	}
	return out
}
