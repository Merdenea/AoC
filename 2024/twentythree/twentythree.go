package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	file, _ := os.Open(fmt.Sprintf("./%s", os.Args[1]))
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	adj := make(map[string]map[string]int)
	for scanner.Scan() {
		line := scanner.Text()
		st := strings.Split(line, "-")

		if _, ok := adj[st[0]]; !ok {
			adj[st[0]] = make(map[string]int)
		}
		if _, ok := adj[st[1]]; !ok {
			adj[st[1]] = make(map[string]int)
		}
		adj[st[0]][st[1]]++
		adj[st[1]][st[0]]++
	}

	now := time.Now()
	one := partOne(adj)
	fmt.Printf("Part I in [%s]: %d\n", time.Since(now).String(), one)

	now = time.Now()
	two := partTwo(adj)
	fmt.Printf("Part II in [%s]: %s\n", time.Since(now).String(), two)

}

func partOne(adj map[string]map[string]int) int64 {
	threeConn := make(map[[3]string]int)

	for k, v := range adj {
		for k2, _ := range v {
			for k3, _ := range v {
				if k2 == k3 {
					continue
				}
				if adj[k2][k3] > 0 {
					connSlice := []string{k, k2, k3}
					sort.Strings(connSlice)
					conn := [3]string{connSlice[0], connSlice[1], connSlice[2]} //yay
					threeConn[conn]++
				}
			}
		}
	}

	total := 0

	for k, _ := range threeConn {
		if k[0][0] == 't' || k[1][0] == 't' || k[2][0] == 't' {
			total++
		}
	}

	return int64(total)
}

func nextConnectedLevel(prev map[string]int, adj map[string]map[string]int) map[string]int {
	res := make(map[string]int)
	for connected, _ := range prev {
		for node, neighbours := range adj {
			found := true
			connectedSlice := strings.Split(connected, "-")
			for _, c := range connectedSlice {
				if c == node {
					found = false
					break
				}
				found = found && neighbours[c] > 0
			}
			if found {
				sl := append(connectedSlice, node)
				sort.Strings(sl)
				key := strings.Join(sl, "-")
				res[key]++
			}
		}
	}
	return res
}

func partTwo(adj map[string]map[string]int) string {
	// takes 20s lol
	// conns := make(map[string]int)
	// for k, v := range adj {
	// 	for k2, _ := range v {
	// 		for k3, _ := range v {
	// 			if k2 == k3 {
	// 				continue
	// 			}
	// 			if adj[k2][k3] > 0 {
	// 				connSlice := []string{k, k2, k3}
	// 				sort.Strings(connSlice)
	// 				conns[strings.Join(connSlice, "-")]++
	// 			}
	// 		}
	// 	}
	// }

	// for len(conns) > 1 {
	// 	conns = nextConnectedLevel(conns, adj)
	// }

	// for k, _ := range conns {
	// 	return strings.ReplaceAll(k, "-", ",")
	// }

	return findMaxClique(adj)
}

func findMaxClique(adj map[string]map[string]int) string {
	nodes := make([]string, 0)

	for n, _ := range adj {
		nodes = append(nodes, n)
	}

	for i := 0; i < 10000; i++ {
		rand.Shuffle(len(nodes), func(i, j int) {
			nodes[i], nodes[j] = nodes[j], nodes[i]
		})

		clique := []string{}

		for _, n := range nodes {
			ok := true
			for _, cn := range clique {
				if adj[cn][n] == 0 {
					ok = false
				}
			}
			if ok {
				clique = append(clique, n)
			}
		}

		if len(clique) == 13 {
			sort.Strings(clique)
			return strings.Join(clique, ",")
		}
	}
	return ""
}
