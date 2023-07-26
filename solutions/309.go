package main

import "fmt"

func explode(str string) []string{
	fmt.Println(str)

	var result = []string{"Hello"}
	return result
}


func main(){
	explode("hello");

}