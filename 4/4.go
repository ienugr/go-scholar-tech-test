package main

import (
    "fmt"
    "math"
    "bufio"
    "os"
    "strconv"
)
	
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    file, _ := os.Open("input.in")
    scanner := bufio.NewScanner(file)
    
    var input[] string
    var lines[] int
    
    for scanner.Scan() {    
        input = append(input, scanner.Text())
    }
    
    // Convert string in array to int
    for _, input := range input {
		j, err := strconv.Atoi(input)
		if err != nil {
			panic(err)
		}
		lines = append(lines, j)
	}
        
    x, y := 1, 4
    ytop := lines[0]*3+1
    
    var answer[] int
    
    for y <= ytop {
        answer = append(answer, TotalDivisible(lines[x], lines[x+1], lines[y-1]))
        x = y
        y += 3
    }
    
    if len(answer) > 0 {
        j := 1
        
        f, err := os.Create("answer.txt")
        check(err)
        
        defer f.Close()
        
        var str string
        for j <= lines[0] {
            str += fmt.Sprintf("Case %d: %d\n", j, answer[j-1])
            j++
        }
        f.WriteString(str)
        
        fmt.Println("Done")
    }
}

func IsDivisible(x int, y int) bool {
    if math.Mod(float64(x), float64(y)) == 0 {
        return true
    } else {
        return false
    }
}

func TotalDivisible(a int, b int, k int) int {
    var counter int = 0
    for a <= b {
        if IsDivisible(a, k) {
            counter++
        }
        a++
    }
    return counter
}

func ValidateAB(x int) bool {
    if x >= 1 && x <=10000 {
        return true
    }
    return false
}

func ValidateK(x int) bool {
    if x >= 1 && x <10000 {
        return true
    }
    return false
}