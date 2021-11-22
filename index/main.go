package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	i := 32
	if i >= 0 && i < len(nums) {
		nums = append(nums[0:i], nums[i+1:]...)
		fmt.Println(nums)
	}
	fmt.Println(nums)
}

// END
