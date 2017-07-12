package main

import (
  "fmt"
)

func main() {
  var i int
  fmt.Print("Number of test cases: ")
  fmt.Scanf("%d", &i)
  fmt.Println("You would like to have", i, "test cases")
  
  var j int = 0
  for j < i {
    var x, y, z int
    fmt.Print("First Number: ")
    fmt.Scanf("%d", &x)
    
    fmt.Print("Second Number: ")
    fmt.Scanf("%d", &y)
    
    fmt.Print("Third Number: ")
    fmt.Scanf("%d", &z)
    
    fmt.Println(x,y,z)
    
    j++
  }
}