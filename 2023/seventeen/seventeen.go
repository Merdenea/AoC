package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type PriorityQueue []*node

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].currentHeatLoss < pq[j].currentHeatLoss
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*node)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *node) {
	heap.Fix(pq, item.index)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

type direction struct {
	i, j int
}

func (d direction) Reversed() direction {
	return direction{d.i * -1, d.j * -1}
}

type node struct {
	currentHeatLoss       int64
	i, j                  int
	currentDirectionCount int
	direction             direction

	index int
}

func (n node) getKey() string {
	return fmt.Sprintf("%d+%d+%d+%d+%d", n.i, n.j, n.direction.i, n.direction.j, n.currentDirectionCount)
}

func (n node) ContinueDirection() (int, int) {
	return n.i + n.direction.i, n.j + n.direction.j
}

func (n node) GetNewIJ(dir direction) (int, int) {
	return n.i + dir.i, n.j + dir.j
}

var (
	UP    = direction{1, 0}
	DOWN  = direction{-1, 0}
	LEFT  = direction{0, -1}
	RIGHT = direction{0, 1}
	NONE  = direction{0, 0}
)

func main() {
	file, _ := os.Open("./seventeen/in")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	m := make([][]int64, 0)
	for scanner.Scan() {
		l := scanner.Text()
		numsL := strings.Split(l, "")
		nums := make([]int64, 0)
		for _, n := range numsL {
			v, _ := strconv.ParseInt(n, 10, 64)
			nums = append(nums, v)
		}
		m = append(m, nums)
	}

	now := time.Now()
	fmt.Printf("Part One: %d in [%s]\n", DjK(m, 1, 3), time.Now().Sub(now))
	now = time.Now()
	fmt.Printf("Part Two: %d in [%s]\n", DjK(m, 4, 10), time.Now().Sub(now))
}

func DjK(m [][]int64, stepSize, maxStraightLine int) int64 {
	visited := make(map[string]bool)
	pq := make(PriorityQueue, 0)

	startNode := node{
		currentHeatLoss:       0,
		currentDirectionCount: 0,
		i:                     0,
		j:                     0,
		direction:             NONE,
	}

	pq.Push(&startNode)
	heap.Init(&pq)

	for pq.Len() > 0 {
		currentNode := heap.Pop(&pq).(*node)

		if currentNode.i == len(m)-1 && currentNode.j == len(m[0])-1 && currentNode.currentDirectionCount >= stepSize {
			return currentNode.currentHeatLoss
		}

		if visited[currentNode.getKey()] {
			continue
		}
		visited[currentNode.getKey()] = true

		if currentNode.currentDirectionCount < maxStraightLine && currentNode.direction != NONE {
			// Continue in the same direction
			newI, newJ := currentNode.ContinueDirection()
			if newI >= 0 && newJ >= 0 && newI < len(m) && newJ < len(m[0]) {
				nn := &node{
					currentHeatLoss:       currentNode.currentHeatLoss + m[newI][newJ],
					i:                     newI,
					j:                     newJ,
					currentDirectionCount: currentNode.currentDirectionCount + 1,
					direction:             currentNode.direction,
				}
				pq.Push(nn)
			}
		}
		// Add all valid turns
		if currentNode.currentDirectionCount >= stepSize || currentNode.direction == NONE {
			for _, dir := range []direction{UP, DOWN, LEFT, RIGHT} {
				if dir != currentNode.direction && dir != currentNode.direction.Reversed() {
					newI, newJ := currentNode.GetNewIJ(dir)
					if newI >= 0 && newJ >= 0 && newI < len(m) && newJ < len(m[0]) {
						nn := &node{
							currentHeatLoss:       currentNode.currentHeatLoss + m[newI][newJ],
							i:                     newI,
							j:                     newJ,
							currentDirectionCount: 1,
							direction:             dir,
						}
						pq.Push(nn)
					}
				}
			}
		}
	}
	return -1
}
