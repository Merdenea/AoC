package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type brick struct {
	x1, y1, z1 int64
	x2, y2, z2 int64

	ID string
}

func main() {
	file, _ := os.Open("./twentytwo/in")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	m := make([]*brick, 0)

	c := 0
	for scanner.Scan() {
		l := scanner.Text()
		sp := strings.Split(l, "~")
		sp1 := strings.Split(sp[0], ",")
		sp2 := strings.Split(sp[1], ",")
		x1, _ := strconv.ParseInt(sp1[0], 10, 64)
		y1, _ := strconv.ParseInt(sp1[1], 10, 64)
		z1, _ := strconv.ParseInt(sp1[2], 10, 64)

		x2, _ := strconv.ParseInt(sp2[0], 10, 64)
		y2, _ := strconv.ParseInt(sp2[1], 10, 64)
		z2, _ := strconv.ParseInt(sp2[2], 10, 64)

		m = append(m, &brick{x1, y1, z1, x2, y2, z2, fmt.Sprintf("b%d", c)})
		c++
	}

	now := time.Now()
	fmt.Printf("Part One %d in [%s]\n", getTotalDisintegrated(m, false), time.Now().Sub(now))
	now = time.Now()
	fmt.Printf("Part Two %d in [%s]\n", getTotalDisintegrated(m, true), time.Now().Sub(now))

}

func getTotalDisintegrated(m []*brick, part2 bool) int {
	sort.Slice(m, func(i, j int) bool {
		return m[i].z1 < m[j].z1
	})

	for i, b := range m {
		maxZ := int64(1)
		for _, next := range m[:i] {
			if xyIntersect(*b, *next) {
				maxZ = max(maxZ, next.z2+1)
			}
		}
		b.z2 -= b.z1 - maxZ
		b.z1 = maxZ
	}

	sort.Slice(m, func(i, j int) bool {
		return m[i].z1 < m[j].z1
	})

	//for _, b := range m {
	//	fmt.Println(*b)
	//}

	supports := make(map[string][]string)
	supportedBy := make(map[string][]string)
	for _, b := range m {
		supports[b.ID] = make([]string, 0)
		supportedBy[b.ID] = make([]string, 0)
	}

	for _, b := range m {
		bricksAbove := getBricksOnLevel(m, int64(b.z2)+1)

		for _, above := range bricksAbove {
			if xyIntersect(*b, above) {
				//fmt.Printf("[%s] supports [%s]\n", b.ID, above.ID)
				supports[b.ID] = append(supports[b.ID], above.ID)
				supportedBy[above.ID] = append(supportedBy[above.ID], b.ID)
			}
		}
	}

	//fmt.Println(supports)
	//fmt.Println(supportedBy)

	if !part2 {
		total := 0
		for _, b := range m {
			check := true
			for _, id := range supports[b.ID] {
				if len(supportedBy[id]) <= 1 {
					check = false
				}
			}
			if check {
				total++
			}
		}
		return total
	}

	total := 0
	for _, b := range m {
		destroyed := make(map[string]bool)
		destroyed[b.ID] = true
		supported := supports[b.ID]
		q := NewQueue()
		for _, a := range supported {
			if len(supportedBy[a]) == 1 {
				q.Push(a)
				destroyed[a] = true
			}
		}

		for !q.IsEmpty() {
			cb := q.Pop()

			for _, a := range supports[cb] {
				if destroyed[a] {
					continue
				}
				aSupportedBy := supportedBy[a]
				allDestroyed := true
				for _, as := range aSupportedBy {
					if !destroyed[as] {
						allDestroyed = false
						break
					}
				}
				if allDestroyed {
					q.Push(a)
					destroyed[a] = true
				}
			}
		}
		total += len(destroyed) - 1
	}
	return total
}

type Queue struct {
	l []string
}

func NewQueue() Queue {
	return Queue{l: make([]string, 0)}
}

func (q *Queue) Pop() string {
	v := q.l[0]
	q.l = q.l[1:]
	return v
}

func (q *Queue) Push(ps ...string) {
	for _, p := range ps {
		q.l = append(q.l, p)
	}
}

func (q *Queue) IsEmpty() bool {
	return len(q.l) == 0
}

func xyIntersect(a brick, b brick) bool {
	return max(a.x1, b.x1) <= min(a.x2, b.x2) && max(a.y1, b.y1) <= min(a.y2, b.y2)
}

func getBricksOnLevel(m []*brick, z int64) []brick {
	res := make([]brick, 0)
	for _, b := range m {
		if b.z1 == z {
			res = append(res, *b)
		}
	}
	return res
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
