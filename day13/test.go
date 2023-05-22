package main

import (
  "fmt"
  "reflect"
)

func Test(list1 interface{}) {
  s := list1.(type)
  switch s {
  case []interface{}:
    fmt.Println(s)
    fmt.Println("is a slice")
  default:
    fmt.Println(s)
    fmt.Println("is not a slice")
  }
}

func main() {
  list1 := []any{'[','1','1','3','1','1',']'}
  Test(list1) 
}
