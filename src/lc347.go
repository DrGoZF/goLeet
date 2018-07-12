package src

import (

)

func topKFrequent(nums []int, k int) []int {
    fre := make(map[int] int)
    result := make(map[int] []int)
    
    for i := 0; i < len(nums); i ++ {
	    fre[nums[i]] = fre[nums[i]] + 1
    }
    
    for key, value := range fre {
	    result[value] = append(result[value], key)
    }
    
    var res []int
    for i := len(nums); i > 0; i-- {
        if arr, ok := result[i]; ok {
            for _, v := range arr {
                res = append(res, v)
                if len(res) == k { 
                    return res 
                }   
            }   
        }   
    }   
    return res
}