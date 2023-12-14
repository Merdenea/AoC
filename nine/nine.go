package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("./nine/in")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	//nums := make([][]int64, 0)
	total := int64(0)
	for scanner.Scan() {
		l := scanner.Text()
		total += getEstimate2(getNums(l))
	}
	fmt.Println(total)
}

func getEstimate2(nums []int64) int64 {
	//fmt.Println(nums)
	ll := len(nums)
	if ll == 1 {
		return nums[0]
	}
	val := nums[0]
	return val - getEstimate2(diff(nums))
}

func getEstimate1(nums []int64) int64 {
	//fmt.Println(nums)
	ll := len(nums)
	if ll == 1 {
		return nums[0]
	}
	val := nums[ll-1]
	return val + getEstimate1(diff(nums))
}

func diff(nums []int64) []int64 {
	res := make([]int64, len(nums)-1)
	for i := 1; i < len(nums); i++ {
		res[i-1] = nums[i] - nums[i-1]
	}
	return res
}

func getNums(l string) []int64 {
	ns := make([]int64, 0)
	sp := strings.Split(l, " ")
	for _, s := range sp {
		ns = append(ns, getNumber(s))
	}
	return ns
}

func getNumber(s string) int64 {
	i := strings.LastIndex(s, " ")
	v, _ := strconv.ParseInt(s[i+1:], 10, 64)
	return v
}
