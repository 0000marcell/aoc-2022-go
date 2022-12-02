package main
 
import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)
 
func main() {
  readFile, err := os.Open("input")
  
  if err != nil {
    fmt.Println(err)
  }
  fileScanner := bufio.NewScanner(readFile)

  fileScanner.Split(bufio.ScanLines)

  for fileScanner.Scan() {
    str := fileScanner.Text()
    if str == "" { 
      continue
    }

    number, err := strconv.Atoi(fileScanner.Text())

    if err != nil {
      fmt.Println("Error during conversion")
      return
    }

    if number > 0 {
      fmt.Println("Number is greater than 0!")
    }

    fmt.Println(fileScanner.Text())
  }

  readFile.Close()
}
