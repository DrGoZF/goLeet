package src

import (

)

func reverseString(s string) string {
    length := len(s)
    c := []byte(s)
    for i := 0; i < length / 2; i ++ {
	    temp := c[i]
	    c[i] = c[length - i - 1]
	    c[length - i - 1] = temp
    }
    return string(c)
}