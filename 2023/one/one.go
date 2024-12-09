package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./one/in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := int64(0)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		val := getValue(scanner.Text())
		// fmt.Println(val)
		sum += val
	}
	fmt.Println(sum)
}

func getValue(l string) int64 {
	// fmt.Println(l)
	digits := map[rune]bool{
		'1': true,
		'2': true,
		'3': true,
		'4': true,
		'5': true,
		'6': true,
		'7': true,
		'8': true,
		'9': true,
		'0': true,
	}
	first := int64(-1)
	lIndex := -1
	rIndex := -1
	last := int64(-1)
	for i, c := range l {
		if digits[c] {
			if first == -1 {
				first, _ = strconv.ParseInt(string(c), 10, 64)
				last = first
				lIndex = i
				rIndex = i
			} else {
				last, _ = strconv.ParseInt(string(c), 10, 64)
				rIndex = i
			}
		}
	}

	fI, fV, rI, rV := getLetterValue(l)

	if lIndex == -1 || (fI != -1 && fI < lIndex) {
		first = int64(fV)
	}
	if rIndex == -1 || (rI != -1 && rI > rIndex) {
		last = int64(rV)
	}
	ret := 10*first + last
	// fmt.Println("ret: ", ret)
	return ret
}

func getLetterValue(l string) (int, int, int, int) {
	// fmt.Println(l)
	vals := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	lIndex := -1
	lValue := -1

	rIndex := -1
	rValue := -1

	for k, v := range vals {
		i := strings.Index(l, k)
		if i >= 0 {
			if lIndex == -1 || i < lIndex {
				lIndex = i
				lValue = v
			}
		}
		j := strings.LastIndex(l, k)
		if j >= 0 {
			if j > rIndex {
				rIndex = j
				rValue = v
			}
		}
	}
	// fmt.Println(lIndex, lValue, rIndex, rValue)

	return lIndex, lValue, rIndex, rValue
}
