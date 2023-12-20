package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Pulse struct {
	source, dest string
	isHigh       bool
}

type Queue struct {
	l []Pulse
}

func NewQueue() Queue {
	return Queue{l: make([]Pulse, 0)}
}

func (q *Queue) Pop() Pulse {
	v := q.l[0]
	q.l = q.l[1:]
	return v
}

func (q *Queue) Push(ps ...Pulse) {
	for _, p := range ps {
		q.l = append(q.l, p)
	}
}

func (q *Queue) IsEmpty() bool {
	return len(q.l) == 0
}

type module struct {
	name  string
	mType string
	isOn  bool
	mem   map[string]*Pulse
	dest  []string
}

func (m module) getState(isHigh bool) string {
	if m.mType == FlipFlop {
		return fmt.Sprintf("%s+%s+%s", m.name, strconv.FormatBool(m.isOn), formatHigh(isHigh))
	}
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%s+%s+", m.name, m.mType))
	for k, v := range m.mem {
		sb.WriteString(fmt.Sprintf("%s+%s", k, formatHigh(v.isHigh)))
	}
	sb.WriteString(formatHigh(isHigh))
	return sb.String()
}

var (
	moduleStates = make(map[string]int)
	modulePeriod = make(map[string]int)
	pathsToRX    = make(map[string]bool)
)

func (m *module) ReceivePulse(source string, isHigh bool) ([]Pulse, bool) {
	if m.mType == FlipFlop {
		if isHigh {
			//ignore
			return nil, false
		}
		//low pulse
		if m.isOn {
			//send low
			m.isOn = false
			return toPulses(m.dest, false, m.name), true
		} else {
			// send high
			m.isOn = true
			return toPulses(m.dest, true, m.name), true
		}
	} else { //Conjunction
		// remember low by default
		m.mem[source] = &Pulse{isHigh: isHigh}
		if m.CheckAllHigh() {
			return toPulses(m.dest, false, m.name), true
		}
		return toPulses(m.dest, true, m.name), true
	}
}

func (m *module) CheckAllHigh() bool {
	for _, v := range m.mem {
		if v.isHigh == false {
			return false
		}
	}
	return true
}

func toPulses(dest []string, isHigh bool, source string) []Pulse {
	res := make([]Pulse, 0, len(dest))
	for _, d := range dest {
		res = append(res, Pulse{
			source: source,
			dest:   d,
			isHigh: isHigh,
		})
	}
	return res
}

const (
	Broadcaster = "broadcaster"
	FlipFlop    = "%"
	Conjunction = "&"
)

func main() {
	file, _ := os.Open("./twenty/in")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	seq := make(map[string]*module)
	for scanner.Scan() {
		l := scanner.Text()
		sp := strings.Split(l, " -> ")

		if strings.Contains(sp[0], "broadcaster") {
			seq[sp[0]] = &module{name: sp[0], dest: strings.Split(sp[1], ", ")}
		} else {
			name := sp[0][1:]
			mType := sp[0][:1]
			seq[name] = &module{
				name:  name,
				mType: mType,
				isOn:  false,
				mem:   make(map[string]*Pulse),
				dest:  strings.Split(sp[1], ", "),
			}
		}
	}

	//set all Conjunction type modules mem to false?
	for k, v := range seq {
		for _, d := range v.dest {
			md, ok := seq[d]
			if ok && md.mType == Conjunction {
				md.mem[k] = &Pulse{isHigh: false}
			}
		}
	}

	now := time.Now()
	//fmt.Printf("Part One %d in [%s]\n", getTotalSignals(seq, 1000, false), time.Now().Sub(now))
	// state needs to be reset for part two
	fmt.Printf("Part Two %d in [%s]\n", findMinButtonPresses(seq), time.Now().Sub(now))

}

// TODO: what a hack...
func findMinButtonPresses(seq map[string]*module) int {
	findAllPaths(seq, "broadcaster", "rx", "broadcaster")
	getTotalSignals(seq, -1, true)

	min := math.MaxInt64
	for path, _ := range pathsToRX {
		// if we can get almost all periods in reasonable time, we can compute the last ones
		// assumes we have all -1
		// and last one before rx is (&) conjunction type -> which requires all other prev connections to have high state saved
		sp := strings.Split(path[1:len(path)-1], ",")

		arr := make([]int, 0)
		secondToLast := sp[len(sp)-2] // "&kc"

		connections := getConnections(seq, secondToLast)
		for _, c := range connections {
			arr = append(arr, modulePeriod[c+"+high"])
		}
		lcm := LCM(1, 1, arr...)
		if lcm < min {
			min = lcm
		}
	}
	return min
}

func getConnections(seq map[string]*module, last string) []string {
	res := make([]string, 0)
	for k, v := range seq {
		for _, d := range v.dest {
			if d == last {
				res = append(res, k)
			}
		}
	}

	return res
}

func findAllPaths(seq map[string]*module, start, end string, current string) {
	if strings.Count(current, start) > 1 {
		return
	}

	if start == end {
		pathsToRX[current] = true
		//fmt.Println(current)
		return
	}

	mod := seq[start]

	for _, d := range mod.dest {
		findAllPaths(seq, d, end, current+","+d)
	}
}

func getTotalSignals(seq map[string]*module, buttonPresses int, partTwo bool) int64 {
	q := NewQueue()

	low := int64(0)
	high := int64(0)

	for i := 0; i < buttonPresses || partTwo; i++ {
		// button to broadcaster
		low++

		for _, b := range seq["broadcaster"].dest {
			q.Push(Pulse{
				source: "broadcaster",
				dest:   b,
				isHigh: false,
			})
		}

		for !q.IsEmpty() {
			pulse := q.Pop()
			if pulse.isHigh {
				high++
			} else {
				low++
			}

			receiver := seq[pulse.dest]
			if receiver == nil {
				continue
			}
			if p, ok := receiver.ReceivePulse(pulse.source, pulse.isHigh); ok {
				if v, okk := moduleStates[receiver.getState(p[0].isHigh)]; okk {
					if _, set := modulePeriod[receiver.name+"+"+formatHigh(p[0].isHigh)]; !set {
						modulePeriod[receiver.name+"+"+formatHigh(p[0].isHigh)] = v
					}
				}
				moduleStates[receiver.getState(p[0].isHigh)] = i + 1
				q.Push(p...)
			}
		}

		if partTwo {
			if len(modulePeriod) == (len(seq)-1)*2-1 { // def can do better than this
				break
			}
		}
	}

	//fmt.Printf("low [%d]  high [%d]\n", low, high)
	return low * high
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

func formatHigh(high bool) string {
	if high {
		return "high"
	}
	return "low"
}
