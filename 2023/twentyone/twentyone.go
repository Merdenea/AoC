package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type direction struct {
	i, j  int
	count int
}

type Queue struct {
	l []direction
}

func NewQueue() Queue {
	return Queue{l: make([]direction, 0)}
}

func (q *Queue) Pop() direction {
	v := q.l[0]
	q.l = q.l[1:]
	return v
}

func (q *Queue) Push(ps ...direction) {
	for _, p := range ps {
		q.l = append(q.l, p)
	}
}

func (q *Queue) IsEmpty() bool {
	return len(q.l) == 0
}

var (
	UP    = direction{-1, 0, 0}
	DOWN  = direction{1, 0, 0}
	LEFT  = direction{0, -1, 0}
	RIGHT = direction{0, 1, 0}
	NONE  = direction{0, 0, 0}
)

func (d direction) apply(new direction) direction {
	dd := direction{d.i + new.i, d.j + new.j, d.count - 1}
	return dd
}

func main() {
	file, _ := os.Open("./twentyone/in")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	m := make([][]rune, 0)
	for scanner.Scan() {
		l := scanner.Text()
		m = append(m, []rune(l))
	}

	si, sj := findStart(m, 'S')
	now := time.Now()
	fmt.Printf("Part One %d in [%s]\n", getCount(m, si, sj, 64), time.Now().Sub(now))
	now = time.Now()
	fmt.Printf("Part Two %d in [%s]\n", getInfiniteGridCount(m, si, sj, 26501365), time.Now().Sub(now))
}

func key(i, j int) string {
	return fmt.Sprintf("%d+%d", i, j)
}

func getInfiniteGridCount(m [][]rune, si, sj, steps int) int {
	size := len(m)
	gridW := steps/size - 1

	oddGrids := gridW * gridW
	evenGrids := (gridW + 1) * (gridW + 1)

	oddPoints := getCount(m, si, sj, 2*size+1)
	evenPoints := getCount(m, si, sj, 2*size)

	topCorner := getCount(m, size-1, sj, size-1)
	rightCorner := getCount(m, si, 0, size-1)
	bottomCorner := getCount(m, 0, sj, size-1)
	leftCorner := getCount(m, si, size-1, size-1)

	smallTopR := getCount(m, size-1, 0, size/2-1)
	smallTopL := getCount(m, size-1, size-1, size/2-1)
	smallBottomR := getCount(m, 0, 0, size/2-1)
	smallBottomL := getCount(m, 0, size-1, size/2-1)

	largeTopR := getCount(m, size-1, 0, 3*size/2-1)
	largeTopL := getCount(m, size-1, size-1, 3*size/2-1)
	largeBottomR := getCount(m, 0, 0, 3*size/2-1)
	largeBottomL := getCount(m, 0, size-1, 3*size/2-1)

	return oddPoints*oddGrids +
		evenPoints*evenGrids +
		topCorner + rightCorner + leftCorner + bottomCorner +
		(gridW+1)*(smallTopL+smallTopR+smallBottomR+smallBottomL) +
		gridW*(largeTopR+largeTopL+largeBottomL+largeBottomR)
}

func getCount(m [][]rune, si, sj, steps int) int {
	q := NewQueue()
	d := direction{si, sj, steps}
	size := len(m) // assumes square grid
	q.Push(d)

	visited := make(map[string]bool)
	finalTile := make(map[string]bool)

	for !q.IsEmpty() {
		current := q.Pop()

		if current.count%2 == 0 {
			finalTile[key(current.i, current.j)] = true
		}

		if current.count == 0 {
			continue
		}

		for _, dir := range []direction{UP, DOWN, RIGHT, LEFT} {
			newD := current.apply(dir)
			if newD.i < 0 || newD.j < 0 || newD.j >= size || newD.i >= size || visited[key(newD.i, newD.j)] {
				continue
			}

			if m[newD.i][newD.j] == '#' {
				continue
			}

			//if m[((newD.i%size)+size)%size][((newD.j%size)+size)%size] == '#' {
			//	continue
			//}
			visited[key(newD.i, newD.j)] = true
			q.Push(newD)
		}
	}

	return len(finalTile)
}

func findStart(m [][]rune, start rune) (int, int) {
	// appears to be the centre of the grid
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			if m[i][j] == start {
				return i, j
			}
		}
	}

	return -1, -1
}
