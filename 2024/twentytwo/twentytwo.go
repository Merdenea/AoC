package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gammazero/deque"
)

func main() {
	file, _ := os.Open(fmt.Sprintf("./%s", os.Args[1]))
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	secretNums := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		v, _ := strconv.ParseInt(line, 10, 64)
		secretNums = append(secretNums, int(v))
	}

	now := time.Now()
	one := partOne(secretNums)
	fmt.Printf("Part I in [%s]: %d\n", time.Since(now).String(), one)

	now = time.Now()
	two := partTwo(secretNums)
	fmt.Printf("Part II in [%s]: %d\n", time.Since(now).String(), two)

}

func mix(secret, value int) int {
	return secret ^ value
}

func prune(secret int) int {
	return secret % 16777216
}

func next(secret int) int {
	next := prune(mix(secret, secret*64))
	next = prune(mix(next, next/32))
	next = prune(mix(next, next*2048))
	return next
}

func partOne(secretNums []int) int64 {
	total := 0

	for _, n := range secretNums {
		for i := 0; i < 2000; i++ {
			n = next(n)
		}
		total += n
	}

	return int64(total)
}

// iffy -~800ms
func partTwo(secretNums []int) int64 {
	// for each buyer
	diffSeqsToPrices := make(map[int]map[[4]int]int)

	for index, n := range secretNums {
		diffSeqsToPrices[index] = make(map[[4]int]int)
		var prices deque.Deque[int]
		prices.PushBack(n % 10)
		for i := 0; i < 2000; i++ {
			n = next(n)
			prices.PushBack(n % 10)
			if prices.Len() > 4 {
				diffs := [4]int{prices.At(1) - prices.At(0), prices.At(2) - prices.At(1), prices.At(3) - prices.At(2), prices.At(4) - prices.At(3)}
				if _, ok := diffSeqsToPrices[index][diffs]; !ok {
					diffSeqsToPrices[index][diffs] = n % 10 // set only once
				}
				prices.PopFront()
			}
		}
	}

	allKeysSum := make(map[[4]int]int)

	for _, m := range diffSeqsToPrices {
		for k, v := range m {
			allKeysSum[k] += v
		}
	}

	max := 0
	for _, v := range allKeysSum {
		if v > max {
			max = v
		}
	}

	return int64(max)
}
