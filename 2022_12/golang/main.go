package main

import "fmt"

func main() {
	result := solution(300, 2, 1)
	fmt.Println(result)
}

func solution(n int, x int, y int) (res int) {
	arr := []int{x, y}
	for i := 2; i < n; i++ {
		arr = append(arr, arr[i-1]+arr[i-2])
	}
	return arr[len(arr)-1]
}
