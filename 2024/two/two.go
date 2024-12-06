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

	m := make([][]int64, 0)
	for scanner.Scan() {
		line := scanner.Text()
		st := strings.Split(line, " ")

		l := make([]int64, 0)
		for _, i := range st {
			n, _ := strconv.ParseInt(i, 10, 64)
			l = append(l, n)
		}

		m = append(m, l)

	}

	now := time.Now()
	fmt.Printf("Part I in [%s]: %d\n", time.Since(now).String(), partOne(m))

	now = time.Now()
	fmt.Printf("Part II in [%s]: %d\n", time.Since(now).String(), partTwo(m))

}

func partTwo(m [][]int64) int64 {
	total := 0

	for _, l := range m {
		if isValid(l) {
			total++
			continue
		}
		// wow so bad
		for i := 0; i < len(l); i++ {
			newArr := make([]int64, 0)

			for j, v := range l {
				if j == i {
					continue
				}
				newArr = append(newArr, v)
			}
			if isValid(newArr) {
				total++
				break
			}
		}
	}
	return int64(total)
}

func partOne(m [][]int64) int64 {
	total := 0

	for _, l := range m {
		if ok := isValid(l); ok {
			total++
		}
	}
	return int64(total)
}

func isValid(l []int64) bool {
	inc := true
	dec := true

	for i := 1; i < len(l); i++ {
		diff := l[i] - l[i-1]
		inc = inc && (diff > 0) && (diff >= 1 && diff <= 3)
		dec = dec && (diff < 0) && (diff >= -3 && diff <= -1)
	}

	if inc || dec {
		return true
	}
	return false
}
