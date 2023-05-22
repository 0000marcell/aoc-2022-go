package main

import (
  "fmt"
  "strings"
  "strconv"
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

  var matches int = 0

  for fileScanner.Scan() {
    str := fileScanner.Text()
    strsplit := strings.Split(str, ",")
    var arrayint = [2]int{} 
    var arrayofarrays [2][len(arrayint)]int
    var arrnums [len(arrayofarrays)][]int
    for i := range strsplit {
      nums := strings.Split(strsplit[i], "-")
      f, err := strconv.Atoi(nums[0])
      if err != nil {
        fmt.Println(err)
      }
      s, err := strconv.Atoi(nums[1])
      if err != nil {
        fmt.Println(err)
      }
      arrnums[i] = append(arrnums[i], f, s)
    }
    fmt.Println(arrnums)
    if arrnums[0][0] < arrnums[1][0] {
      if arrnums[0][1] >= arrnums[1][1] {
        matches += 1
        fmt.Println(true)
      } 
    } else if arrnums[0][0] > arrnums[1][0]{
      if arrnums[1][1] >= arrnums[0][1] {
        matches += 1
        fmt.Println(true)
      }
    } else {
      if arrnums[0][1] >= arrnums[1][1] ||
         arrnums[1][1] >= arrnums[0][1] {
        matches += 1
        fmt.Println(true)
      }
    }
    fmt.Println("==================")
  }
  fmt.Println(matches)
  
  //   str := fileScanner.Text()
  //   var strLen int = len(str)
  //   var n int = strLen/2
  //   str1 := str[0:n]
  //   str2 := str[n:strLen]

  //   var validID = regexp.MustCompile("["+str1+"]")
  //   var char = validID.Find([]byte(str2))
  //   charN := int(big.NewInt(0).SetBytes(char).Uint64())
    
  //   if charN < 65 {
  //     fmt.Println("error char lower than 65")
  //     return 
  //   }
  //   if charN > 90 && charN < 97   {
  //     fmt.Println("error char between 90 and 97")
  //     return 
  //   }
  //   if charN > 122 {
  //     fmt.Println("error char greather than 122")
  //     return 
  //   }
  //   var points int = 0
  //   if charN > 64 {
  //     points = charN - 38 
  //   }
  //   if charN > 96 {
  //     points = charN - 96
  //   }
  //   totalscore += points
  // }
  // fmt.Println(totalscore)
}


// package main

// import (
//   "fmt"
//   "regexp"
//   "math/big"
//   "os"
//   "bufio"
// )

// func main() {
//   readFile, err := os.Open("testinput-1")
//   if err != nil {
//     fmt.Println(err)
//   }
//   fileScanner := bufio.NewScanner(readFile)

//   fileScanner.Split(bufio.ScanLines)

//   for fileScanner.Scan() {
//     str := fileScanner.Text()
//     var strLen int = len(str)
//     var n int = strLen/2
//     str1 := str[0:n]
//     str2 := str[n:strLen]

//     var validID = regexp.MustCompile("["+str1+"]")
//     var char = validID.Find([]byte(str2))
//     charN := int(big.NewInt(0).SetBytes(char).Uint64())
    
//     if charN < 65 {
//       fmt.Println("error char lower than 65")
//       return 
//     }
//     if charN > 90 && charN < 97   {
//       fmt.Println("error char between 90 and 97")
//       return 
//     }
//     if charN > 122 {
//       fmt.Println("error char greather than 122")
//       return 
//     }
//     var points int = 0
//     if charN > 64 {
//       points = charN - 38 
//     }
//     if charN > 96 {
//       points = charN - 96
//     }
//     totalscore += points
//   }
//   fmt.Println(totalscore)
// }
