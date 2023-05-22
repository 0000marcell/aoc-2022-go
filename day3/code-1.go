package main

import (
  "fmt"
  "regexp"
  "math/big"
  "os"
  "bufio"
)

func main() {
  readFile, err := os.Open("input-1")
  if err != nil {
    fmt.Println(err)
  }
  fileScanner := bufio.NewScanner(readFile)

  fileScanner.Split(bufio.ScanLines)

  var totalscore int = 0

  for fileScanner.Scan() {
    // split string in half
    str := fileScanner.Text()
    var strLen int = len(str)
    var n int = strLen/2
    str1 := str[0:n]
    str2 := str[n:strLen]

    // get the duplicate item
    var validID = regexp.MustCompile("["+str1+"]")
    var char = validID.Find([]byte(str2))
    //charN, err := strconv.Atoi(string(char))
    charN := int(big.NewInt(0).SetBytes(char).Uint64())
    //fmt.Println(charN)
    // if err != nil {
    //   fmt.Println("Error during conversion")
    //   return
    // }
    
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
    // fmt.Println(points)
    // var result = string(char)
    // fmt.Println(result)
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
