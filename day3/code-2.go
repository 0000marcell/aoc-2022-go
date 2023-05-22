package main

import (
  "fmt"
  "regexp"
  "math/big"
  "os"
  "bufio"
  "strings"
)

func main() {
  readFile, err := os.Open("input-2")
  if err != nil {
    fmt.Println(err)
  }
  fileScanner := bufio.NewScanner(readFile)

  fileScanner.Split(bufio.ScanLines)

  var totalscore int = 0

  var linecount int = 1 

  var prevline string = ""
  var partialR string = ""

  for fileScanner.Scan() {
    
    str := fileScanner.Text()

    if linecount % 3 == 0 {
      var validID = regexp.MustCompile("["+str+"]")
      var char = validID.Find([]byte(partialR))
      charN := int(big.NewInt(0).SetBytes(char).Uint64())
      if charN < 65 {
        fmt.Println("error char lower than 65")
        return 
      }
      if charN > 90 && charN < 97   {
        fmt.Println("error char between 90 and 97")
        return 
      }
      if charN > 122 {
        fmt.Println("error char greather than 122")
        return 
      }
      var points int = 0
      if charN > 64 {
        points = charN - 38 
      }
      if charN > 96 {
        points = charN - 96
      }
      totalscore += points
    } else {
      if prevline != "" {
        r := regexp.MustCompile("["+str+"]")
        matches := r.FindAllString(prevline, -1)
        partialR = strings.Join(matches[:], "")
        prevline = ""
      } else {
        prevline = str
      }
    }
    linecount += 1
  }
  fmt.Println(totalscore)
}

// package main

// import (
//   "fmt"
//   "regexp"
// )

// func main() {

  // how to compare two strings
  //var str2 = "vJrwpWtwJgWrhcsFMMfFFhFp"
  // var str1 = "jqHRNqRjqzjGDLGL" 
  // var validID = regexp.MustCompile("["+str1+"]")

  // var str2 = "rsFMfFZSrLrFZsSL"

  // fmt.Printf("%q\n", validID.Find([]byte(str2)))
//}

// package main

// import (
// 	"fmt"
// 	"regexp"
// )

// func main() {
// 	re := regexp.MustCompile(`foo.?`)
// 	fmt.Printf("%q\n", re.Find([]byte(`seafood fool`)))

// }
