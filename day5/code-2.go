package main

import (
  "fmt"
  "strings"
  "regexp"
  "os"
  "bufio"
  "strconv"
)

func main() {
  readFile, err := os.Open("input-1")
  if err != nil {
    fmt.Println(err)
  }
  fileScanner := bufio.NewScanner(readFile)

  fileScanner.Split(bufio.ScanLines)
  var cratestomove int = 0
  var fromc int = 0
  var toc int = 0
	iscrate := regexp.MustCompile(`\[`)
  m := make([][]string, 10)
  for fileScanner.Scan() {
    str := fileScanner.Text()
    strsplit := strings.Split(str, " ")
    //fmt.Println(strsplit[0])
    var blankchars int = 0
    var startingmove int = 0 
    var ccount int = 0
    for _, e := range strsplit {
      // handling matrix creation
      if e != "" {
        if iscrate.Match([]byte(e)) {
          m[ccount] = append(m[ccount], e)
          ccount += 1
          blankchars = 0
        }
      }else {
          // if blankchars == 8 {
          //   m[lcount] = append(m[lcount], "", "")
          // }
        blankchars += 1
        if blankchars == 4 {
          blankchars = 0
          ccount += 1
        }
      }

      if e == "move" {
        startingmove = 1
      }
      if startingmove == 1 {
        n, err := strconv.Atoi(e)
        // check if it's a number
        if err == nil {
          if cratestomove == 0 {
            cratestomove = n
          } else if fromc == 0 {
            fromc = n
          } else if toc == 0 {
            toc = n
          }
        }
      }
      if cratestomove != 0 && fromc != 0 && toc != 0 {
        fmt.Println("cratestomove: ", cratestomove)
        fmt.Println("fromc: ", fromc)
        fmt.Println("toc: ", toc)
        //var valuetomove string = ""
        //valuetomove = m[fromc - 1][cratestomove - 1 - i]
        //x = append([]int{1}, x...)
        fmt.Println("before move: ", m)
        itemstoappend := []string{}
        for i := 0; i < cratestomove; i++ {
          itemstoappend = append(itemstoappend, m[fromc - 1][i])
        }
        m[fromc - 1] = m[fromc - 1][cratestomove:]
        fmt.Println("itemstoappend: ", itemstoappend)
        m[toc - 1] = append(itemstoappend, m[toc - 1]...)
        //m[fromc - 1] = m[fromc - 1][cratestomove:]
        //fmt.Println("value to move: ", valuetomove)
        //for i := 0; i < cratestomove; i++ {
        //  //valuetomove = m[fromc - 1][cratestomove - 1 - i]
        //  m[fromc - 1] = m[fromc - 1][1:]
        //  //x = append([]int{1}, x...)

        //  //m[toc - 1] = append([]string{valuetomove}, m[toc - 1]...)
        //  fmt.Println("value to move: ", valuetomove)
        //}
        startingmove = 0
        cratestomove = 0
        fromc = 0
        toc = 0
        fmt.Println("after move: ", m)
      }
    }
    fmt.Println("================")
  }

  // first column is [x][0]
  // second column is [x][1]
  // third column is [x][2]

  // first column 
  for i := range m[0] {
    fmt.Println("column 1", m[0][i]) // empty 
  }

  for i := range m[1] {
    fmt.Println("column 2", m[1][i]) // empty 
  }

  for i := range m[2] {
    fmt.Println("column 3", m[2][i]) // empty 
  }
  fmt.Println("final result is: ", m[0][0], m[1][0], m[2][0], m[3][0], m[4][0], m[5][0], m[6][0], m[7][0], m[8][0])
  //fmt.Println("final result is: ", m[0][0], m[1][0], m[2][0])
  // fmt.Println("0 0", m[0][0]) // empty 
  // fmt.Println("0 1", m[0][1]) // N
  // fmt.Println("0 2", m[0][2]) // Z

  // fmt.Println("1 0", m[1][0]) // D
  // fmt.Println("1 1", m[1][1]) // C
  // fmt.Println("1 2", m[1][2]) // M

  // fmt.Println("2 0", m[2][0]) // empty
  // fmt.Println("2 1", m[2][1]) // empty
  // fmt.Println("2 2", m[2][2]) // P

  //fmt.Println(matches)
}
