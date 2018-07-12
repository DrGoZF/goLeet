package src

import (

)

func productExceptSelf(nums []int) []int {
    length := len(nums)
    if length == 0 {
        return nums
    }
    part1, part2 := make([]int, length), make([]int, length)
    part1[0] = 1
    part2[length - 1] = 1
    for i := 1; i < length; i ++ {
	    part1[i] = part1[i - 1] * nums[i - 1]
    }
    for i := length - 2; i >= 0; i -- {
	    part2[i] = part2[i + 1] * nums[i + 1]
    }
    for i := 0; i < length; i ++ {
	    part1[i] = part1[i] * part2[i]
    }
    return part1
}