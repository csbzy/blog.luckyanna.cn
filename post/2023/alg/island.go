// 200. 岛屿数量
// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

// 此外，你可以假设该网格的四条边均被水包围。
// 使用visited数组和直接修改岛屿为海两种实现方式有以下区别：

// 使用visited数组：在使用visited数组的实现方式中，我们使用一个二维布尔数组来记录每个位置是否被访问过。当我们遍历到一个陆地格子时，我们将其标记为已访问，并通过深度优先搜索（DFS）遍历与该陆地格子相邻的其他陆地格子。这种方式可以确保每个岛屿只被计数一次，避免重复计数。

// 直接修改岛屿为海：在直接修改岛屿为海的实现方式中，我们在遍历到一个陆地格子时，将其修改为海洋格子（'0'）。通过这种方式，我们可以避免使用额外的visited数组，节省了空间复杂度。但是，需要注意的是，这种方式会直接修改输入的二维网格，可能会影响到其他部分代码对原始数据的使用。

// 选择使用哪种实现方式取决于具体的需求和场景。如果需要保持原始数据不变，并且需要多次计算岛屿数量，那么使用visited数组的方式更合适。如果只需要计算一次岛屿数量，并且可以修改原始数据，那么直接修改岛屿为海的方式可能更简洁和高效。
package alg

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	lands := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				lands++
			}
		}
	}
	return lands
}

func dfs(grid [][]byte, i, j int) {
	if i < 0 || i == len(grid) || j < 0 || j == len(grid[0]) || grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}

// func numIslands(grid [][]byte) int {
// 	if len(grid) == 0 {
// 		return 0
// 	}

// 	landnums := 0
// 	// visited := make([][]bool, len(grid))
// 	// for i := range visited {
// 	// 	visited[i] = make([]bool, len(grid[0]))
// 	// }
// 	fmt.Println("-----------------------------------")

// 	for i := 0; i < len(grid); i++ {
// 		s := []byte{}

// 		for j := 0; j < len(grid[0]); j++ {
// 			s = append(s, grid[i][j])
// 			fmt.Println(i, j, grid[i][j])
// 			if grid[i][j] == '1' { //&& !visited[i][j] {
// 				landnums++

// 				dfs(grid, i, j) //, visited)
// 			}
// 		}
// 		fmt.Println(s)

// 	}
// 	fmt.Println("-----------------------------------")
// 	for i := 0; i < len(grid); i++ {
// 		s := []byte{}
// 		for j := 0; j < len(grid[0]); j++ {
// 			s = append(s, grid[i][j])
// 		}
// 		fmt.Println(s)
// 	}
// 	fmt.Println("-----------------------------------")

// 	return landnums
// }

// func dfs(grid [][]byte, i, j int) { //}, visited [][]bool) {
// 	if i >= len(grid) || i < 0 || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' { //|| visited[i][j] {
// 		return
// 	}

// 	grid[i][j] = '0'
// 	fmt.Println(i, j, grid[i][j])
// 	// visited[i][j] = true

// 	dfs(grid, i-1, j) //, visited)
// 	dfs(grid, i+1, j) //, visited)
// 	dfs(grid, i, j-1) //, visited)
// 	dfs(grid, i, j+1) //, visited)
// }
