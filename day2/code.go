package main
 
import (
  "bufio"
  "fmt"
  "os"
)
 
func main() {
  readFile, err := os.Open("input")
  
  if err != nil {
    fmt.Println(err)
  }
  fileScanner := bufio.NewScanner(readFile)

  fileScanner.Split(bufio.ScanLines)

  var totalscore int = 0

  for fileScanner.Scan() {
    str := fileScanner.Text()
    //s := strings.Split(str, " ")
    if str == "A X" {
      totalscore += 0 + 3
    } else if str == "A Y" {
      totalscore += 3 + 1
    } else if str == "A Z" {
      totalscore += 6 + 2
    } else if str == "B X" {
      totalscore += 0 + 1
    } else if str == "B Y" {
      totalscore += 3 + 2
    } else if str == "B Z" {
      totalscore += 6 + 3
    } else if str == "C X" {
      totalscore += 0 + 2
    } else if str == "C Y" {
      totalscore += 3 + 3
    } else if str == "C Z" {
      totalscore += 6 + 1
    } else {
      fmt.Println("ERROR!!!")
    }
  }

  fmt.Println(totalscore)

  // output = append(output, sum)
  // sort.Ints(output)
  // fmt.Println(output)
  readFile.Close()
}
