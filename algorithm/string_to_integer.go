package algorithm

import (
	"math"
)

func Atoi(str string) int32 {
	if len(str) == 0 {
		return 0
	}
	var ret int32
	isPositiveNumber := true
	for i := 0; i < len(str); i++ {
		a := str[i]
		if a == byte(' ') {
			continue
		}
		if ret == 0 {
			if a == byte('-') {
				isPositiveNumber = false
			} else if a == byte('+') {
				isPositiveNumber = true
			} else if a > byte('0') && a <= byte('9') {
				ret = int32(a - byte(0))
			} else {
				return 0
			}
		} else {
			if a > byte('0') && a <= byte('9') {
				newRet := ret*10 + int32(a-byte('0'))
				if (newRet-int32(a-byte('0')))/10 != ret {
					if isPositiveNumber {
						return math.MaxInt32
					} else {
						return math.MinInt32
					}
				}
				ret = newRet
			} else {
				break
			}
		}
	}
	if isPositiveNumber {
		return ret
	} else {
		return -ret
	}
}
