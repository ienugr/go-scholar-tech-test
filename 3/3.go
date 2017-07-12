package main

import "fmt"

func g(str string) string {
	var i int = 0
	var new_str string = ""
	for i < len(str) - 1 {
		new_str = new_str + string(str[i+1])
		i++
	}
	return new_str
}

func f(str string) string {
  if len(str) == 0 {
    return ""
  } else if len(str) == 1 {
    return str
  } else {
    return f(g(str)) + string(str[0])
  }
}

func h(n int, str string) string {
	for n != 1 {
		if n % 2 == 0 {
			n = n/2
		} else {
			n = 3*n + 1
		}
		str = f(str)
	}
	return str
}

func pow(x int, y int) int {
	if y == 0 {
		return 1
	} else {
		return x * pow(x, y-1)
	}
}

func main() {
	fmt.Print(h(1, "fruits"))
	fmt.Print(h(2, "fruits"))
	fmt.Print(h(5, "fruits"))
	fmt.Print(h(pow(2, 10), "fruits"))
	fmt.Print(h(pow(2, 37), "fruits"))
}