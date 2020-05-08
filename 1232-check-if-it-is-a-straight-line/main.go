package main

func checkStraightLine(coordinates [][]int) bool {
	x1, y1 := coordinates[0][0], coordinates[0][1]
	x2, y2 := coordinates[1][0], coordinates[1][1]

	a := float64(y1-y2) / float64(x1-x2)
	b := float64(x1*y2-x2*y1) / float64(x1-x2)

	for i := 2; i < len(coordinates); i++ {
		xi := float64(coordinates[i][0])
		yi := float64(coordinates[i][1])
		if yi != a*xi+b {
			return false
		}
	}

	return true
}
