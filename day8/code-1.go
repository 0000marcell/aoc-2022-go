package main

import (
  "fmt"
  "os"
  "bufio"
  //"regexp"
  //"strings"
  "strconv"
)

func IsVisible(m [][]int, line int, column int, 
                       linescount int, columnscount int) bool {
  var itemheight int = m[line][column]
  fmt.Println("item height: ", itemheight)
  var visup bool = true
  for i := 1; line - i >= 0 ; i++ {
    fmt.Println("visup comparing to ", m[line - i][column], line - i)
    if itemheight <= m[line - i][column] {
      visup = false
      break
    }   
  }
  if visup {
    fmt.Println("visup is true")
    return true
  }
  var visdown bool = true
  for i := 1; line + i <= linescount ; i++ {
    fmt.Println("visdown accessing: ", line + i, column)
    fmt.Println("visdown comparing to ", m[line + i][column])
    if itemheight <= m[line + i][column] {
      visdown = false
      break
    }   
  }
  if visdown {
    fmt.Println("visdown is true")
    return true

  }
  var visleft bool = true

  for i := 1; column - i >= 0 ; i++ {
    fmt.Println("visleft comparing to ", m[line][column - i])
    if itemheight <= m[line][column - i] {
      visleft = false
      break
    }   
  }
  if visleft {
    fmt.Println("visleft is true")
    return true
  }
  var visright bool = true
  for i := 1; column + i <= columnscount ; i++ {
    fmt.Println("visright accesing ", line, column + i)
    fmt.Println("visright comparing to ", m[line][column + i])
    if itemheight <= m[line][column + i] {
      visright = false
      break
    }   
  }
  if visright {
    fmt.Println("visright is true")
    return true
  }
  fmt.Println("not true for: ", itemheight)
  fmt.Println("===========================")
  return false
}


func CountVisibleTrees(m [][]int, lastline int, lastcolumn int) int {
  var visible = 0
  for i := 0; i < len(m); i++ {
    if len(m[i]) > 0 && 
       i != 0 &&
       i != lastline {
      for j := 0; j < len(m[i]); j++ {
        if j != 0 &&
           j != lastcolumn {
          if IsVisible(m, i, j, lastline, lastcolumn) {
            visible += 1
          }
        }
      }
    }
  }
  return visible 
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
  fmt.Println("columns: ", columns)
  var lines int = CountMatLines(m)
  fmt.Println("lines: ", lines)
  var totaledge int = columns * 2 + lines * 2 - 4
  fmt.Println("totaledge: ", totaledge)
  var lastline int = lines - 1  
  var lastcolumn int = columns - 1
  var visibletreecount int = CountVisibleTrees(m, lastline, lastcolumn)
  fmt.Println(visibletreecount)
  var result int = totaledge + visibletreecount
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
