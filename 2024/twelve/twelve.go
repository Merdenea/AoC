package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	two := partTwo(m)
	fmt.Printf("Part II in [%s]: %d\n", time.Since(now).String(), two)

}

var seen = make(map[string]int)

func partOne(m [][]rune) int64 {
	total := 0
	for i, l := range m {
		for j, _ := range l {
			total += findAreaAndPer(false, i, j, m)
		}
	}

	return int64(total)
}

func key(i, j int) string {
	return fmt.Sprintf("%s-%s", i, j)
}

func findAreaAndPer(partTwo bool, i, j int, m [][]rune) int {
	if seen[key(i, j)] > 0 {
		return 0
	}

	q := make([][2]int, 0)
	q = append(q, [2]int{i, j})

	area := make([][2]int, 0)

	for len(q) > 0 {
		current := q[0]
		q = q[1:]

		if seen[key(current[0], current[1])] > 0 {
			continue
		}

		area = append(area, [2]int{current[0], current[1]})
		seen[key(current[0], current[1])]++

		for _, dir := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
			ni, nj := current[0]+dir[0], current[1]+dir[1]

			if ni < 0 || nj < 0 || ni >= len(m) || nj >= len(m[0]) {
				continue
			}
			if m[ni][nj] == m[i][j] {
				q = append(q, [2]int{ni, nj})
			}
		}
	}

	per, sides := findPer(area, m)
	if partTwo {
		return sides * len(area)
	}
	return per * len(area)
}

func findPer(area [][2]int, m [][]rune) (int, int) {
	i, j := area[0][0], area[0][1]
	per := 0
	for _, p := range area {
		sides := 0
		for _, dir := range [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
			ni, nj := p[0]+dir[0], p[1]+dir[1]

			if ni >= 0 && nj >= 0 && ni < len(m) && nj < len(m[0]) {
				if m[ni][nj] != m[i][j] {
					sides++
				}
			} else {
				if ni < 0 || nj < 0 {
					sides++
				}

				if ni >= len(m) || nj >= len(m[0]) {
					sides++
				}
			}
		}
		per += sides

	}
	sCount := findSides(area, m)
	return per, sCount
}

func findSides(area [][2]int, m [][]rune) int {
	if len(area) == 1 {
		return 4
	}

	areaMap := make(map[[2]int]int)
	for _, p := range area {
		areaMap[p]++
	}

	crossHorizontal := make([][2]int, 0)
	for i, l := range m {
		in := false
		for j, _ := range l {
			if areaMap[[2]int{i, j}] > 0 {
				if !in {
					in = true
					crossHorizontal = append(crossHorizontal, [2]int{i, j})
				}
				if j == len(m[0])-1 {
					in = false
					d := 1
					if j == 0 {
						d = -1
					}
					crossHorizontal = append(crossHorizontal, [2]int{i, j + d})
				}
			} else {
				if in {
					in = false
					crossHorizontal = append(crossHorizontal, [2]int{i, j})
				}
			}
		}
	}

	jFreq := make(map[int][]int)

	for _, p := range crossHorizontal {
		jFreq[p[1]] = append(jFreq[p[1]], p[0])
	}

	hCross := 0

	for j, iPos := range jFreq {
		if len(iPos) == 1 {
			hCross++
		} else {
			sort.Ints(iPos)
			hCross++
			for index := 1; index < len(iPos); index++ {
				if iPos[index]-1 != iPos[index-1] {
					hCross++
				} else if areaMap[[2]int{iPos[index], j}] != areaMap[[2]int{iPos[index-1], j}] {
					hCross++
				}
			}
		}
	}

	crossVertical := make([][2]int, 0)
	for j := 0; j < len(m[0]); j++ {
		in := false
		for i := 0; i < len(m); i++ {
			// r := m[i][j]
			if areaMap[[2]int{i, j}] > 0 {
				if !in {
					in = true
					crossVertical = append(crossVertical, [2]int{i, j})
				}
				if i == len(m)-1 {
					in = false
					d := 1
					if i == 0 {
						d = -1
					}
					crossVertical = append(crossVertical, [2]int{i + d, j})
				}
			} else {
				if in {
					in = false
					crossVertical = append(crossVertical, [2]int{i, j})
				}
			}
		}
	}

	iFreq := make(map[int][]int)

	for _, p := range crossVertical {
		iFreq[p[0]] = append(iFreq[p[0]], p[1])
	}

	vCross := 0

	for i, jPos := range iFreq {
		if len(jPos) == 1 {
			vCross++
		} else {
			sort.Ints(jPos)
			vCross++
			for index := 1; index < len(jPos); index++ {
				if jPos[index]-1 != jPos[index-1] {
					vCross++
				} else if areaMap[[2]int{i, jPos[index]}] != areaMap[[2]int{i, jPos[index-1]}] {
					hCross++
				}
			}
		}
	}

	return hCross + vCross
}

func partTwo(m [][]rune) int64 {
	seen = make(map[string]int)
	total := 0
	for i, l := range m {
		for j, _ := range l {
			total += findAreaAndPer(true, i, j, m)
		}
	}
	return int64(total)
}
