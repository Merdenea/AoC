package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type rule struct {
	propType string
	cond     string
	value    int64
	dest     string
}

func (r rule) reverse() rule {
	newCond := "<"
	newVal := r.value + 1
	if r.cond == "<" {
		newCond = ">"
		newVal = r.value - 1
	}
	nr := rule{
		propType: r.propType,
		cond:     newCond,
		value:    newVal,
	}
	return nr
}

type part struct {
	x, m, a, s int64
}

func (p part) Sum() int64 {
	return p.x + p.m + p.a + p.s
}

func (p part) GetValue(s string) int64 {
	switch s {
	case "x":
		return p.x
	case "m":
		return p.m
	case "a":
		return p.a
	case "s":
		return p.s
	}
	return 0
}

const (
	Accept = "A"
	Reject = "R"

	StartWorkflow = "in"
)

func main() {
	file, _ := os.Open("./nineteen/in")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	workflows := make(map[string][]rule)
	for scanner.Scan() {
		l := scanner.Text()
		if len(l) > 0 {
			sp := strings.Split(l, "{")

			key := sp[0]
			sp[1] = sp[1][:len(sp[1])-1]
			rules := parseRules(sp[1])
			workflows[key] = rules
		} else {
			break
		}
	}

	parts := make([]part, 0)
	for scanner.Scan() {
		l := scanner.Text()
		sp := strings.Split(l[1:len(l)-1], ",")
		x, _ := strconv.ParseInt(sp[0][2:], 10, 64)
		m, _ := strconv.ParseInt(sp[1][2:], 10, 64)
		a, _ := strconv.ParseInt(sp[2][2:], 10, 64)
		s, _ := strconv.ParseInt(sp[3][2:], 10, 64)
		parts = append(parts, part{x, m, a, s})
	}

	now := time.Now()
	fmt.Printf("Part One %d in [%s]\n", getAcceptedSum(parts, workflows), time.Now().Sub(now))

	st := state{
		sMax: 4001,
		mMax: 4001,
		aMax: 4001,
		xMax: 4001,
		sMin: 0,
		mMin: 0,
		aMin: 0,
		xMin: 0,
	}
	now = time.Now()
	fmt.Printf("Part Two %d in [%s]\n", dfsTraverse(workflows, StartWorkflow, st), time.Now().Sub(now))

}

type state struct {
	sMax, mMax, aMax, xMax int64
	sMin, mMin, aMin, xMin int64
}

func (s state) comb() int64 {
	return (s.sMax - 1 - s.sMin) * (s.aMax - 1 - s.aMin) * (s.mMax - 1 - s.mMin) * (s.xMax - 1 - s.xMin)
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func (s state) update(r rule) state {
	if r.cond == "" {
		return s
	}
	ns := state{
		sMax: s.sMax,
		mMax: s.mMax,
		aMax: s.aMax,
		xMax: s.xMax,
		sMin: s.sMin,
		mMin: s.mMin,
		aMin: s.aMin,
		xMin: s.xMin,
	}
	switch r.propType {
	case "s":
		if r.cond == "<" {
			ns.sMax = min(ns.sMax, r.value)
		} else {
			ns.sMin = max(ns.sMin, r.value)
		}
		break
	case "m":
		if r.cond == "<" {
			ns.mMax = min(ns.mMax, r.value)
		} else {
			ns.mMin = max(ns.mMin, r.value)
		}
		break
	case "a":
		if r.cond == "<" {
			ns.aMax = min(ns.aMax, r.value)
		} else {
			ns.aMin = max(ns.aMin, r.value)
		}
		break
	case "x":
		if r.cond == "<" {
			ns.xMax = min(ns.xMax, r.value)
		} else {
			ns.xMin = max(ns.xMin, r.value)
		}
		break
	}
	return ns
}

func dfsTraverse(ws map[string][]rule, current string, state state) int64 {
	if current == Accept {
		return state.comb()
	}
	if current == Reject {
		return 0
	}
	rules := ws[current]
	total := int64(0)
	for _, r := range rules {
		total += dfsTraverse(ws, r.dest, state.update(r))
		state = state.update(r.reverse())
	}
	return total
}

func getAcceptedSum(parts []part, workflows map[string][]rule) int64 {
	total := int64(0)
	for _, p := range parts {
		workflowKey := StartWorkflow
		for !(workflowKey == Accept || workflowKey == Reject) {
			workflowKey = applyWorkflow(p, workflows[workflowKey])
		}

		if workflowKey == Accept {
			total += p.Sum()
		}
	}

	return total
}

func applyWorkflow(p part, rules []rule) string {
	for _, r := range rules {
		switch r.cond {
		case ">":
			if p.GetValue(r.propType) > r.value {
				return r.dest
			}
			break
		case "<":
			if p.GetValue(r.propType) < r.value {
				return r.dest
			}
			break
		default:
			return r.dest
		}
	}
	return ""
}

func parseRules(s string) []rule {
	sp := strings.Split(s, ",")
	rules := make([]rule, 0)
	for _, st := range sp {
		spp := strings.Split(st, ":")
		if len(spp) == 2 {
			val, _ := strconv.ParseInt(spp[0][2:], 10, 64)
			r := rule{
				propType: spp[0][0:1],
				cond:     spp[0][1:2],
				value:    val,
				dest:     spp[1],
			}
			rules = append(rules, r)
		} else {
			rules = append(rules, rule{dest: st})
		}
	}
	return rules
}
