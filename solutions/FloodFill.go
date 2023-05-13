package solutions

func FloodFill(image [][]int, sr int, sc int, color int) [][]int {
	return floodFill(image, sr, sc, image[sr][sc], color)
}

func floodFill(image [][]int, sr int, sc int, prevColor int, color int) [][]int {
	if sr < 0 || sr >= len(image) || sc < 0 || sc >= len(image[sr]) {
		return image
	}

	if image[sr][sc] == color || image[sr][sc] != prevColor {
		return image
	}

	image[sr][sc] = color

	image = floodFill(image, sr-1, sc, prevColor, color)
	image = floodFill(image, sr+1, sc, prevColor, color)
	image = floodFill(image, sr, sc-1, prevColor, color)
	image = floodFill(image, sr, sc+1, prevColor, color)

	return image
}
