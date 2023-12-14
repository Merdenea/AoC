package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type tree struct {
	val   string
	left  *tree
	right *tree
	count int
}

var (
	mm = make(map[string]*tree)
)

func main() {
	file, _ := os.Open("./eight/in")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	instructions := scanner.Text()

	for scanner.Scan() {
		l := scanner.Text()
		if len(l) == 0 {
			continue
		}
		n := getNode(l)
		mm[n.val] = n
	}

	//ONE

	//found := false
	//count := 0
	//cr := mm["AAA"]
	//
	ins := []rune(instructions)

	//for !found {
	//	for _, r := range ins {
	//		//fmt.Println(cr.val)
	//		if cr.val == "ZZZ" {
	//			found = true
	//			break
	//		}
	//		if r == 'L' {
	//			cr = mm[cr.left.val]
	//		} else {
	//			cr = mm[cr.right.val]
	//		}
	//		count++
	//	}
	//fmt.Println(count)

	starts := make(map[string]*tree)

	now := time.Now()
	for k, v := range mm {
		if k[2:3] == "A" {
			starts[k] = v
		}
	}
	//fmt.Println(starts)

	ends := make(map[string]map[string]*tree)

	for k, v := range starts {
		ends[k] = make(map[string]*tree)
		done := false
		cn := v
		count := 0
		for !done {
			for _, r := range ins {
				//fmt.Println(v.val)
				ok, newNode := traverseOne(cn, r)
				count++
				if ok {
					if ends[k][newNode.val] != nil {
						done = true
					} else {
						newNode.count = count
						ends[k][newNode.val] = newNode
					}
				}
				cn = newNode
			}
		}
	}

	cc := make([]int, 0)
	for _, v := range ends {
		for _, p := range v {
			cc = append(cc, p.count)
		}
	}

	fmt.Println(LCM(1, 1, cc...))
	fmt.Println(time.Now().Sub(now))

	//count := 0
	//found := false
	//for !found {
	//
	//	for _, r := range ins {
	//		for k, v := range starts {
	//			//fmt.Println(v.val)
	//			_, newNode := traverseOne(v, r)
	//			starts[k] = newNode
	//		}
	//		count++
	//		if checkStarts(starts) {
	//			found = true
	//			break
	//		}
	//	}
	//
	//}
	//fmt.Println(count)
}

func checkStarts(starts map[string]*tree) bool {
	for _, v := range starts {
		if v.val[2:3] != "Z" {
			return false
		}
	}
	return true
}

func traverseOne(cr *tree, r rune) (bool, *tree) {
	if r == 'L' {
		cr = mm[cr.left.val]
	} else {
		cr = mm[cr.right.val]
	}
	if cr.val[2:3] == "Z" {
		return true, cr
	}
	return false, cr
}

func getNode(l string) *tree {
	sp := strings.Split(l, " = ")
	val := sp[0]
	spp := strings.Split(sp[1][1:len(sp[1])-1], ", ")

	left := spp[0]
	right := spp[1]
	return &tree{
		val:   val,
		left:  &tree{val: left},
		right: &tree{val: right},
	}
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
