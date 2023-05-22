package main

import (
  "fmt"
  "os"
  "bufio"
)

func Uniqarray(arr []string) bool {
  var result bool = true
  for i, a := range arr {
    for j, b := range arr {
      if i != j {
        if a == b {
          result = false
          break
        }
      }
         
    }
  }
  return result
}

func main() {
  readFile, err := os.Open("input-1")
  if err != nil {
    fmt.Println(err)
  }
  s := bufio.NewScanner(readFile)
  s.Split(bufio.ScanRunes)
  chars := []string{}
  for s.Scan() {
    if s.Text() != "" {
      chars = append(chars, s.Text())    
      if len(chars) > 4 {
        if Uniqarray(chars[len(chars) - 4:len(chars)]) {
          break
        }
      }
    }
    fmt.Println("=============================")
  }
  fmt.Println(chars)
  fmt.Println("the answer is: ", len(chars))
}
