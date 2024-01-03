package partb

import (
	"errors"
)

func undoom(dieA, dieB [6]int) ([6]int, [6]int, error) {
	var standard_dice = [6]int{1, 2, 3, 4, 5, 6}

	if (dieA != standard_dice) || (dieB != standard_dice) {
		return [6]int{}, [6]int{}, errors.New("Not standard pair of dice")
	}

	return [6]int{1, 2, 2, 3, 3, 4}, [6]int{1, 3, 4, 5, 6, 8}, nil
}
