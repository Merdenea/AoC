package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"

	"github.com/addrummond/heap"
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
	two := partTwo()
	fmt.Printf("Part II in [%s]: %d\n", time.Since(now).String(), two)

}

type node struct {
	score, x, y, dir_x, dir_y int
	path                      string
}

var seen = make(map[[4]int]int)

var (
	paths = make([]string, 0)
)

func (t1 *node) Cmp(t2 *node) int {
	return t1.score - t2.score
}

func partOne(m [][]rune) int64 {
	si, sj := findStart(m)

	var pq heap.Heap[node, heap.Min]
	heap.PushOrderable(&pq, node{0, si, sj, 0, 1, fmt.Sprintf("%d-%d", si, sj)})

	seen[[4]int{si, sj, 0, 1}]++

	minPath := math.MaxInt64

	for heap.Len(&pq) > 0 {
		current, _ := heap.PopOrderable(&pq)
		seen[[4]int{current.x, current.y, current.dir_x, current.dir_y}]++

		if m[current.x][current.y] == 'E' {
			if current.score > minPath {
				continue
			}
			minPath = current.score
			paths = append(paths, current.path)
		}

		for _, next := range []node{
			node{current.score + 1, current.x + current.dir_x, current.y + current.dir_y, current.dir_x, current.dir_y, current.path + fmt.Sprintf(",%d-%d", current.x+current.dir_x, current.y+current.dir_y)},
			node{current.score + 1000, current.x, current.y, current.dir_y, -current.dir_x, current.path}, // clockwise
			node{current.score + 1000, current.x, current.y, -current.dir_y, current.dir_x, current.path}, // counter clockwise
		} {
			if m[next.x][next.y] == '#' {
				continue
			}
			if seen[[4]int{next.x, next.y, next.dir_x, next.dir_y}] > 0 {
				continue
			}
			heap.PushOrderable(&pq, next)
		}

	}
	return int64(minPath)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findStart(m [][]rune) (int, int) {
	for i, l := range m {
		for j, r := range l {
			if r == 'S' {
				return i, j
			}
		}
	}
	return -1, -1
}

func partTwo() int64 {
	best := make(map[string]int)
	for _, st := range paths {
		sp := strings.Split(st, ",")
		for _, s := range sp {
			best[s]++
		}
	}
	return int64(len(best))
}
