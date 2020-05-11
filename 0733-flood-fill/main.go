package main

func dfs(image *[][]int, height int, width int, row int, column int, oldColor int, newColor int) {
	if (*image)[row][column] != oldColor {
		return
	}

	(*image)[row][column] = newColor
	if row-1 >= 0 && (*image)[row-1][column] == oldColor {
		dfs(image, height, width, row-1, column, oldColor, newColor)
	}
	if row+1 < height && (*image)[row+1][column] == oldColor {
		dfs(image, height, width, row+1, column, oldColor, newColor)
	}
	if column-1 >= 0 && (*image)[row][column-1] == oldColor {
		dfs(image, height, width, row, column-1, oldColor, newColor)
	}
	if column+1 < width && (*image)[row][column+1] == oldColor {
		dfs(image, height, width, row, column+1, oldColor, newColor)
	}
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] != newColor {
		dfs(&image, len(image), len(image[0]), sr, sc, image[sr][sc], newColor)
	}

	return image
}
