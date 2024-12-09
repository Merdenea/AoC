package main

import (
	cmm "aoc/common"
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	file, _ := os.Open(fmt.Sprintf("./%s", os.Args[1]))
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	m := make([][]rune, 0)
	for scanner.Scan() {
		line := scanner.Text()
		m = append(m, []rune(line))
	}

	now := time.Now()
	one := partOne(m)
	fmt.Printf("Part I in [%s]: %d\n", time.Since(now).String(), one)

	now = time.Now()
	two := partTwoX2(m)
	fmt.Printf("Part II in [%s]: %d\n", time.Since(now).String(), two)

}

var (
	seen map[string]int
	adj  = make(map[string][]cmm.Pair)
)

func partOne(m [][]rune) int64 {
	seen = make(map[string]int)
	i, j, dir := findStart(m)

	for i >= 0 && j >= 0 && i < len(m) && j < len(m[0]) {

		if i+int(dir.I) < 0 || j+int(dir.J) < 0 || i+int(dir.I) >= len(m) || j+int(dir.J) >= len(m[0]) {
			seen[getKey(i, j)]++
			break
		}

		if m[i+int(dir.I)][j+int(dir.J)] == '#' {
			dir = cmm.TurnRight(dir)
		} else {
			seen[getKey(i, j)]++
			i += int(dir.I)
			j += int(dir.J)
		}
	}

	return int64(len(seen))
}

func getKey(i, j int) string {
	return fmt.Sprintf("%d-%d", i, j)
}

func getKey2(i, j int, dir cmm.Direction) string {
	return fmt.Sprintf("%d-%d-%d-%d", i, j, dir.I, dir.J)
}

func findStart(m [][]rune) (int, int, cmm.Direction) {
	for i, l := range m {
		for j, r := range l {
			if r != '#' && r != '.' {
				return i, j, findDirection(r)
			}
		}
	}
	return -1, -1, cmm.Direction{}
}

func findDirection(r rune) cmm.Direction {
	switch r {
	case '^':
		return cmm.UP
	case 'v':
		return cmm.DOWN
	case '>':
		return cmm.RIGHT
	case '<':
		return cmm.LEFT
	}
	return cmm.Direction{}
}

// = 6 seconds.....
func partTwoX2(m [][]rune) int64 {
	si, sj, dir := findStart(m)
	total := int64(0)

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			if i == si && j == sj {
				continue
			}
			if m[i][j] != '.' {
				continue
			}
			if seen[getKey(i, j)] == 0 {
				continue
			}

			m[i][j] = '#'

			total += hasCycle(si, sj, dir, m)
			m[i][j] = '.'
		}
	}

	return total
}

func hasCycle(i, j int, dir cmm.Direction, m [][]rune) int64 {
	seenDir := make(map[string]int)

	for i >= 0 && j >= 0 && i < len(m) && j < len(m[0]) {

		if seenDir[getKey2(i, j, dir)] > 0 {
			return 1
		}

		if i+int(dir.I) < 0 || j+int(dir.J) < 0 || i+int(dir.I) >= len(m) || j+int(dir.J) >= len(m[0]) {
			break
		}

		if m[i+int(dir.I)][j+int(dir.J)] == '#' {
			dir = cmm.TurnRight(dir)
		} else {
			seenDir[getKey2(i, j, dir)]++
			i += int(dir.I)
			j += int(dir.J)
		}

	}
	return 0
}
