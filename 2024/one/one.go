package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, _ := os.Open(fmt.Sprintf("./%s", os.Args[1]))
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	left := make([]int64, 0)
	right := make([]int64, 0)

	for scanner.Scan() {
		line := scanner.Text()
		st := strings.Split(line, "   ")

		l, _ := strconv.ParseInt(st[0], 10, 64)
		r, _ := strconv.ParseInt(st[1], 10, 64)

		left = append(left, l)
		right = append(right, r)
	}

	now := time.Now()
	fmt.Printf("Part I in [%s]: %d\n", time.Since(now).String(), partOne(left, right))

	now = time.Now()
	fmt.Printf("Part II in [%s]: %d\n", time.Since(now).String(), partTwo(left, right))

}

func partOne(left, right []int64) int64 {
	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	total := int64(0)

	for i := 0; i < len(left); i++ {
		if right[i] > left[i] {
			total += right[i] - left[i]
		} else {
			total += left[i] - right[i]
		}
	}
	return total
}

func partTwo(left, right []int64) int64 {
	freqMap := make(map[int64]int64)

	for _, r := range right {
		freqMap[r]++
	}

	total := int64(0)
	for _, l := range left {
		total += l * freqMap[l]
	}
	return total
}
