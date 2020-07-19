/*
 * @lc app=leetcode.cn id=874 lang=golang
 *
 * [874] 模拟行走机器人
 */

// @lc code=start
func robotSim(commands []int, obstacles [][]int) int {
	obstacleSet := map[string]bool{}
	for _, obstacle := range obstacles {
		key := fmt.Sprintf("%d-%d", obstacle[0], obstacle[1])
		obstacleSet[key] = true
	}
	max := 0
	direction := 0
	xFactor := []int{0, 1, 0, -1}
	yFactor := []int{1, 0, -1, 0}
	x, y := 0, 0

	for _, command := range commands {
		if command == -1 {
			direction = (direction + 1) % 4
		} else if command == -2 {
			direction = (direction + 3) % 4
		} else {
			for i := 0; i < command; i++ {
				nextX := x + xFactor[direction]
				nextY := y + yFactor[direction]
				key := fmt.Sprintf("%d-%d", nextX, nextY)
				if obstacleSet[key] {
					break
				}
				x = nextX
				y = nextY
				distance := x*x + y*y
				if distance > max {
					max = distance
				}
			}
		}
	}
	return max
}
// @lc code=end

