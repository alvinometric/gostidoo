package main

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// Given two strings s and t, return true if t is an anagram of s
// and false otherwise.
func isAnagram(s string, t string) bool{
	// they're the same string
	if reflect.DeepEqual(s, t){
		return false
	}
	arr1 := strings.Split(s, "")
	arr2 := strings.Split(t, "")

	sort.Strings(arr1)
	sort.Strings(arr2)

	return reflect.DeepEqual(arr1, arr2)
}

func main(){
	fmt.Println(isAnagram("race", "care"));
	fmt.Println(isAnagram("elbow", "below"));
	fmt.Println(isAnagram("chien", "niche"));
}