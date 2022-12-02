package main
 
import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "sort"
)
 
func main() {
  readFile, err := os.Open("input")
  
  if err != nil {
    fmt.Println(err)
  }
  fileScanner := bufio.NewScanner(readFile)

  fileScanner.Split(bufio.ScanLines)

  var sum int = 0
  //var gsum int = 0
  var output []int

  for fileScanner.Scan() {
    str := fileScanner.Text()
    if str == "" { 
      // if gsum < sum {
      //   gsum = sum
      //   fmt.Println("new biggest gsum: ", gsum)
      // }
      output = append(output, sum)
      sum = 0
      continue
    }

    num, err := strconv.Atoi(fileScanner.Text())

    if err != nil {
      fmt.Println("Error during conversion")
      return
    }

    sum += num
  }

  output = append(output, sum)
  sort.Ints(output)
  fmt.Println(output)

  // if gsum < sum {
  //   gsum = sum

  //   fmt.Println("new biggest gsum", gsum)
  // }

  readFile.Close()
}
