package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"

	"github.com/gammazero/deque"
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

func partOne(m [][]rune) int64 {
	si, sj := -1, -1
	ei, ej := -1, -1
	for i, l := range m {
		for j, r := range l {
			if r == 'S' {
				si, sj = i, j
			}
			if r == 'E' {
				ei, ej = i, j
			}
		}
	}

	// get cost from path position to end
	bfsNoCheat(ei, ej, si, sj, m)
	return int64(bfs(2, si, sj, ei, ej, m))
}

type node struct {
	i, j, time int
	cheated    bool
}

var timeCost = make(map[[2]int]int)

func bfsNoCheat(si, sj, ei, ej int, m [][]rune) int {
	q := make([]node, 0)
	q = append(q, node{i: si, j: sj, time: 0})
	seen := make(map[[2]int]int)

	for len(q) > 0 {
		current := q[0]
		q = q[1:]
		if current.i == ei && current.j == ej {
			timeCost[[2]int{current.i, current.j}] = current.time
			return current.time
		}

		if seen[[2]int{current.i, current.j}] > 0 {
			timeCost[[2]int{current.i, current.j}] = min(current.time, timeCost[[2]int{current.i, current.j}])
			continue
		}
		seen[[2]int{current.i, current.j}]++
		timeCost[[2]int{current.i, current.j}] = current.time

		for _, dir := range [][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}} {
			ni, nj := current.i+dir[0], current.j+dir[1]

			if m[ni][nj] == '#' {
				continue
			}
			q = append(q, node{i: ni, j: nj, time: current.time + 1})
		}
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func bfs(cheatLen, si, sj, ei, ej int, m [][]rune) int {
	// q := make([]node, 0)
	// no actual improvement..
	var q deque.Deque[node]

	q.PushBack(node{si, sj, 0, false})

	noCheatTime := -1
	cheatTimes := make([]int, 0)

	var cheats map[[2]int][][3]int
	cheats = getAllCheatsX2(cheatLen, m)

	seen := make(map[[2]int]int)
	for q.Len() > 0 {
		current := q.PopFront()

		if current.i == ei && current.j == ej {
			if !current.cheated {
				noCheatTime = current.time
			} else {
				cheatTimes = append(cheatTimes, current.time)
			}
			continue
		}

		if t, ok := timeCost[[2]int{current.i, current.j}]; ok && current.cheated {
			cheatTimes = append(cheatTimes, current.time+t)
			continue
		}

		if seen[[2]int{current.i, current.j}] > 0 {
			continue
		}
		seen[[2]int{current.i, current.j}]++

		for _, cheat := range cheats[[2]int{current.i, current.j}] {
			ni, nj, nt := cheat[0], cheat[1], cheat[2]
			q.PushBack(node{ni, nj, current.time + nt, true})
		}

		for _, dir := range [][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}} {
			ni, nj := current.i+dir[0], current.j+dir[1]

			if m[ni][nj] == '#' {
				continue
			}

			q.PushBack(node{ni, nj, current.time + 1, current.cheated})
		}

	}

	count := 0
	for _, t := range cheatTimes {
		if noCheatTime-t >= 100 {
			count++
		}
	}

	return count
}

// takes 2 seconds..
func getAllCheats(cheatLen int, m [][]rune) map[[2]int][][3]int {
	cheats := make(map[[4]int]int)

	// save the min cost cheat if it has the same start finish

	for k, _ := range timeCost {
		si, sj := k[0], k[1]

		// si, sj, time
		q := make([][3]int, 0)
		q = append(q, [3]int{si, sj, 0})
		seen := make(map[[2]int]int)

		for len(q) > 0 {
			cur := q[0]
			q = q[1:]

			ci, cj, ct := cur[0], cur[1], cur[2]
			if ct > cheatLen {
				continue
			}

			if seen[[2]int{ci, cj}] > 0 {
				continue
			}
			seen[[2]int{ci, cj}]++

			// check cw(curretn walls) > 0 to only add cheats
			if ct >= 2 && m[ci][cj] != '#' {
				if v, ok := cheats[[4]int{si, sj, ci, cj}]; ok {
					cheats[[4]int{si, sj, ci, cj}] = min(v, ct)
				} else {
					cheats[[4]int{si, sj, ci, cj}] = ct
				}
			}

			for _, dir := range [][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}} {
				ni, nj := ci+dir[0], cj+dir[1]

				if ni >= 0 && nj >= 0 && ni < len(m) && nj < len(m[0]) {
					q = append(q, [3]int{ni, nj, ct + 1})
				}
			}
		}
	}

	// si, sj to ei, ej, time
	res := make(map[[2]int][][3]int)

	for k, v := range cheats {
		res[[2]int{k[0], k[1]}] = append(res[[2]int{k[0], k[1]}], [3]int{k[2], k[3], v})
	}

	return res
}

// takes 1.9s whomp whomp
func getAllCheatsX2(cheatLen int, m [][]rune) map[[2]int][][3]int {
	cheats := make(map[[4]int]int)

	for i, l := range m {
		for j, r := range l {
			if r == '#' {
				continue
			}
			for k, _ := range timeCost {
				si, sj := k[0], k[1]
				dist := absDist(si, sj, i, j)
				if dist >= 2 && dist <= cheatLen {
					if v, ok := cheats[[4]int{si, sj, i, j}]; ok {
						cheats[[4]int{si, sj, i, j}] = min(v, dist)
					} else {
						cheats[[4]int{si, sj, i, j}] = dist
					}
				}
			}
		}
	}

	// si, sj to ei, ej, time
	res := make(map[[2]int][][3]int)

	for k, v := range cheats {
		res[[2]int{k[0], k[1]}] = append(res[[2]int{k[0], k[1]}], [3]int{k[2], k[3], v})
	}

	return res
}

func absDist(si, sj, ei, ej int) int {
	return int(math.Abs(float64(si-ei))) + int(math.Abs(float64(sj-ej)))
}

func partTwo(m [][]rune) int64 {
	si, sj := -1, -1
	ei, ej := -1, -1
	for i, l := range m {
		for j, r := range l {
			if r == 'S' {
				si, sj = i, j
			}
			if r == 'E' {
				ei, ej = i, j
			}
		}
	}

	// todo - takes 2 seconds fix this
	return int64(bfs(20, si, sj, ei, ej, m))
}
