package main

import (
	"fmt"
	"strings"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    var lenNums1 = len(nums1)
	var lenNums2 = len(nums2)
	var totalSize = lenNums1 + lenNums2
    var target = make([]int, 0, totalSize)
    var i, j, k = 0, 0, 0
    for ; i <lenNums1 && j < lenNums2; k++ {
        if nums1[i] <= nums2[j] {
            target[k] = nums1[i]
            i = i + 1
        } else {
            target[k] = nums2[j]
            j = j + 1
        }
    }
    if i == lenNums1 {
        target = append(target, nums2[j:]...)
    } else {
        target = append(target, nums1[i:]...)
	}
	if totalSize % 2 == 0 {
		return float64(target[totalSize/2] + target[(totalSize-1)/2])
	}
	return float64(target[totalSize/2])
	
}

func main() {
	findMedianSortedArrays([1,3,4],[2,5,6])
}
