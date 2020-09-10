package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

var testCaseSqSize = 0
var testCaseCount = 0
var testCaseNumber = 0

func main() {
	content, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatalf("Cannot read text file: %s", err)
	}
	lines := strings.Split(string(content), "\r\n")
	loopInput(lines, 0)
	if testCaseCount != testCaseNumber{
		log.Fatalf("Invalid test case number: %s", strconv.Itoa(testCaseNumber))
	}
}

func loopInput(lines []string, index int) string {
	if index == len(lines) {
		return lines[index-1]
	} else {
		n := lines[index]
		if index == 0 {
			testCaseNo, errtestCaseNo := strconv.Atoi(n)
			if errtestCaseNo != nil {
				log.Fatalf("Cannot parse string to int: %s", errtestCaseNo)
			}
			if testCaseNo > 100 || testCaseNo < 1 {
				log.Fatalf("Invalid test case number: %s", n)
			}

			testCaseNumber = testCaseNo
		}

		if index%2 != 0 && index != 0 {
			testCaseCount++
			testCaseInt, errtestCaseInt := strconv.Atoi(n)
			if errtestCaseInt != nil {
				log.Fatalf("Cannot parse string to int: %s", errtestCaseInt)
			}
			if testCaseInt > 100 || testCaseInt < 1 {
				log.Fatalf("Invalid test case integer: %s", n)
			}

			testCaseSqSize = testCaseInt
		}

		if index%2 == 0 && index != 0 {
			squares := strings.Fields(n)
			if len(squares) == testCaseSqSize{
				loopSquares(squares, 0, 0)
			}else{
				log.Fatalf("Invalid test case square integer size: %s", strconv.Itoa(testCaseSqSize))
			}
		}
		return loopInput(lines, index+1)
	}
}

func loopSquares(squares []string, result int, index int) string {
	var tempOutput int

	if index == len(squares) {
		fmt.Println(result)
		return squares[index-1]
	}else{
		n := squares[index]
		testCaseSqInt, errtestCaseSqInt := strconv.Atoi(n)
		if errtestCaseSqInt != nil {
			log.Fatalf("Cannot parse string to int: %s", errtestCaseSqInt)
		}
		if testCaseSqInt > 100 || testCaseSqInt < -100 {
			log.Fatalf("Invalid test case square integer: %s", n)
		}

		if testCaseSqInt >= 0 {
			tempOutput = int(math.Pow(float64(testCaseSqInt), 2))
			result += tempOutput
		}
		return loopSquares(squares, result, index+1)
	}
}
