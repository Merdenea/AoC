package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	hits = 0
	mm   map[string]int

	valid map[string]bool
	vHits = 0
)

func main() {
	mm = make(map[string]int)
	valid = make(map[string]bool)

	file, _ := os.Open("./twelve/in")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	now := time.Now()
	fmt.Println("start: ", now.String())

	total := 0
	partTwo := true
	for scanner.Scan() {
		l := scanner.Text()
		sp := strings.Split(l, " ")
		//total += getArrangementsV1(replicate(sp[0], partTwo), replicateNums(partTwo, getNums(sp[1])))
		total += getArrangementsV2(replicate(sp[0], partTwo), replicateNums(partTwo, getNums(sp[1])))
	}

	fmt.Println(time.Now().Sub(now))
	fmt.Println(total)
	fmt.Println("end: ", time.Now().String())
	fmt.Println("hits: ", hits)
	fmt.Println("vhits: ", vHits)
}

func replicate(s string, partTwo bool) string {
	if partTwo {
		return fmt.Sprintf("%s?%s?%s?%s?%s", s, s, s, s, s)
	}
	return s
}

func replicateNums(partTwo bool, nums []int) []int {
	if partTwo {
		res := make([]int, 0)
		res = append(res, nums...)
		res = append(res, nums...)
		res = append(res, nums...)
		res = append(res, nums...)
		res = append(res, nums...)
		return res
	}
	return nums
}

func getArrangementsV1(st string, config []int) int {
	key := getKey(st, config)
	if val, ok := mm[key]; ok {
		hits++
		return val
	}

	if !strings.ContainsRune(st, '?') {
		if isValid(st, config) {
			return 1
		}
		return 0
	} else if !isValid(st[0:strings.Index(st, "?")+1], config) {
		return 0
	}

	s1 := strings.Replace(st, "?", ".", 1)
	s2 := strings.Replace(st, "?", "#", 1)
	res := getArrangementsV1(s1, config) + getArrangementsV1(s2, config)
	mm[getKey(st, config)] = res
	return res
}

func getArrangementsV2(st string, config []int) int {
	key := getKey(st, config)
	if val, ok := mm[key]; ok {
		hits++
		return val
	}

	if len(st) == 0 {
		if len(config) == 0 {
			return 1
		}
		return 0
	}

	res := 0

	if strings.HasPrefix(st, ".") {
		res = getArrangementsV2(st[1:], config)
	} else if strings.HasPrefix(st, "?") {
		s1 := strings.Replace(st, "?", ".", 1)
		s2 := strings.Replace(st, "?", "#", 1)
		res = getArrangementsV2(s1, config) + getArrangementsV2(s2, config)
	} else { // prefix == #
		if len(config) == 0 {
			// too many broken springs groups
			res = 0
		} else if len(st) < config[0] {
			// too few springs left
			res = 0
		} else if strings.Contains(st[:config[0]], ".") { // should only contain ?#, so not enough broken
			res = 0
		} else if len(config) > 1 {
			if len(st) < config[0]+1 || st[config[0]] == '#' {
				// group needs to be separated with a . (hence +1 or # check)
				res = 0
			} else {
				// increment over the group +plus the '.' separating the groups
				res = getArrangementsV2(st[config[0]+1:], config[1:])
			}
		} else {
			// advance config and st to the end of config
			res = getArrangementsV2(st[config[0]:], config[1:])
		}
	}

	mm[getKey(st, config)] = res
	return res
}

func getKey(st string, c []int) string {
	sb := strings.Builder{}
	for _, i := range c {
		sb.WriteRune(rune(i))
		sb.WriteRune('-')
	}
	return fmt.Sprintf("%s:%s", st, sb.String())
}

func isValid(s string, config []int) bool {
	key := getKey(s, config)
	if v, ok := valid[key]; ok {
		vHits++
		return v
	}
	c := 0
	sb := strings.Builder{}
	for _, r := range []rune(s) {
		if r == '?' {
			return true
		}
		if r == '#' {
			sb.WriteRune(r)
		} else if sb.Len() > 0 {
			if c >= len(config) {
				valid[key] = false
				return false
			}
			if sb.Len() == config[c] {
				sb.Reset()
				c++
			} else {
				valid[key] = false
				return false
			}
		}
	}
	if sb.Len() > 0 && c < len(config) {
		if sb.Len() == config[c] {
			sb.Reset()
			c++
		} else {
			valid[key] = false
			return false
		}
	} else if sb.Len() > 0 {
		c++
	}
	if c == len(config) {
		valid[key] = true
		return true
	}
	valid[key] = false
	return false
}

func isValidV2(s string, conf []int) bool {
	if strings.ContainsRune(s, '?') {
		return false
	}
	re := regexp.MustCompile("\\.+")
	sp := re.Split(s, -1)

	if len(sp) > 1 && sp[0] == "" {
		sp = sp[1:]
	}
	if len(sp) > 0 && sp[len(sp)-1] == "" {
		sp = sp[:len(sp)-1]
	}

	if len(conf) != len(sp) {
		return false
	}

	for i := 0; i < len(conf); i++ {
		if len(sp[i]) != conf[i] {
			return false
		}
	}
	return true
}

func getNums(s string) []int {
	nums := make([]int, 0)
	sp := strings.Split(s, ",")
	for _, c := range sp {
		n, _ := strconv.ParseInt(c, 10, 32)
		nums = append(nums, int(n))
	}
	return nums
}
