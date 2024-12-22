package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gammazero/deque"
)

var numKeypad = [][]rune{
	{'7', '8', '9'},
	{'4', '5', '6'},
	{'1', '2', '3'},
	{' ', '0', 'A'},
}

var dirKeypad = [][]rune{
	{' ', '^', 'A'},
	{'<', 'v', '>'},
}

var (
	numRobot_i, numRobot_j = 3, 2

	robotOne_i, robotOne_j = 0, 2
)

func main() {
	file, _ := os.Open(fmt.Sprintf("./%s", os.Args[1]))
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	codes := make([]string, 0)
	for scanner.Scan() {
		codes = append(codes, scanner.Text())
	}

	now := time.Now()
	one := partOne(codes)
	fmt.Printf("Part I in [%s]: %d\n", time.Since(now).String(), one)

	now = time.Now()
	two := partTwo(codes)
	fmt.Printf("Part II in [%s]: %d\n", time.Since(now).String(), two)

}

func partOne(codes []string) int64 {
	total := int64(0)

	for _, c := range codes {
		numericMoves = make(map[string]int)
		moveNumericRobot(c, "", 3, 2)
		seconds := make(map[string]int)
		for k, _ := range numericMoves {
			seconds[moveDirectionalRobot(k)]++
		}

		third := ""
		for k, _ := range seconds {
			newThird := moveDirectionalRobot(k)
			if third == "" {
				third = newThird
				continue
			}
			if len(newThird) < len(third) {
				third = newThird
			}
		}
		val, _ := strconv.ParseInt(c[:len(c)-1], 10, 64)
		total += int64(len(third)) * val
	}

	return total
}

func moveDirectionalRobot(moves string) string {
	res := strings.Builder{}

	for _, m := range moves {
		ni, nj := findNextDir(rune(m))

		diff_i, diff_j := ni-robotOne_i, nj-robotOne_j

		var upDown, leftRight string
		if diff_i < 0 {
			upDown = getMoves('^', diff_i*-1)
		} else if diff_i > 0 {
			upDown = getMoves('v', diff_i)
		}

		if diff_j < 0 {
			leftRight = getMoves('<', diff_j*-1)
		} else if diff_j > 0 {
			leftRight = getMoves('>', diff_j)
		}

		cs := fmt.Sprintf("%s%sA", upDown, leftRight)

		res.WriteString(cs)
		robotOne_i, robotOne_j = ni, nj
	}

	return res.String()
}

var numericMoves = make(map[string]int)

func validateNumericPath(moves string) bool {
	si, sj := 3, 2
	xi, xj := 3, 0

	for _, m := range moves {
		switch m {
		case '^':
			si -= 1
		case 'v':
			si += 1
		case '<':
			sj -= 1
		case '>':
			sj += 1
		}
		if si == xi && sj == xj {
			return false
		}
	}
	return true
}

func moveNumericRobot(code, current string, si, sj int) {
	if len(code) == 0 {
		if validateNumericPath(current) {
			numericMoves[current]++
		}
		return
	}

	c := code[0]
	ni, nj := findNext(rune(c))

	diff_i, diff_j := ni-si, nj-sj
	var upDown, leftRight string
	if diff_i < 0 {
		upDown = getMoves('^', diff_i*-1)
	} else if diff_i > 0 {
		upDown = getMoves('v', diff_i)
	}

	if diff_j < 0 {
		leftRight = getMoves('<', diff_j*-1)
	} else if diff_j > 0 {
		leftRight = getMoves('>', diff_j)
	}

	moveNumericRobot(code[1:], fmt.Sprintf("%s%s%sA", current, upDown, leftRight), ni, nj)
	moveNumericRobot(code[1:], fmt.Sprintf("%s%s%sA", current, leftRight, upDown), ni, nj)
}

var movesM = make(map[[2]int]string)

func getMoves(r rune, times int) string {
	if v, ok := movesM[[2]int{int(r), times}]; ok {
		return v
	}
	sb := strings.Builder{}
	for i := 0; i < times; i++ {
		sb.WriteRune(r)
	}
	movesM[[2]int{int(r), times}] = sb.String()
	return sb.String()
}

func findNextDir(c rune) (int, int) {
	for i, l := range dirKeypad {
		for j, r := range l {
			if r == c {
				return i, j
			}
		}
	}
	return -1, -1
}

func findNext(c rune) (int, int) {
	for i, l := range numKeypad {
		for j, r := range l {
			if r == c {
				return i, j
			}
		}
	}
	return -1, -1
}

var (
	dirSeq = getPossibleSeqs(dirKeypad)
	numSeq = getPossibleSeqs(numKeypad)
)

func partTwo(codes []string) int64 {
	total := int64(0)

	for _, c := range codes {
		numericMoves = make(map[string]int)
		moveNumericRobot(c, "", 3, 2)

		minVal := 0

		for k, _ := range numericMoves {
			currentVal := moveRobotCount(k, 25)
			if minVal == 0 {
				minVal = currentVal
				continue
			}

			if currentVal < minVal {
				minVal = currentVal
			}
		}

		val, _ := strconv.ParseInt(c[:len(c)-1], 10, 64)

		total += int64(minVal) * val
	}

	return total
}

var memo = make(map[string]int)

// can solve PART I with level = 2
func moveRobotCount(moves string, level int) int {
	if v, ok := memo[fmt.Sprintf("%s-%d", moves, level)]; ok {
		return v
	}

	if level == 1 {
		res := 0
		for i := 0; i < len(moves); i++ {
			a, b := ' ', ' '
			if i == 0 {
				a, b = 'A', rune(moves[i])
			} else {
				a, b = rune(moves[i-1]), rune(moves[i])
			}
			res += len(dirSeq[[2]rune{a, b}][0])
		}
		return res
	}

	totalMin := 0
	for i := 0; i < len(moves); i++ {
		a, b := ' ', ' '
		if i == 0 {
			a, b = 'A', rune(moves[i])
		} else {
			a, b = rune(moves[i-1]), rune(moves[i])
		}

		currentMin := 0
		for _, s := range dirSeq[[2]rune{a, b}] {
			cs := moveRobotCount(s, level-1)
			if currentMin == 0 {
				currentMin = cs
				continue
			}
			if cs < currentMin {
				currentMin = cs
			}
		}
		totalMin += currentMin
	}

	memo[fmt.Sprintf("%s-%d", moves, level)] = totalMin
	return totalMin
}

var seqs = make(map[[4]int]string)

// from si, sj to ei, ej
// get all valid (si, sj) -> (ei, ej)

type PosSeq struct {
	i, j  int
	moves string
}

func getPossibleSeqs(keyPad [][]rune) map[[2]rune][]string {
	symToPos := make(map[rune][2]int)
	for i, l := range keyPad {
		for j, r := range l {
			if r != ' ' {
				symToPos[r] = [2]int{i, j}
			}
		}
	}

	// symbol to symbol
	seqs := make(map[[2]rune][]string)

	for k1, _ := range symToPos {
		for k2, _ := range symToPos {
			if k1 == k2 {
				seqs[[2]rune{k1, k2}] = []string{"A"}
				continue
			}

			possible := make([]string, 0)

			var q deque.Deque[PosSeq]
			i, j := symToPos[k1][0], symToPos[k1][1]

			q.PushBack(PosSeq{i, j, ""})

			currentMin := math.MaxInt64
			foundMin := false
			for q.Len() > 0 && !foundMin {
				cur := q.PopFront()
				i, j, m := cur.i, cur.j, cur.moves

				for _, dir := range []PosSeq{{0, 1, ">"}, {1, 0, "v"}, {-1, 0, "^"}, {0, -1, "<"}} {
					ni, nj, nm := i+dir.i, j+dir.j, dir.moves

					if ni < 0 || nj < 0 || ni >= len(keyPad) || nj >= len(keyPad[0]) {
						continue
					}
					if keyPad[ni][nj] == ' ' {
						continue
					}
					if keyPad[ni][nj] == k2 {
						if currentMin < len(m)+1 {
							foundMin = true
							break
						}
						currentMin = len(m) + 1
						possible = append(possible, m+nm+"A")
						continue
					}
					q.PushBack(PosSeq{ni, nj, m + nm})
				}
			}

			seqs[[2]rune{k1, k2}] = possible
		}
	}
	return seqs
}
