package solve

type CMD struct {
	Len int
	Dir byte
}

func toInt(a bool) int {
	if a {
		return 1
	}
	return 0
}

func Solve(commands []CMD) int64 {
	curDir := commands[0].Dir
	var nextDir byte

	curLen := commands[0].Len
	nextLen := 0

	y := 0
	area := int64(0)
	clockwise := true

	for i := 1; i < len(commands); i++ {
		nextDir = commands[i].Dir
		nextLen = commands[i].Len

		nextClockwise := nextDir == ((curDir + 1) & 0b11)
		curLen += toInt(nextClockwise) + toInt(clockwise) - 1
		if curDir&1 != 0 {
			// up or down
			if curDir < 2 {
				//down
				y += curLen
			} else {
				//up
				y -= curLen
			}
		} else {
			// left or right
			if curDir > 1 {
				// left
				area += int64(curLen) * int64(y)
			} else {
				//right
				area += int64(-curLen) * int64(y)
			}
		}
		clockwise = nextClockwise
		curDir = nextDir
		curLen = nextLen
	}

	return area
}
