package ch001

import (
	"strconv"
	"strings"
)

func Solve(boxes []int) int {
	cache = make(map[string]int)
	return solve(boxes)
}

var cache map[string]int

func solve(boxes []int) int {
	if len(boxes) == 1 {
		return boxes[0]
	}

	key := getKey(boxes)
	if _, ok := cache[key]; ok {
		return cache[key]
	}

	rev := reverse(boxes)

	result := max(boxes[0]-solve(boxes[1:]), rev[0]-solve(rev[1:]))
	cache[key] = result

	return result
}

func getKey(boxes []int) string {
	s := make([]string, len(boxes))
	for _, v := range boxes {
		s = append(s, strconv.Itoa(v))
	}

	return strings.Join(s, "#")
}

func reverse(collection []int) []int {
	length := len(collection)
	result := make([]int, length)
	copy(result, collection)

	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = collection[j], collection[i]
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
