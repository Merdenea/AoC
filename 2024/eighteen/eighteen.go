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

	bytes := make([][2]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		st := strings.Split(line, ",")

		y, _ := strconv.ParseInt(st[0], 10, 64)
		x, _ := strconv.ParseInt(st[1], 10, 64)
		bytes = append(bytes, [2]int{int(x), int(y)})
	}

	now := time.Now()
	one := partOneX2(bytes)
	fmt.Printf("Part I in [%s]: %d\n", time.Since(now).String(), one)

	now = time.Now()
	two := partTwo(bytes)
	fmt.Printf("Part II in [%s]: %s\n", time.Since(now).String(), two)

}

type node struct {
	x, y, steps int
}

func bfs(si, sj, ei, ej int, corrupted map[[2]int]int) int {
	q := make([]node, 0)
	q = append(q, node{si, sj, 0})
	seen := make(map[[2]int]int)
	for len(q) > 0 {
		current := q[0]
		q = q[1:]

		if seen[[2]int{current.x, current.y}] > 0 {
			continue
		}

		seen[[2]int{current.x, current.y}]++

		for _, d := range [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			ni, nj := current.x+d[0], current.y+d[1]
			if ni < 0 || nj < 0 || ni >= ei+1 || nj >= ej+1 {
				continue
			}
			if seen[[2]int{ni, nj}] > 0 || corrupted[[2]int{ni, nj}] > 0 {
				continue
			}
			if ni == ei && nj == ej {
				return current.steps + 1
			}
			q = append(q, node{ni, nj, current.steps + 1})
		}
	}
	return -1
}

func partOneX2(bytes [][2]int) int {
	corrupted := make(map[[2]int]int)

	si, sj := 0, 0
	ei, ej := 70, 70
	noOfBytes := 1024
	if len(bytes) < 1024 {
		// test vals
		ei, ej = 6, 6
		noOfBytes = 12
	}
	for i := 0; i < noOfBytes; i++ {
		corrupted[[2]int{bytes[i][0], bytes[i][1]}]++
	}

	res := bfs(si, sj, ei, ej, corrupted)
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func partTwo(bytes [][2]int) string {

	si, sj := 0, 0
	ei, ej := 70, 70
	noOfBytes := 1024

	if len(bytes) < 1024 {
		// test vals
		ei, ej = 6, 6
		noOfBytes = 12
	}

	left, right := noOfBytes, len(bytes)-1

	for left < right {
		mid := (right + left) / 2

		corrupted := make(map[[2]int]int)
		for i := 0; i < mid; i++ {
			corrupted[[2]int{bytes[i][0], bytes[i][1]}]++
		}

		res := bfs(si, sj, ei, ej, corrupted)

		if res > 0 {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return fmt.Sprintf("%d,%d", bytes[left-1][1], bytes[left-1][0])
}
