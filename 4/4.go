package main

import (
  "fmt"
  "math"
)

func main() {
  var i int
  
  fmt.Print("Number of test cases: ")
  _, err := fmt.Scanf("%d", &i)
  
  if err != nil {
    fmt.Println(err)
  } else {
    if i >= 1 && i <= 100 {
      fmt.Println("You would like to have", i, "test cases")

      var j int = 0
      var answer[] int
      for j < i {
        var A, B, K int
        fmt.Println("Case", j+1)
        fmt.Print("First Number: ")
        fmt.Scanf("%d", &A)

        if ValidateAB(A) {
          fmt.Print("Second Number: ")
          fmt.Scanf("%d", &B)

          if ValidateAB(B) {
            fmt.Print("Third Number: ")
            fmt.Scanf("%d", &K)

            if ValidateK(K) {
              answer = append(answer, TotalDivisible(A,B,K))
              j++ 
            } else {
              answer = answer[:0]
              fmt.Println("The third number must be between 0-10000")
            } 
          } else {
            fmt.Println("The second number must be between 0-10001")
          }
        } else {
          fmt.Println("The second number must be between 0-10001")
        } 
      }

      fmt.Print("\nOutput\n")

      if len(answer) > 0 {
        j = 1
        for j <= i {
          fmt.Printf("Case %d: ", j)
          fmt.Print(answer[j-1], "\n")
          j++
        }
      }
    } else {
      fmt.Println("Please input number from 1 to 100 only")
    } 
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