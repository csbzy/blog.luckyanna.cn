package alg

// 给定两个字符串，求出他们的最长公共子序列
func MaxSameSubString(first string, second string) string {
	var (
		charExistMap  = make(map[rune]struct{})
		sameSubString = []rune{}
	)

	for _, e := range first {
		charExistMap[e] = struct{}{}
	}

	for _, e := range second {
		if _, ok := charExistMap[e]; ok {
			sameSubString = append(sameSubString, e)
		}
	}
	return string(sameSubString)
}
