package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	file, _ := os.Open(fmt.Sprintf("./%s", os.Args[1]))
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	m := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		pos := make([]int, 0)
		for _, r := range line {
			p, _ := strconv.ParseInt(string(r), 10, 64)
			if r == '.' {
				p = -123456
			}
			pos = append(pos, int(p))
		}
		m = append(m, pos)
	}

	now := time.Now()
	one := partOne(m)
	fmt.Printf("Part I in [%s]: %d\n", time.Since(now).String(), one)

	now = time.Now()
	two := partTwo(m)
	fmt.Printf("Part II in [%s]: %d\n", time.Since(now).String(), two)

}

var (
	seen map[string]int
	ends map[string]int
)

func key(i, j int) string {
	return fmt.Sprintf("%d-%d", i, j)
}

var paths = int64(0)

func partOne(m [][]int) int64 {
	total := 0
	for i, l := range m {
		for j, v := range l {
			if v == 0 {
				// seen = make(map[string]int)
				ends = make(map[string]int)
				dfs(m, i, j)
				updatePaths() // part2
				total += len(ends)
			}
		}
	}

	return int64(total)
}

func updatePaths() {
	for _, v := range ends {
		paths += int64(v)
	}
}

func partTwo(m [][]int) int64 {
	return paths
}

func dfs(m [][]int, i, j int) {
	k := key(i, j)
	// if _, ok := seen[k]; ok {
	// 	return
	// }
	if m[i][j] == 9 {
		ends[k]++
		return
	}

	for _, dir := range [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}} {
		ni, nj := i+dir[0], j+dir[1]
		if ni < 0 || nj < 0 || ni >= len(m) || nj >= len(m[0]) {
			continue
		}

		if m[ni][nj] == m[i][j]+1 {
			dfs(m, ni, nj)
		}
	}
	// seen[k]++
}
