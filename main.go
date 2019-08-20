5package main

import (
	"fmt"
)

func main() {
	iArr := []int{1, 4, 6, 7, 8, 10, 12, 13, 18, 19, 20, 100, 101, 7, 19, 20, 21, 1, 2, 3, 4, 5, 6}
	knownrangeidxs := []int{} // index positions in iArr that are part of a known range group
	
	//First Pass : Find all the ranges in the array
	seriesCnt := 1
	for x := 1; x < len(iArr); x++ {
		endOfARange := false // reset

		if iArr[x] == iArr[x-1]+1 {
			seriesCnt++
			if seriesCnt >= 3 {
				if x == len(iArr)-1 { // last element
					endOfARange = true
				} else {
					if iArr[x+1] != iArr[x]+1 {
						endOfARange = true
					}
				}

				if endOfARange {
					//fmt.Printf("End of series found at x=%d\n", x)
					startOfRange := x - seriesCnt + 1
					endOfRange := x

					//rangesMap[x-seriesCnt][seriesCnt]
					fmt.Printf("Range Found :  start:%d  end:%d  count:%d\n", startOfRange, endOfRange, seriesCnt)
					// now we can save the idx's that are part of a known range

					for y := startOfRange; y <= endOfRange; y++ {
						fmt.Println(y)
						knownrangeidxs = append(knownrangeidxs, y)
					}

					seriesCnt = 1 // reset
				}
			}
		} else {
			seriesCnt = 1 // reset
		}
	}

	fmt.Println("known range idx's :")
	fmt.Printf("%v\n", knownrangeidxs)

}