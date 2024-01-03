package main

import (
	"errors"
	"fmt"
)

func freq(Sum int) (int, error) {
    if Sum > 12 || Sum < 2 {
        return 0, errors.New("Sum can only be in range 2-12")
    }

    if Sum <= 7 {
        return Sum-1, nil
    } else {
        return 12 - Sum + 1, nil
    }

}



func ques1()  {
    var (
        dieA [6]int = [6]int{1,2,3,4,5,6}
        dieB [6]int = [6]int{1,2,3,4,5,6}
        total int = 0
    )

    for _, i := range dieA {
        for _, j := range dieB {
            fmt.Printf("[%d,%d]\t", i, j)
            total += 1
        }
        fmt.Println()
    }

    fmt.Printf("Total: %d possible combinations\n", total)

}

func ques2()  {
    var (
        dieA [6]int = [6]int{ 1, 2, 3, 4, 5, 6, }
        dieB [6]int = [6]int{ 1, 2, 3, 4, 5, 6}
        freq map[int]int = make(map[int]int)
    )

    for _, i := range dieA {
        for _, j := range dieB {
            freq[i+j] = freq[i+j] + 1
            fmt.Printf("%d\t", i + j)
        }
        fmt.Println()
    }

    fmt.Println("Sum\tFrequency")
    for key := 2; key <= 12; key++ {
        fmt.Printf("%d:\t%d \n", key, freq[key])
    }

}

func ques3()  {

    fmt.Println("Sum\tFreq\tProbability")
    for key := 2; key <= 12; key++ {
        f, _ := freq(key)
        fmt.Printf("%d:\t%d\t%f \n", key, f, float64(f)/36.0)
    }

}

func main() {
    ques1()
    ques2()
    ques3()
}
