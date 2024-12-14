package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"
)

type node struct {
	i, j     int
	distance int
}

type Graph map[string]map[node]int

func NewGraph() Graph {
	return make(map[string]map[node]int)
}

func (n node) eq(b node) bool {
	return n.i == b.i && n.j == b.j
}

func (n node) Key() string {
	return fmt.Sprintf("%d+%d", n.i, n.j)
}

func key(i, j int) string {
	return fmt.Sprintf("%d+%d", i, j)
}

type Stack struct {
	l []node
}

func (s *Stack) Push(n node) {
	s.l = append(s.l, n)
}

func (s *Stack) Pop() node {
	v := s.l[len(s.l)-1]
	s.l = s.l[:len(s.l)-1]
	return v
}

func (s *Stack) IsEmpty() bool {
	return len(s.l) == 0
}

func NewStack() Stack {
	return Stack{l: make([]node, 0)}
}

type direction struct {
	i, j int
}

var (
	UP    = direction{-1, 0}
	DOWN  = direction{1, 0}
	LEFT  = direction{0, -1}
	RIGHT = direction{0, 1}

	dirs = map[rune][]direction{
		'^': {UP},
		'v': {DOWN},
		'<': {LEFT},
		'>': {RIGHT},
		'.': {UP, DOWN, LEFT, RIGHT},
	}

	gVisited   = make(map[string]bool)
	allVisited = make(map[string]bool)
)

func main() {
	file, _ := os.Open("./twentythree/in")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	m := make([][]rune, 0)
	for scanner.Scan() {
		l := scanner.Text()
		m = append(m, []rune(l))
	}

	now := time.Now()
	fmt.Printf("Part One %d in [%s]\n", getLongestPath(m, true), time.Now().Sub(now))
	now = time.Now()
	fmt.Printf("Part Two %d in [%s]\n", getLongestPath(m, false), time.Now().Sub(now))
}

func getLongestPath(m [][]rune, one bool) int {
	si, sj := getStart(m, 1)
	di, dj := getStart(m, -1)

	start := node{i: si, j: sj}
	end := node{i: di, j: dj}

	points := make(map[string]node, 0)
	points[start.Key()] = start
	points[end.Key()] = end

	for i, l := range m {
		for j, r := range l {
			if r == '#' {
				continue
			}
			neighbours := 0
			for _, d := range []direction{UP, DOWN, LEFT, RIGHT} {
				ni, nj := i+d.i, j+d.j
				if ni < 0 || nj < 0 || ni >= len(m) || nj >= len(m[0]) {
					continue
				}
				if m[ni][nj] != '#' {
					neighbours++
				}
			}
			if neighbours >= 3 {
				nn := node{i: i, j: j}
				points[nn.Key()] = nn
			}
		}
	}

	graph := NewGraph()

	for _, p := range points {
		graph[p.Key()] = make(map[node]int)
	}

	for _, startNode := range points {
		st := NewStack()
		st.Push(startNode)
		visited := make(map[string]bool)
		visited[startNode.Key()] = true

		for !st.IsEmpty() {
			cNode := st.Pop()

			if _, ok := points[cNode.Key()]; ok && cNode.distance != 0 {
				graph[startNode.Key()][cNode] = cNode.distance
				continue
			}

			if one {
				for _, d := range dirs[m[cNode.i][cNode.j]] {
					ni, nj := cNode.i+d.i, cNode.j+d.j
					if !(ni < 0 || nj < 0 || ni >= len(m) || nj >= len(m[0])) && m[ni][nj] != '#' && !visited[node{i: ni, j: nj}.Key()] {
						st.Push(node{ni, nj, cNode.distance + 1})
						visited[node{i: ni, j: nj}.Key()] = true
					}
				}
			} else {
				//{UP, DOWN, LEFT, RIGHT}
				for _, d := range []direction{UP, DOWN, LEFT, RIGHT} {
					ni, nj := cNode.i+d.i, cNode.j+d.j
					if !(ni < 0 || nj < 0 || ni >= len(m) || nj >= len(m[0])) && m[ni][nj] != '#' && !visited[node{i: ni, j: nj}.Key()] {
						st.Push(node{ni, nj, cNode.distance + 1})
						visited[node{i: ni, j: nj}.Key()] = true
					}
				}
			}

		}
	}

	//for k, v := range graph {
	//	fmt.Println(k, ":", v)
	//}

	return maxDfs(graph, start, end)
}

func maxDfs(graph Graph, start, end node) int {
	if start.eq(end) {
		return 0
	}

	maxPath := math.MinInt32
	gVisited[start.Key()] = true
	for next, distance := range graph[start.Key()] {
		if !gVisited[next.Key()] {
			maxPath = max(maxPath, maxDfs(graph, next, end)+distance)
		}
	}
	gVisited[start.Key()] = false
	return maxPath
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getStart(m [][]rune, mode int) (int, int) {
	if mode == -1 {
		for i := len(m) - 1; i >= 0; i-- {
			for j, r := range m[i] {
				if r == '.' {
					return i, j
				}
			}
		}
	}
	for i, l := range m {
		for j, r := range l {
			if r == '.' {
				return i, j
			}
		}
	}
	return -1, -1
}
