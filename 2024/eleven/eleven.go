package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, _ := os.Open(fmt.Sprintf("./%s", os.Args[1]))
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	line := scanner.Text()
	st := strings.Split(line, " ")

	val := make([]int, 0)
	for _, s := range st {
		n, _ := strconv.ParseInt(s, 10, 64)
		val = append(val, int(n))
	}

	now := time.Now()
	one := partOne(val)
	fmt.Printf("Part I in [%s]: %d\n", time.Since(now).String(), one)

	now = time.Now()
	two := partTwo(val)
	fmt.Printf("Part II in [%s]: %d\n", time.Since(now).String(), two)

}

var seen = make(map[string]int)

func key(i, j int) string {
	return fmt.Sprintf("%d-%d", i, j)
}

func partOne(val []int) int64 {
	total := 0
	for _, v := range val {
		total += updateStones(v, 25)
	}
	return int64(total)
}

func updateStones(val, blinks int) int {
	if v, ok := seen[key(val, blinks)]; ok {
		return v
	}

	next := blink(val)
	if blinks == 1 {
		return len(next)
	}
	res := 0
	for _, v := range next {
		res += updateStones(v, blinks-1)
	}
	seen[key(val, blinks)] = res

	return res
}

func blink(v int) []int {
	res := make([]int, 0)
	if v == 0 {
		res = append(res, 1)
	} else if evenNoDigits(v) {
		n1, n2 := splitNo(v)
		res = append(res, n1, n2)
	} else {
		res = append(res, v*2024)
	}
	return res
}

func evenNoDigits(i int) bool {
	return len(strconv.Itoa(i))%2 == 0
}

func splitNo(i int) (int, int) {
	s := strconv.Itoa(i)
	l, r := s[:len(s)/2], s[len(s)/2:]
	n1, _ := strconv.ParseInt(l, 10, 64)
	n2, _ := strconv.ParseInt(r, 10, 64)
	return int(n1), int(n2)
}

func partTwo(val []int) int64 {
	total := 0
	for _, v := range val {
		total += updateStones(v, 75)
	}
	return int64(total)
}
