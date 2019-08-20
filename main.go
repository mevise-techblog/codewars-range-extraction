package main

import (
	"fmt"
	"strconv"
)

func main() {

	Solution([]int{-3, -2, 0, 3, 4, 5})
	Solution([]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20})
	Solution([]int{1, 4, 6, 7, 8, 10, 12, 13, 18, 19, 20, 100, 101, 7, 19, 20, 21, 1, 2, 3, 4, 5, 6})
}

func Solution(iArr []int) string {
	knownRangeIdxs := []int{} // index positions in iArr that are part of a known range group
	rangeStartIdxs := []int{}
	rangeEndIdxs := []int{}
	resultString := ""

	// First Pass : Find all the ranges in the array
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
					startOfRange := x - seriesCnt + 1
					endOfRange := startOfRange + seriesCnt - 1
					rangeStartIdxs = append(rangeStartIdxs, startOfRange)
					rangeEndIdxs = append(rangeEndIdxs, endOfRange)

					for y := startOfRange; y <= endOfRange; y++ {
						knownRangeIdxs = append(knownRangeIdxs, y)
					}

					seriesCnt = 1 // reset
				}
			}
		} else {
			seriesCnt = 1 // reset
		}
	}

	// This loop builds the string
	for x := 0; x < len(iArr); x++ {
		if contains(knownRangeIdxs, x) == false {
			resultString += strconv.Itoa(iArr[x])
			if x < len(iArr)-1 {
				resultString += ","
			}

		} else {
			isStartOfRange := contains(rangeStartIdxs, x)
			isEndOfRange := contains(rangeEndIdxs, x)

			if isStartOfRange {
				resultString += strconv.Itoa(iArr[x]) + "-"
			} else if isEndOfRange {
				resultString += strconv.Itoa(iArr[x])
				if x < len(iArr)-1 {
					resultString += ","
				}
			}
		}

	}

	fmt.Println("Result : " + resultString)
	return resultString
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
