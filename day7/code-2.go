package main

import (
  "fmt"
  "os"
  "bufio"
  "regexp"
  "strings"
  "strconv"
)

type Node struct {
  name string
  kind string
  size int
  parent *Node
  children []*Node
}

func SetSizes (node *Node) int {
  var size = 0 
  size += node.size
  if len(node.children) > 0 {
    for i := 0; i < len(node.children); i++ {
      size += SetSizes(node.children[i])
    }
  }
  node.size = size
  return size
}

func ShowSizes(node *Node) {
  if node.kind == "dir" {
    //fmt.Println("Node: ", node.name, node.size)
  }
  if len(node.children) > 0 {
    for i := 0; i < len(node.children); i++ {
      ShowSizes(node.children[i])
    }
  }
}

var totalsum int = 0
func FindDirSizes(node *Node) {
  if node.kind == "dir" {
    if node.size <= 100000 {
      totalsum += node.size
    }
  }
  if len(node.children) > 0 {
    for i := 0; i < len(node.children); i++ {
      FindDirSizes(node.children[i])
    }
  }
}

var spacetofree int = 0
var smallestdir int = 999999999999999999
func FreeSpace(node *Node) {
  if node.kind == "dir" {
    if node.size >= spacetofree {
      if node.size < smallestdir  {
        smallestdir = node.size
      }
    }
  }
  if len(node.children) > 0 {
    for i := 0; i < len(node.children); i++ {
      FreeSpace(node.children[i])
    }
  }
}




func main() {
  //var n1 Node 
  readFile, err := os.Open("input-1")
  if err != nil {
    fmt.Println(err)
  }
  fileScanner := bufio.NewScanner(readFile)
  var cdreg = regexp.MustCompile(` cd (\w)+`)
  var dirreg = regexp.MustCompile("dir .*")
  var filereg = regexp.MustCompile(`\d+ .*`)
  var root Node
  var curnode* Node
  for fileScanner.Scan() {
    str := fileScanner.Text()
    if str == "$ cd .." {
      curnode = curnode.parent
    } else if str == "$ cd /" {
      strsplit := strings.Split(str, " ")
      var foldername string = strsplit[2]
      root.name = foldername
      root.kind = "dir"
      curnode = &root
    } else if cdreg.Match([]byte(str)) {
      strsplit := strings.Split(str, " ")
      var foldername string = strsplit[2]
      // can I potentially recreated the same node
      // I gonna assume we don't cd into the same folder two times
      var nnode Node
      nnode.name = foldername
      nnode.kind = "dir"
      nnode.parent = curnode
      curnode.children  = append(curnode.children, &nnode)
      curnode = &nnode
    }
    if str == "$ ls" {
    }

    if dirreg.Match([]byte(str)) {
    }

    if filereg.Match([]byte(str)) {
      strsplit := strings.Split(str, " ")
      var nnode Node
      nnode.name =  strsplit[1]
      size, err := strconv.Atoi(string(strsplit[0]))
      if err != nil {
        fmt.Println("Error during conversion")
        return
      }
      nnode.size = size 
      nnode.kind = "file"
      nnode.parent = curnode
      curnode.children  = append(curnode.children, &nnode)
    }
    //fmt.Println("====================")
  }
  //fmt.Printf("%+v\n", root)
  SetSizes(&root) 
  ShowSizes(&root)
  var unusedspace int = 70000000 - root.size
  spacetofree = 30000000 - unusedspace
  FreeSpace(&root) 
  fmt.Println("smallestdir to size: ", smallestdir)
}
