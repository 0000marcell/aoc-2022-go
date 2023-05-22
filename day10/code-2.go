package main

import (
  "fmt"
  "os"
  "bufio"
  //"math"
  "strings"
  "strconv"
)

func PrintImage(image []string) {
  var startpixel int = 0
  for i := 0; i < 6; i++ {
    fmt.Println(image[startpixel:startpixel + 40])
    startpixel += 40
  }
}

func Draw(image []string, regvalue int) []string {
  var postodraw int = len(image) - len(image)/40 * 40
  var spritepos int = regvalue
  var symboltodraw string = "."
  if postodraw == spritepos ||
     postodraw == spritepos - 1 ||
     postodraw == spritepos + 1 {
    symboltodraw = "#"
  }
  image = append(image, symboltodraw)
  return image
}

func main() {
  readFile, err := os.Open("input-1")
  if err != nil {
    fmt.Println(err)
  }
  fileScanner := bufio.NewScanner(readFile)
  var cyclecount int = 0
  var regvalue int = 1
  image := []string{}
  for fileScanner.Scan() {
    // if cyclecount > 240 {
    //   break
    // }
    image = Draw(image, regvalue)
    str := fileScanner.Text()
    if str == "noop" {
      cyclecount += 1
      continue
    } 
    strsplit := strings.Split(str, " ")
    if strsplit[0] == "addx" {
      cyclecount += 1
      image = Draw(image, regvalue)
      number, err := strconv.Atoi(string(strsplit[1]))
      if err != nil {
        fmt.Println("Error during conversion")
        return
      }
      cyclecount += 1
      regvalue += number
    }
    //fmt.Println("=================")
  }
  fmt.Println("image len: ", len(image))
  //fmt.Println("image: ", image)
  PrintImage(image)
}
