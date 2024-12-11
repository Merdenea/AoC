package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	file, _ := os.Open(fmt.Sprintf("./%s", os.Args[1]))
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	line := scanner.Text()

	now := time.Now()
	one := partOne(line)
	fmt.Printf("Part I in [%s]: %d\n", time.Since(now).String(), one)

	now = time.Now()
	two := partTwo(line)
	fmt.Printf("Part II in [%s]: %d\n", time.Since(now).String(), two)

}

func partOne(l string) int64 {
	freeSpace := make([]int, 0)
	files := make([]int, 0)

	for i, s := range l {
		n, _ := strconv.ParseInt(string(s), 10, 64)
		if i%2 == 0 {
			files = append(files, int(n))
		} else {
			freeSpace = append(freeSpace, int(n))
		}
	}

	res := make([]int, 0)

	for i := 0; i < len(files); i++ {
		if i < len(files) {
			res = append(res, getDigits(i, files[i])...)
		}
		if i < len(freeSpace) {
			res = append(res, getDigits(-1, freeSpace[i])...)
		}
	}

	left, right := 0, len(res)-1

	for left < right {
		for res[left] >= 0 {
			left++
		}

		for res[right] < 0 {
			right--
		}

		if left < right {
			res[left], res[right] = res[right], res[left]
		}
	}

	return int64(checksum(res))
}

func checksum(arr []int) int {
	res := 0
	for i, v := range arr {
		if v >= 0 {
			res += i * v
		}
	}
	return res
}

func getDigits(val, repeated int) []int {
	res := make([]int, 0)
	for i := 0; i < repeated; i++ {
		res = append(res, val)
	}
	return res
}

func partTwo(l string) int64 {
	freeSpace := make([]int, 0)
	files := make([]int, 0)

	for i, s := range l {
		n, _ := strconv.ParseInt(string(s), 10, 64)
		if i%2 == 0 {
			files = append(files, int(n))
		} else {
			freeSpace = append(freeSpace, int(n))
		}
	}

	res := make([]int, 0)

	for i := 0; i < len(files); i++ {
		if i < len(files) {
			res = append(res, getDigits(i, files[i])...)
		}
		if i < len(freeSpace) {
			res = append(res, getDigits(-1, freeSpace[i])...)
		}
	}

	for right := len(res) - 1; right >= 0; {
		// print(res)
		for res[right] < 0 {
			right--
		}

		j := right
		for j >= 0 && res[j] == res[right] {
			j--
		}

		for left := 0; left < right; {
			for res[left] >= 0 {
				left++
			}

			if left < right {

				i := left
				for i < right && res[i] < 0 {
					i++
				}

				// fmt.Println(left, i-1)
				// fmt.Println(j+1, right)

				swap(res, left, i-1, j+1, right)
				// print(res)
				left = i
			}
		}

		right = j
	}

	return int64(checksum(res))
}

func print(res []int) {

	for _, v := range res {
		if v == -1 {
			fmt.Print(".")
		} else {
			fmt.Print(v)
		}
	}
	fmt.Println()
}

var swapped = make(map[int]bool)

func swap(arr []int, l1, l2, r1, r2 int) bool {
	if swapped[arr[r1]] {
		return false
	}
	if l2-l1 < r2-r1 {
		// fmt.Println("not swapping", arr[r1:r2+1])
		return false
	}

	swapped[arr[r1]] = true

	for i := 0; i <= r2-r1; i++ {
		arr[l1+i], arr[r2-i] = arr[r2-i], arr[l1+i]
	}
	return true
}
