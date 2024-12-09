package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	UP    = "UP"
	DOWN  = "DOWN"
	LEFT  = "LEFT"
	RIGHT = "RIGHT"
)

var (
	visitedGlobal map[string]int
	visited       map[string]int
)

func main() {
	file, _ := os.Open("./sixteen/in")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	m := make([][]rune, 0)
	for scanner.Scan() {
		l := scanner.Text()
		m = append(m, []rune(l))
	}

	visitedGlobal = make(map[string]int)

	// Part One
	now := time.Now()
	one := getEnergisedCount(m, 0, 0, RIGHT)
	fmt.Printf("Part One [%d] in %s\n", one, time.Now().Sub(now))

	// Part Two
	now = time.Now()
	max := -1
	dir := RIGHT
	for i := 0; i < len(m); i++ {
		if _, ok := visitedGlobal[getKey(i, 0, dir)]; ok {
			continue
		}
		v := getEnergisedCount(m, i, 0, dir)
		if v > max {
			max = v
		}
	}

	dir = DOWN
	for j := 0; j < len(m[0]); j++ {
		if _, ok := visitedGlobal[getKey(0, j, dir)]; ok {
			continue
		}
		v := getEnergisedCount(m, 0, j, dir)
		if v > max {
			max = v
		}
	}

	dir = LEFT
	for i := 0; i < len(m); i++ {
		if _, ok := visitedGlobal[getKey(i, len(m[i])-1, dir)]; ok {
			continue
		}
		v := getEnergisedCount(m, i, len(m[i])-1, dir)
		if v > max {
			max = v
		}
	}

	dir = UP
	for j := 0; j < len(m[0]); j++ {
		if _, ok := visitedGlobal[getKey(len(m)-1, j, dir)]; ok {
			continue
		}
		v := getEnergisedCount(m, len(m)-1, j, dir)
		if v > max {
			max = v
		}
	}

	fmt.Printf("Part Two [%d] in %s\n", max, time.Now().Sub(now))
}

func getEnergisedCount(m [][]rune, i, j int, dir string) int {
	visited = make(map[string]int)
	traceBeam(m, i, j, dir)

	visitedAnyDir := make(map[string]int)

	for k, v := range visited {
		newK := k[:strings.Index(k, ":")]
		visitedAnyDir[newK] += v
	}
	return len(visitedAnyDir)
}

func traceBeam(m [][]rune, i, j int, dir string) {
	if i < 0 || i >= len(m) {
		return
	}
	if j < 0 || j >= len(m[i]) {
		return
	}

	key := getKey(i, j, dir)
	if _, ok := visited[key]; ok {
		return
	}
	visited[key]++
	visitedGlobal[key]++

	switch m[i][j] {
	case '.':
		ii, jj, d := getNextEmpty(i, j, dir)
		traceBeam(m, ii, jj, d)
		break
	case '/', '\\':
		ii, jj, d := getNextMirror(i, j, dir, m[i][j])
		traceBeam(m, ii, jj, d)
		break
	case '|':
		if dir == UP || dir == DOWN {
			ii, jj, d := getNextEmpty(i, j, dir)
			traceBeam(m, ii, jj, d)
		} else {
			upI, upJ, upD := up(i, j)
			downI, downJ, downD := down(i, j)
			traceBeam(m, upI, upJ, upD)
			traceBeam(m, downI, downJ, downD)
		}
		break
	case '-':
		if dir == LEFT || dir == RIGHT {
			ii, jj, d := getNextEmpty(i, j, dir)
			traceBeam(m, ii, jj, d)
		} else {
			leftI, leftJ, leftD := left(i, j)
			rightI, rightJ, rightD := right(i, j)
			traceBeam(m, leftI, leftJ, leftD)
			traceBeam(m, rightI, rightJ, rightD)
		}
		break
	}
}

func getNextEmpty(i, j int, dir string) (int, int, string) {
	switch dir {
	case UP:
		return up(i, j)
	case DOWN:
		return down(i, j)
	case LEFT:
		return left(i, j)
	case RIGHT:
		return right(i, j)
	}
	return -1, -1, ""
}

func getNextMirror(i, j int, dir string, mirror rune) (int, int, string) {
	if mirror == '/' {
		switch dir {
		case UP:
			return right(i, j)
		case DOWN:
			return left(i, j)
		case LEFT:
			return down(i, j)
		case RIGHT:
			return up(i, j)
		}
		return -1, -1, ""
	}

	// \
	switch dir {
	case UP:
		return left(i, j)
	case DOWN:
		return right(i, j)
	case LEFT:
		return up(i, j)
	case RIGHT:
		return down(i, j)
	}
	return -1, -1, ""

}

func left(i, j int) (int, int, string) {
	return i, j - 1, LEFT
}

func right(i, j int) (int, int, string) {
	return i, j + 1, RIGHT
}

func up(i, j int) (int, int, string) {
	return i - 1, j, UP
}

func down(i, j int) (int, int, string) {
	return i + 1, j, DOWN
}

func getKey(i, j int, dir string) string {
	return fmt.Sprintf("%d+%d:%s", i, j, dir)
}
