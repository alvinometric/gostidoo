// You have a faulty keyboard.
// Whenever you type a vowel on it (a,e,i,o,u,y), it reverses the string that you have written,
// instead of typing the character.
// Typing other characters works as expected. Given a string, return what will be on the screen after typing with your faulty keyboard.
// > faultyKeeb('string')
// > 'rtsng'

package main

import (
	"fmt"
	"slices"
	"strings"
)

func reverse(str string) string{
	chars := strings.Split(str, "");
	reversed := ""
	for i := len(chars) - 1; i >= 0; i-- {
		reversed += chars[i];
	}
	return reversed
}

func isVowel(letter string) bool{
	vowels := []string{"a","e","i","o","u","y"}
	return slices.Contains(vowels, letter)
}

func faulty(input string) string{
	chars := strings.Split(input, "")
	result := ""

	for _, char := range chars {
		if isVowel(char) {
			result = reverse(result)
		} else {
			result += char
		}
  }

	return result
}

func main(){
	reversed := faulty("string") // rtsng
	fmt.Println(reversed)
}