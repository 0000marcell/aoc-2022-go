package main

import (
  "fmt"
  "os"
  "bufio"
  //"math"
  "strings"
  "strconv"
)

func CalcCycle(cyclecount int, regvalue int) int {
  var result int = 0
  if cyclecount == 20 || 
    cyclecount == 60 ||
    cyclecount == 100 ||
    cyclecount == 140 ||
    cyclecount == 180 ||
    cyclecount == 220 {
    // fmt.Println("cyclecount: ", cyclecount)
    // fmt.Println("regvalue: ", regvalue)
    result = regvalue * cyclecount
    //fmt.Println("new result: ", result)
  }
  return result
}

func main() {
  readFile, err := os.Open("input-1")
  if err != nil {
    fmt.Println(err)
  }
  fileScanner := bufio.NewScanner(readFile)
  var total int = 0
  var cyclecount int = 0
  var regvalue int = 1
  for fileScanner.Scan() {
    str := fileScanner.Text()
    if str == "noop" {
      cyclecount += 1
      total += CalcCycle(cyclecount, regvalue) 
      continue
    } 
    strsplit := strings.Split(str, " ")
    if strsplit[0] == "addx" {
      cyclecount += 1
      total += CalcCycle(cyclecount, regvalue) 
      number, err := strconv.Atoi(string(strsplit[1]))
      if err != nil {
        fmt.Println("Error during conversion")
        return
      }
      if cyclecount == 220 {
        fmt.Println("before 220")
        fmt.Println("regvalue: ", regvalue)
      }
      cyclecount += 1
      if cyclecount == 220 {
        fmt.Println("after 220")
        fmt.Println("regvalue: ", regvalue)
      }
      total += CalcCycle(cyclecount, regvalue) 
      regvalue += number
    }
    fmt.Println("=================")
    if cyclecount == 220 {
      break
    }
  }
  fmt.Println("total: ", total)
}
