package main

import (
	"fmt"
	"strconv"
	"strings"
)

func spellOut(num int) string{
	spelledOut := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
	return spelledOut[num - 1]
}

func countAndSay(sequence int) string{
	digits := strconv.Itoa(sequence);
	acc := make(map[string]int)
	result := ""

	for _,v := range strings.Split(digits, ""){
		_, hasKey := acc[v];
		if hasKey{
			acc[v]++
		} else{
			acc[v] = 1;
		}
	}

	i := 0;
	for k, v := range acc{
		if i != 0{
			result += ", then "
		}
		result += spellOut(v) + " " + k + "s"
		i++
	}

	return result
}

func main(){
	fmt.Println(countAndSay(111226666))
}