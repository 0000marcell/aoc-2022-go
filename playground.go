package main
import (
  "fmt"
  "math"
)

func main() {
  if -4 - -2 > 0 {
    fmt.Println("greather than zero")
  } else {
    fmt.Println("smaller than zero")
  }
  if math.Abs(float64(-4 - -2)) > 1 {
    fmt.Println("true")
  } else {
    fmt.Println("not true")
  }
}

