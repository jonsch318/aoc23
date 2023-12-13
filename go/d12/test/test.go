package test

type SpringType int

type Entry struct {
	springTypesLength int
	numbersLength     int
	remaining         int
}

type Counter struct {
	Cache map[Entry]int
}

func (c Counter) CountArrangements(springTypes string, numbers []int, remaining int) int {
	if len(springTypes) == 0 {
		if len(numbers) == 0 || (len(numbers) == 1 && remaining == 0) {
			return 1
		}
		return 0
	}

	entry := Entry{springTypesLength: len(springTypes), numbersLength: len(numbers), remaining: remaining}
	if v, exists := c.Cache[entry]; exists {
		return v
	}

	springType := springTypes[0]
	res := 0
	switch springType {
	case '?':
		switch remaining {
		case -1:
			with := 0
			if len(numbers) != 0 {
				with = c.CountArrangements(springTypes[1:], numbers, numbers[0]-1)
			}
			without := c.CountArrangements(springTypes[1:], numbers, -1)
			res = with + without
		case 0:
			res = c.CountArrangements(springTypes[1:], numbers[1:], -1)
		default:
			res = c.CountArrangements(springTypes[1:], numbers, remaining-1)
		}
	case '.':
		switch remaining {
		case -1:
			res = c.CountArrangements(springTypes[1:], numbers, remaining)
		case 0:
			res = c.CountArrangements(springTypes[1:], numbers[1:], -1)
		default:
		}
	case '#':
		switch remaining {
		case -1:
			if len(numbers) == 0 {
				break
			}
			res = c.CountArrangements(springTypes[1:], numbers, numbers[0]-1)
		case 0:
		default:
			res = c.CountArrangements(springTypes[1:], numbers, remaining-1)
		}
	default:
		panic(springType)
	}

	c.Cache[entry] = res
	return res
}
