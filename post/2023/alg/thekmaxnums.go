// 215. 数组中的第K个最大元素

// 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

// 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。

package alg

import "fmt"

type IntSlice []int

func (s IntSlice) Len() int {
	return len(s)
}

func (s IntSlice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s IntSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func findKthLargest(nums []int, k int) int {
	// if len(nums) == 0 {
	// 	return 0
	// }
	// sort.Sort(IntSlice(nums))
	// return nums[len(nums)-k]
	return quickSelect(nums, 0, len(nums)-1, k)
}

// 在这段代码中，pivotIndex 是用来表示快速选择算法中的枢轴元素的索引。
// 快速选择算法是一种用于在无序数组中查找第 k 大元素的算法。
// 它的基本思想是通过每次选择一个枢轴元素，将数组分为两部分，一部分比枢轴元素小，一部分比枢轴元素大。
// 然后根据枢轴元素的位置，可以确定第 k 大元素在哪一部分中，从而缩小问题的规模。

// 在这段代码中，pivotIndex 的值是通过调用 partition 函数得到的。
// partition 函数根据给定的左右边界，在数组 nums 中选择一个枢轴元素，
// 并将数组重新排列，使得枢轴元素左边的元素都比它小，右边的元素都比它大。
// 然后，pivotIndex 的值就是枢轴元素的最终位置。

// 在快速选择算法中，通过比较 pivotIndex 和 len(nums)-k 的值，
// 可以确定第 k 大元素在左边还是右边。如果 pivotIndex 等于 len(nums)-k，
// 则找到了第 k 大元素；如果 pivotIndex 大于 len(nums)-k，则第 k 大元素在左边部分，
// 继续在左边部分进行快速选择；如果 pivotIndex 小于 len(nums)-k，则第 k 大元素在右边部分，
// 继续在右边部分进行快速选择。

// 总之，pivotIndex 的值在快速选择算法中起到了确定枢轴元素位置的作用，
// 帮助我们缩小问题的规模并找到第 k 大元素。
func quickSelect(nums []int, left, right, k int) int {
	if len(nums) == 0 {
		return 0
	}
	pivotIndex := partition(nums, left, right)
	fmt.Println("pivotIndex", pivotIndex, len(nums)-k+1)
	if pivotIndex == len(nums)-k {
		return nums[pivotIndex]
	} else if pivotIndex > len(nums)-k {
		return quickSelect(nums, left, pivotIndex-1, k)
	} else {
		return quickSelect(nums, pivotIndex+1, right, k)
	}
}

func partition(nums []int, left, right int) int {
	if left < 0 {
		left = 0
	}
	if right >= len(nums) {
		right = len(nums) - 1
	}

	provit := nums[right]
	i := left
	for j := left; j < right; j++ {
		if nums[j] <= provit {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		fmt.Println("nums=", nums)
	}
	nums[i], nums[right] = nums[right], nums[i]
	fmt.Println("return i", i)
	return i
	// provit := nums[right]
	// i := left
	// for j := left; j < right; j++ {
	// 	if nums[j] <= provit {
	// 		nums[i], nums[j] = nums[j], nums[i]
	// 		i++
	// 	}
	// }
	// nums[i], nums[right] = nums[right], nums[i]
	// return i
}

// func threeSum(nums []int) [][]int {
// 	sum := map[int]int{}
// 	answers := [][]int{}
// 	visited := map[string]bool{}
// 	for i, e := range nums {
// 		sum[e] = i
// 	}
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {

// 			want := 0 - nums[i] - nums[j]
// 			if n, ok := sum[want]; ok {
// 				if n == i || n == j {
// 					continue
// 				}
// 				if visited(toString(nums[i], nums[j], nums[n])) {
// 					continue
// 				}

// 				answer := []int{nums[i], nums[j], nums[n]}

// 				answers = append(answers, answer)

// 				visited[toString(nums[i], nums[j], nums[n])] = true

// 			}
// 		}
// 	}
// 	return answers
// }

// func toString(a, b, c int) string {
// 	var t int
// 	if a > b {
// 		t = a
// 		a = b
// 		b = t
// 	}
// 	if a > c {
// 		t = a
// 		a = c
// 		c = t
// 	}
// 	if b > c {
// 		t = b
// 		b = c
// 		c = t
// 	}
// 	return fmt.Sprintf("%v%v%v", a, b, c)
// }

// if n<i {
//     return fmt.Sprintf("%v%v%v",n,i,j)
// }else if n>i && n<j {
//             return fmt.Sprintf("%v%v%v",i,n,j)

// }else {
// return fmt.Sprintf("%v%v%v",i,j,n)
// }
