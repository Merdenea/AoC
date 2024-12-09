package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	cards = map[rune]int{
		'A': 13,
		'K': 12,
		'Q': 11,
		'J': 0,
		'T': 9,
		'9': 8,
		'8': 7,
		'7': 6,
		'6': 5,
		'5': 4,
		'4': 3,
		'3': 2,
		'2': 1,
	}
)

type hand struct {
	h     string
	value int
	hSt   int
}

func main() {
	file, _ := os.Open("./seven/in")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	mm := make([]hand, 0)
	for scanner.Scan() {
		sp := strings.Split(scanner.Text(), " ")
		value, _ := strconv.ParseInt(sp[1], 10, 64)
		mm = append(mm, hand{
			h:     sp[0],
			value: int(value),
			hSt:   getHandValue(sp[0], true),
		})
	}

	sort.Slice(mm, func(i, j int) bool {
		if mm[i].hSt == mm[j].hSt {
			rOne := []rune(mm[i].h)
			rTwo := []rune(mm[j].h)

			for k := 0; k < len(rOne); k++ {
				if cards[rOne[k]] < cards[rTwo[k]] {
					return true
				} else if cards[rOne[k]] > cards[rTwo[k]] {
					return false
				}
			}
		}
		return mm[i].hSt < mm[j].hSt
	})

	fmt.Println(mm)

	total := 0
	for i := 0; i < len(mm); i++ {
		total += (i + 1) * mm[i].value
	}
	fmt.Println(total)
}

func getHandValue(hand string, checkJoker bool) int {
	hr := []rune(hand)
	frqMap := make(map[rune]int)
	for _, h := range hr {
		frqMap[h]++
	}

	max := 0
	if frqMap['J'] > 0 && checkJoker {
		for k, _ := range frqMap {
			if k == 'J' {
				continue
			}
			val := getHandValue(strings.ReplaceAll(hand, "J", string(k)), false)
			//fmt.Println(hand, val)
			max = int(math.Max(float64(max), float64(val)))
		}
		val := getHandValue(hand, false)
		//fmt.Println(hand, val)
		max = int(math.Max(float64(max), float64(val)))
		return max
	}

	ln := len(frqMap)
	switch ln {
	case 1:
		// five of a kind all the same AAAAA
		return 7
	case 2:
		//four of a kind AAAAB
		if isKind(frqMap, 4) {
			return 6
		} else {
			// or full house 23332
			return 5
		}
	case 3:
		// three of a kind
		if isKind(frqMap, 3) {
			return 4
		}
		//two pair
		return 3
	case 4:
		// one pair
		return 2
	case 5:
		return 1
	default:
		return 1
	}
}

func isKind(hr map[rune]int, no int) bool {
	for _, v := range hr {
		if v == no {
			return true
		}
	}
	return false
}
