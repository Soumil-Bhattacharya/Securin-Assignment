package main

import (
	"errors"
	"fmt"
)

func undoom(dieA, dieB [6]int) ([6]int, [6]int, error) {
	var standard_dice = [6]int{1, 2, 3, 4, 5, 6}
    
	if (dieA != standard_dice) || (dieB != standard_dice) {
		return [6]int{}, [6]int{}, errors.New("Not standard pair of dice")
	}

	return [6]int{1, 2, 2, 3, 3, 4}, [6]int{1, 3, 4, 5, 6, 8}, nil
}

func main() {
	var (
		dieA [6]int = [6]int{1, 2, 3, 4, 5, 6}
		dieB [6]int = [6]int{1, 2, 3, 4, 5, 6}
	)

	fmt.Println("Die_A=", dieA)
	fmt.Println("Die_B=", dieB)

	New_Die_A, New_Die_B, err := undoom(dieA, dieB)

	if err != nil {
		fmt.Println("Standard dice should be used")
	} else {
		fmt.Println("New_Die_A=", New_Die_A)
		fmt.Println("New_Die_B=", New_Die_B)

	}
}
