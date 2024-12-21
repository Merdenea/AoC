package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	file, _ := os.Open(fmt.Sprintf("./%s", os.Args[1]))
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	ll := scanner.Text()
	ts := strings.Split(ll, ", ")

	ps := make([]string, 0)
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		ps = append(ps, line)
	}

	now := time.Now()
	one := partOne(ts, ps)
	fmt.Printf("Part I in [%s]: %d\n", time.Since(now).String(), one)

	now = time.Now()
	two := partTwo(ts, ps)
	fmt.Printf("Part II in [%s]: %d\n", time.Since(now).String(), two)

}

var (
	towels = make(map[string]bool)
)

func partOne(ts, ps []string) int64 {
	maxLen := 0
	for _, t := range ts {
		towels[t] = true
		if len(t) > maxLen {
			maxLen = len(t)
		}
	}

	total := 0
	for _, p := range ps {
		if isPossible(p, maxLen) {
			total++
		}
	}

	return int64(total)
}

func isPossible(p string, maxLen int) bool {
	if towels[p] {
		return true
	}

	for currLen := maxLen; currLen >= 1; currLen-- {
		if len(p) > currLen {
			if isPossible(p[0:currLen], maxLen) && isPossible(p[currLen:], maxLen) {
				return true
			}
		}
	}
	return false
}

var memo = make(map[string]int)

func countPossible(targetPattern string, ts []string) int {
	if targetPattern == "" {
		return 1
	}

	if v, ok := memo[targetPattern]; ok {
		return v
	}

	count := 0
	for _, t := range ts {
		if strings.HasPrefix(targetPattern, t) {
			count += countPossible(targetPattern[len(t):], ts)
		}
	}

	memo[targetPattern] = count
	return count
}

func partTwo(ts, ps []string) int64 {
	total := 0
	for _, p := range ps {
		total += countPossible(p, ts)
	}

	return int64(total)
}
