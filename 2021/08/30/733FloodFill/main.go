package main

func main() {
	floodFill([][]int{
		{0, 0, 0},
		{0, 0, 0},
	}, 0, 0, 2)
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	stack := make([][]int, 0)
	stack = append(stack, []int{sr, sc})
	row, col := len(image), len(image[0])
	oldColor := image[sr][sc]
	for len(stack) != 0 {
		size := len(stack)
		for i := range stack {
			for _, v := range move {
				nextr := stack[i][0] + v[0]
				nextc := stack[i][1] + v[1]
				if nextr >= 0 && nextr < row && nextc >= 0 && nextc < col && image[nextr][nextc] == oldColor {
					stack = append(stack, []int{nextr, nextc})
				}
			}
			image[stack[i][0]][stack[i][1]] = newColor
		}
		stack = stack[size:]
	}

	return image
}

var move = [][]int{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}
