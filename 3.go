package main

import "fmt"

func g(str string) string {
	var i int = 0
	var new_str string = ""
	for i < len(str)-1 {
		new_str = new_str + string(str[i+1])
		i++
	}
	return new_str
}

func f(str string) string {
  
}

func main() {
	fmt.Println(g("fruits"))
}