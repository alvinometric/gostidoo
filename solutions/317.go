package main

import (
	"fmt"
	"slices"
)

func separateAndSort(arr []int) ([][]int){
	even := []int{}
	odd := []int{}
	
	for _, v := range arr {
		if v % 2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}

	slices.Sort(even)
	slices.Sort(odd)
	return [][]int{even, odd}
}

func main(){
	result := separateAndSort([]int{4,3,2,1,5,7,8,9})
	fmt.Println(result)
}