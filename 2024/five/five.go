package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, _ := os.Open(fmt.Sprintf("./%s", os.Args[1]))
	// file, _ := os.Open(fmt.Sprintf("./test"))

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	m := make(map[int64][]int64)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}
		st := strings.Split(line, "|")
		x, _ := strconv.ParseInt(st[0], 10, 64)
		y, _ := strconv.ParseInt(st[1], 10, 64)

		if _, ok := m[x]; ok {
			m[x] = append(m[x], y)
		} else {
			m[x] = make([]int64, 0)
			m[x] = append(m[x], y)
		}
	}

	updates := make([][]int64, 0)
	for scanner.Scan() {
		update := make([]int64, 0)
		line := scanner.Text()
		st := strings.Split(line, ",")
		for _, n := range st {
			nn, _ := strconv.ParseInt(n, 10, 64)
			update = append(update, nn)
		}
		updates = append(updates, update)
	}

	now := time.Now()
	one := partOne(m, updates)
	fmt.Printf("Part I in [%s]: %d\n", time.Since(now).String(), one)

	now = time.Now()
	two := partTwo(m, updates)
	fmt.Printf("Part II in [%s]: %d\n", time.Since(now).String(), two)

}

func partOne(adj map[int64][]int64, updates [][]int64) int64 {
	total := int64(0)
	for _, u := range updates {
		ok := true
		for i := 0; i < len(u); i++ {
			for j := i + 1; j < len(u); j++ {
				l := adj[u[i]]
				if !slices.Contains(l, u[j]) {
					ok = false
					break
				}
			}
			if !ok {
				break
			}
		}
		if ok {
			total += u[(len(u)-1)/2]
		}
	}
	return total
}

func checkUpdate(adj map[int64][]int64, u []int64) bool {
	for i := 0; i < len(u); i++ {
		for j := i + 1; j < len(u); j++ {
			l := adj[u[i]]
			// fmt.Println(u[i], u[j], l)
			if !slices.Contains(l, u[j]) {
				return false
			}
		}
	}
	return true
}

func partTwo(adj map[int64][]int64, updates [][]int64) int64 {
	total := int64(0)
	toFix := make([][]int64, 0)
	for _, u := range updates {
		ok := true
		for i := 0; i < len(u); i++ {
			for j := i + 1; j < len(u); j++ {
				l := adj[u[i]]
				if !slices.Contains(l, u[j]) {
					ok = false
					break
				}
			}
			if !ok {
				break
			}
		}
		if !ok {
			toFix = append(toFix, u)
		}
	}

	keys := make(map[int64]bool)
	for k, _ := range adj {
		keys[k] = true
	}

	for _, update := range toFix {
		// total += findFixV2(update, adj)
		total += findFixX2(update, adj)
	}
	return total
}

func findFixX2(update []int64, adj map[int64][]int64) int64 {
	sort.Slice(update, func(i, j int) bool {
		l := update[i]
		r := update[j]
		return slices.Contains(adj[l], r)
	})

	return update[(len(update)-1)/2]
}

// LOL - takes 1 minute
func findFixV2(update []int64, adj map[int64][]int64) int64 {
	updateMap := make(map[int64]bool)
	for _, u := range update {
		updateMap[u] = true
	}
	total := int64(0)

	for _, u := range update {
		node := bfs(u, adj, updateMap)
		if int(node.dist) == len(update)-1 {
			sp := strings.Split(node.path, ",")
			mid, _ := strconv.ParseInt(sp[(len(sp)-1)/2], 10, 64)
			total += mid
		}
	}
	return total
}

type node struct {
	val  int64
	dist int64
	path string
}

func bfs(root int64, adj map[int64][]int64, update map[int64]bool) node {
	q := make([]node, 0)
	q = append(q, node{root, 0, fmt.Sprintf("%d", root)})

	maxNode := node{}
	for len(q) > 0 {
		val := q[0]
		q = q[1:]

		next := adj[val.val]

		for _, n := range next {
			if update[n] {
				q = append(q, node{n, val.dist + 1, fmt.Sprintf("%s,%d", val.path, n)})
			}
		}
		maxNode = maxN(maxNode, val)
	}
	return maxNode
}

func maxN(a, b node) node {
	if a.dist > b.dist {
		return a
	}
	return b
}
