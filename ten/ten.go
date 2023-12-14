package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var (
	visited   map[string]bool
	dots      map[string]bool
	inTheLoop map[string]bool

	minI = 0
	minJ = 0
	maxI = math.MaxInt64
	maxJ = math.MaxInt64
)

const (
	UP    = "UP"
	DOWN  = "DOWN"
	LEFT  = "LEFT"
	RIGHT = "RIGHT"
)

func main() {
	file, _ := os.Open("./ten/in")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	mmap := make([][]rune, 0)

	for scanner.Scan() {
		l := scanner.Text()
		rr := make([]rune, 0)
		for _, r := range []rune(l) {
			rr = append(rr, r)
		}
		mmap = append(mmap, rr)
	}

	sI, sJ := findStart(mmap)
	key := fmt.Sprintf("%d+%d", sI, sJ)

	visited = make(map[string]bool)
	inTheLoop = make(map[string]bool)
	visited[key] = true

	sVal, a, b, direction := getStartingVal(mmap, sI, sJ)
	dfs(mmap, a, b, direction)

	for i := 0; i < len(mmap); i++ {
		for j := 0; j < len(mmap[i]); j++ {
			k := fmt.Sprintf("%d+%d", i, j)
			if strings.Contains("FLJ7S|-", string(mmap[i][j])) && !inTheLoop[k] {
				mmap[i][j] = '.'
			}
		}
	}

	mmap[sI][sJ] = sVal

	mm := make([][]rune, 0)

	// Do I need this?
	for i := 0; i < len(mmap); i++ {
		mm = append(mm, getLine(mmap[i]))
		mm = append(mm, getLine(mmap[i]))
		for j := 0; j < len(mmap[i]); j++ {
			switch mmap[i][j] {
			case 'F':
				mm[len(mm)-2][j] = 'F'
				mm[len(mm)-1][j] = '|'
				break
			case '7':
				mm[len(mm)-2][j] = '7'
				mm[len(mm)-1][j] = '|'
				break
			case '|':
				mm[len(mm)-2][j] = '|'
				break
			default:
				mm[len(mm)-1][j] = 'x'
			}
		}
	}

	fm := make([][]rune, 0)

	for i := 0; i < len(mm); i++ {
		fm = append(fm, getLine(mm[i]))
		fm = append(fm, getLine(mm[i]))
		for j := 0; j < len(mm[i]); j++ {
			switch mm[i][j] {
			case 'J':
				fm[len(fm)-2][j] = '|'
				fm[len(fm)-1][j] = 'J'
				break
			case 'L':
				fm[len(fm)-2][j] = '|'
				fm[len(fm)-1][j] = 'L'
				break
			case '|':
				fm[len(fm)-1][j] = '|'
				break
			default:
				fm[len(fm)-2][j] = 'x'
			}
		}
	}

	oneMap := make([][]rune, 0)
	for i := 0; i < len(fm); i++ {
		line := getEmpty(3 * len(fm[i]))
		for j := 0; j < 3*len(fm[i]); j += 3 {
			line[j+1] = fm[i][j/3]
			switch line[j+1] {
			case 'F':
				line[j] = 'x'
				line[j+2] = '-'
				break
			case '7':
				line[j] = '-'
				line[j+2] = 'x'
				break
			case 'J':
				line[j] = '-'
				line[j+2] = 'x'
				break
			case 'L':
				line[j] = 'x'
				line[j+2] = '-'
				break
			case '-':
				line[j] = '-'
				line[j+2] = '-'
			default:
				line[j] = 'x'
				line[j+2] = 'x'
			}
		}
		oneMap = append(oneMap, line)
	}

	visited = make(map[string]bool)
	dots = make(map[string]bool)

	floodFill(oneMap, 0, 0)
	print(oneMap)

}

// Not fully tested
func getStartingVal(mmap [][]rune, i int, j int) (rune, int, int, string) {
	if (j+1 < len(mmap[i]) && j-1 >= 0) && strings.Contains("J7-", string(mmap[i][j+1])) && strings.Contains("LF-", string(mmap[i][j-1])) {
		return '-', i, j + 1, RIGHT
	}
	if (i+1 < len(mmap) && i-1 >= 0) && strings.Contains("F7|", string(mmap[i-1][j])) && strings.Contains("LJ|", string(mmap[i+1][j])) {
		return '|', i + 1, j, UP
	}
	if (j+1 < len(mmap[i]) && i+1 < len(mmap)) && strings.Contains("|JL", string(mmap[i+1][j])) && strings.Contains("-7J", string(mmap[i][j+1])) {
		return 'F', i, j + 1, RIGHT
	}
	if (i+1 < len(mmap) && j-1 >= 0) && strings.Contains("-FL", string(mmap[i][j-1])) && strings.Contains("|LJ", string(mmap[i+1][j])) {
		return '7', i + 1, j, DOWN
	}
	if (j+1 < len(mmap[i]) && i-1 >= 0) && strings.Contains("J-7", string(mmap[i][j+1])) && strings.Contains("|7F", string(mmap[i-1][j])) {
		return 'L', i, j + 1, RIGHT
	}
	if (i-1 >= 0 && j-1 >= 0) && strings.Contains("-FL", string(mmap[i][j-1])) && strings.Contains("|F7", string(mmap[i-1][j])) {
		return 'J', i, j - 1, LEFT
	}

	return ' ', i, j, ""
}

func getEmpty(i int) []rune {
	s := make([]rune, 0)
	for j := 0; j < i; j++ {
		s = append(s, ' ')
	}
	return s
}

func getLine(l []rune) []rune {
	res := make([]rune, 0)
	for i := 0; i < len(l); i++ {
		res = append(res, l[i])
	}
	return res
}

func print(mmap [][]rune) {
	count := 0

	for i := 0; i < len(mmap); i++ {
		for j := 0; j < len(mmap[i]); j++ {
			key := fmt.Sprintf("%d+%d", i, j)
			if mmap[i][j] == '.' && !dots[key] {
				count++
			}
		}
	}
	fmt.Println(count)

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

func getNext(i, j int, r rune, direction string) (int, int, string) {
	switch r {
	case '-':
		if direction == LEFT {
			return left(i, j)
		}
		return right(i, j)
	case '|':
		if direction == UP {
			return up(i, j)
		}
		return down(i, j)
	case '7':
		if direction == RIGHT {
			return down(i, j)
		}
		return left(i, j)
	case 'J':
		if direction == DOWN {
			return left(i, j)
		}
		return up(i, j)
	case 'L':
		if direction == LEFT {
			return up(i, j)
		}
		return right(i, j)
	case 'F':
		if direction == UP {
			return right(i, j)
		}
		return down(i, j)
	}
	return -1, -1, ""
}

func dfs(mmap [][]rune, i, j int, dir string) int {
	if i < 0 || i >= len(mmap) {
		return 0
	}

	if j < 0 || j >= len(mmap[i]) {
		return 0
	}

	if mmap[i][j] == '.' {
		return 0
	}
	key := fmt.Sprintf("%d+%d", i, j)

	if strings.Contains("FLJ7S|-", string(mmap[i][j])) {
		inTheLoop[key] = true
	}

	if visited[key] {
		return 0
	}
	visited[key] = true
	minI = min(i, minI)
	minJ = min(j, minJ)

	maxI = max(i, maxI)
	maxJ = max(j, maxJ)

	nextI, nextJ, nextD := getNext(i, j, mmap[i][j], dir)

	return 1 + dfs(mmap, nextI, nextJ, nextD)
}

func floodFill(mmap [][]rune, i, j int) int {
	if i < 0 || i >= len(mmap) {
		return 0
	}

	if j < 0 || j >= len(mmap[i]) {
		return 0
	}
	key := fmt.Sprintf("%d+%d", i, j)
	if visited[key] {
		return 0
	}

	if strings.Contains("FLJ7S|-", string(mmap[i][j])) {
		visited[key] = true
		return 0
	}

	visited[key] = true
	if mmap[i][j] == '.' {
		dots[key] = true
	}

	u1, u2, _ := up(i, j)
	floodFill(mmap, u1, u2)

	u1, u2, _ = down(i, j)
	floodFill(mmap, u1, u2)

	u1, u2, _ = left(i, j)
	floodFill(mmap, u1, u2)

	u1, u2, _ = right(i, j)
	floodFill(mmap, u1, u2)
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findStart(mmap [][]rune) (int, int) {
	for i := 0; i < len(mmap); i++ {
		for j := 0; j < len(mmap[i]); j++ {
			if mmap[i][j] == 'S' {
				return i, j
			}
		}
	}
	return -1, -1
}
