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
	dd := direction{d.i + new.i, d.j + new.j, d.count + 1}
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
	fmt.Printf("Part One %d in [%s]\n", getPlotsCount(m, direction{si, sj, 0}, 64), time.Now().Sub(now))

	// Part two - Nope - TBDs
}

func key(i, j int) string {
	return fmt.Sprintf("%d+%d", i, j)
}

func getPlotsCount(m [][]rune, d direction, steps int) int {
	q := NewQueue()

	size := len(m) // assumes square grid
	q.Push(d)

	count := 0

	visited := make(map[string]bool)

	for !q.IsEmpty() {
		current := q.Pop()

		if visited[key(current.i, current.j)] {
			continue
		}
		visited[key(current.i, current.j)] = true

		if current.count == steps {
			count++
			continue
		}
		if current.count%2 == 0 {
			count++
		}

		for _, dir := range []direction{UP, DOWN, RIGHT, LEFT} {
			newD := current.apply(dir)
			//
			//if newD.i < 0 || newD.j < 0 || newD.j >= size || newD.i > size {
			//	continue
			//}

			if m[((newD.i%size)+size)%size][((newD.j%size)+size)%size] == '#' {
				continue
			}
			q.Push(newD)
		}
	}

	return count
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
