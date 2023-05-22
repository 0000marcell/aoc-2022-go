package main

import (
  "fmt"
  "os"
  "bufio"
  //"regexp"
  //"strings"
  "strconv"
)

func VisibleScore(m [][]int, line int, column int, 
                       linescount int, columnscount int) int {
  var itemheight int = m[line][column]

  var vupscore int = 0 
  for i := 1; line - i >= 0 ; i++ {
    vupscore += 1
    if itemheight <= m[line - i][column] {
      break
    }   
  }

  var vdownscore int = 0 
  for i := 1; line + i <= linescount ; i++ {
    vdownscore += 1
    if itemheight <= m[line + i][column] {
      break
    }   
  }

  var vleftscore int = 0 
  for i := 1; column - i >= 0 ; i++ {
    vleftscore += 1
    if itemheight <= m[line][column - i] {
      break
    }   
  }
  
  var vrightscore int = 0 
  for i := 1; column + i <= columnscount ; i++ {
    vrightscore += 1
    if itemheight <= m[line][column + i] {
      break
    }   
  }
  
  var total int = vupscore * vdownscore * vleftscore * vrightscore 
  return total 
}


func GetBiggestVisibleScore(m [][]int, lastline int, 
                            lastcolumn int) int {
  var biggestscore int = 0
  for i := 0; i < len(m); i++ {
    if len(m[i]) > 0 && 
       i != 0 &&
       i != lastline {
      for j := 0; j < len(m[i]); j++ {
        if j != 0 &&
           j != lastcolumn {
          
          var score int = VisibleScore(m, i, j, lastline, lastcolumn)
          if score > biggestscore {
            biggestscore = score
          }
        }
      }
    }
  }
  return biggestscore 
}

func CountMatLines(m [][]int) int {
  var linescount int = 0
  for i := 0; i < len(m); i++ {
    if len(m[i]) > 0 {
      linescount += 1
      // for j := 0; j < len(m[i]); j++ {
      //   fmt.Println(m[i][j])
      // }
      //fmt.Println("=============")
    }
  }
  return linescount 
}

func main() {
  //var n1 Node 
  readFile, err := os.Open("input-1")
  if err != nil {
    fmt.Println(err)
  }
  fileScanner := bufio.NewScanner(readFile)
  // var cdreg = regexp.MustCompile(` cd (\w)+`)
  // var dirreg = regexp.MustCompile("dir .*")
  // var filereg = regexp.MustCompile(`\d+ .*`)
  // var root Node
  // var curnode* Node
  m := make([][]int, 99)
  var lcount int = 0
  for fileScanner.Scan() {
    str := fileScanner.Text()
    for _, ch := range str {
      num, err := strconv.Atoi(string(ch))
      if err != nil {
        fmt.Println("Error during conversion")
        return
      }
      m[lcount] = append(m[lcount], num)
	  }
    //fmt.Println("====================")
    lcount += 1
  }
  fmt.Println(m)
  var columns int = len(m[0])
  var lines int = CountMatLines(m)
  var lastline int = lines - 1  
  var lastcolumn int = columns - 1

  var result int = GetBiggestVisibleScore(m, lastline, lastcolumn)
  fmt.Println("final result: ", result)
  // how can I cound the edges?
  // [0][anything] is a edge
  // [last line][anything] is another edge
  // [anything][0] is a edge
  // [anything][lastcolumn] is a edge
  //fmt.Printf("%+v\n", root)
  // SetSizes(&root) 
  // ShowSizes(&root)
  // FindDirSizes(&root)
  // fmt.Println("totalsum: ", totalsum)
}
