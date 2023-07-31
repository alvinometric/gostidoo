package main

import (
	"fmt"
	"sort"
	"strings"
)

// This week's question:
// Given a string, separate it into groups of non-space equal characters, sorted.
//
// Example:
// explodeString('Ahh, abracadabra!')
// ['!',',','A','aaaaa','bb','c','d','hh','rr']
func explode(str string) []string{
	var letters []string = strings.Split(str, "");
	var result []string;
	var acc = make(map[string]int)

	for _,letter := range letters{
		acc[letter]++;
	}

	for k, v := range acc{
		 str := "";
		for i := 0; i < v; i++ {
			str = str + k
		}
		result = append(result, str)
	}

	sort.Strings(result)
	return result
}


func main(){
	res := explode("Ahh, abracadabra!")
	fmt.Println(res)
}