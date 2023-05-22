package main

import (
  "fmt"
  "os"
  "bufio"
  "math"
  "strings"
  "strconv"
)

func CheckIfExists(arr []string, element string) bool {
  var result bool = false
  for _, x := range arr {
    if x == element {
        result = true
        break
    }
  }
  return result 
}

func Move(knot []int, pos int, dir string) []int {
  if dir == "+"  {
    knot[pos] += 1
  }

  if dir == "-" {
    knot[pos] -= 1
  }
  return knot
}

func main() {
  //var n1 Node 
  readFile, err := os.Open("testinput-2")
  if err != nil {
    fmt.Println(err)
  }
  fileScanner := bufio.NewScanner(readFile)
  m := make([][]int, 10)
  for i := 0; i < 10; i++ {
    m[i] = append(m[i], 0, 0)
  }
  fmt.Println("initial pos: ", m)
  tvisitedpos := []string{"0|0"}
  for fileScanner.Scan() {
    str := fileScanner.Text()
    strsplit := strings.Split(str, " ")

    number, err := strconv.Atoi(string(strsplit[1]))
    if err != nil {
      fmt.Println("Error during conversion")
      return
    }

    for i := 0; i < number; i++ {
      if strsplit[0] == "D" {
        m[0] = Move(m[0], 0, "+")
      }

      if strsplit[0] == "U" {
        m[0] = Move(m[0], 0, "-")
      }

      if strsplit[0] == "R" {
        m[0] = Move(m[0], 1, "+")
      }

      if strsplit[0] == "L" {
        m[0] = Move(m[0], 1, "-")
      }
      fmt.Println(m[0])

      // Others logic 
      for i := 0; i < 9; i++ {
        // H is m[i]
        // T is m[i+1]
        //fmt.Println("comparing: ", i, i+1)
        var itomove int = -1
        var dirtomove string = ""
        var directmove bool = false
        if math.Abs(float64(m[i][0] - m[i+1][0])) > 1 {
          itomove = 0
          if m[i][0] - m[i+1][0] > 0 {
            dirtomove = "+"
          } else {
            dirtomove = "-"
          }
          if m[i][1] == m[i+1][1] {
            directmove = true
          } else {
            directmove = false 
          }
        } else if math.Abs(float64(m[i][1] - m[i+1][1])) > 1 {
          itomove = 1
          if m[i][1] - m[i+1][1] > 0 {
            dirtomove = "+"
          } else {
            dirtomove = "-"
          }
          if m[i][0] == m[i+1][0] {
            directmove = true
          } else {
            directmove = false 
          }
        }

        if itomove == -1  {
          continue
        }

        if dirtomove == "" {
          continue
        }

        if dirtomove == "" {
          fmt.Println("errror!!!!!!!!!!!!!>>>>>>>>>>.")
          continue
        }

        if itomove == -1 {
          fmt.Println("errror !!!!!!!!!!!!!!!>>>>>>>>>>>>>>>")
          return
        }

        m[i+1] = Move(m[i+1], itomove, dirtomove)

        if directmove == false {
          if itomove == 0 {
            if m[i][1] != m[i+1][1] {
              m[i+1][1] = m[i][1]
            }
          } else if itomove == 1 {
            if m[i][0] != m[i+1][0] {
              m[i+1][0] = m[i][0]
            }
          }
        }
        

        
        fmt.Println(m)
        if i + 1 == 9 {
          var tpos string = strconv.Itoa(m[i + 1][0]) + "|" + strconv.Itoa(m[i+1][1])
          //fmt.Println("tvisitedpos: ", tpos)
          if !CheckIfExists(tvisitedpos, tpos) {
            //fmt.Println(tpos)
            //fmt.Println(m)
            tvisitedpos = append(tvisitedpos, tpos)
          }else {
            fmt.Println("pos already exist!")
          }
          //fmt.Println("====================")
        }
      }
    }
  }
  fmt.Println(tvisitedpos)
  fmt.Println("all tvisitedpos: ", len(tvisitedpos))
}
