package main

import (
	"fmt"
  "os"
  "bufio"
  "strconv"
)

type Queue []rune

type StrInt struct {
  s string
  i int
  isNumber bool
}

func (q *Queue) Push(val rune) {
  *q = append(*q, val)
}

func (q *Queue) Pop() (*StrInt, error) {
  var strInt *StrInt
  if len(*q) == 0 {
    return strInt, fmt.Errorf("Queue is empty")
  }
  strInt = &StrInt{}
  val := (*q)[0]
  *q = (*q)[1:]
  strInt.s = string(val)
  integer, err := strconv.Atoi(strInt.s)
  if err != nil {
    strInt.isNumber = false
  } else {
    strInt.isNumber = true
  }
  strInt.i = integer
  return strInt, nil
}

func (q *Queue) Print() {
  var res = []string{}
  for i := 0; i < len(*q); i++ {
    res = append(res, string((*q)[i]))
  }
  fmt.Println(res)
}

func (q *Queue) Len() int {
  return len(*q)
}

type Node struct {
  values Queue
  next *Node
}

type NodeInteger struct {
  values []int 
  next *NodeInteger
  isInOrder bool
  isClosed bool
}

func ( n *NodeInteger)  NewNode() *NodeInteger {
  n.next = &NodeInteger{}
  return n.next
}

func CreateNode(n *NodeInteger) *NodeInteger {
  if n == nil {
    return &NodeInteger{}
  } else if n.isClosed {
    return n.NewNode()
  } else {
    return n
  }
}

type FileChars struct {
  left *Node 
  right *Node 
}

type SigTreeNode struct {
  left *NodeInteger
  right *NodeInteger
  next *SigTreeNode 
}

type SigTree struct {
  head *SigTreeNode 
}

func Compare(sigTree *SigTree) {
  sigTree := sigTree.head 
  for sigTree != nil {
    sigTree.left.isInOrder = true
    for i, _ := range sigTree.left.values {
      valL := sigTree.left.values[i]
      if i > len(sigTree.right.values) - 1 {
        continue
      }
      valR := sigTree.right.values[i]
      if valL > valR {
        sigTree.left.isInOrder = false
      }
    }
    sigTree = sigTree.next
  }
  fmt.Println("testing the result here")
}

func Part1() {
  readFile, err := os.Open("testinput")
  if err != nil {
    fmt.Println(err)
  }
  fileScanner := bufio.NewScanner(readFile)
  fileScanner.Split(bufio.ScanLines)
  lineN := 0
  var left *Node = &Node{}
  var right *Node = &Node{}
  fc := FileChars{
    left: left,
    right: right,
  }
  for fileScanner.Scan() {
    str := fileScanner.Text()
    if str == "" {
      continue
    }
    for _, ch := range str {
      if ch == ',' {
        continue
      }
      if lineN % 2 == 0 {
        left.values.Push(ch)
      } else {
        right.values.Push(ch)
      }
    }
    if lineN % 2 == 0 {
      left.next = &Node{}
      left = left.next
    } else {
      right.next = &Node{}
      right = right.next
    }
    lineN++
  }
  // now I want to go through fc and check all characters
  leftSide := fc.left
  rightSide := fc.right
  var sigTreeNode *SigTreeNode
  var sigTree *SigTree
  var currL *NodeInteger
  var currR *NodeInteger
  for leftSide != nil {
    var leftCh *StrInt
    var rightCh *StrInt
    currL = nil
    currR = nil
    for leftSide.values.Len() > 0 {
      leftCh, _ = leftSide.values.Pop()
      rightCh, _ = rightSide.values.Pop()

      if leftCh == nil || rightCh == nil {
        // currL.isClosed = true
        // currR.isClosed = true
        break
      }

      // both sides are integers
      if leftCh.isNumber && rightCh.isNumber {
        currL = CreateNode(currL)
        currL.values = append(currL.values, leftCh.i)
        currR = CreateNode(currR)
        currR.values = append(currR.values, rightCh.i)
      }

      // both sides are open brackets
      if leftCh.s == "[" && rightCh.s == "[" {
        currL = CreateNode(currL)
        currR = CreateNode(currR)
        if sigTree == nil {
          sigTree = &SigTreeNode{
            left: currL,
            right: currR,
          }
        }
      }

      // both sides are close brackets
      if leftCh.s == "]" && rightCh.s == "]" {
        currL.isClosed = true
        currR.isClosed = true
      }

      // one side is an open bracket the other an integer
      if leftCh.s == "[" && rightCh.isNumber {
        currL = CreateNode(currL)
        currR = CreateNode(currR)
        currR.values = append(currR.values, rightCh.i)
      }

      // one side is an open bracket the other an integer
      if leftCh.isNumber && rightCh.s == "[" {
        currL = CreateNode(currL)
        currR = CreateNode(currR)
        currL.values = append(currL.values, leftCh.i)
      }
      
      // one side is a close bracket and the other side has nothing
      if leftCh.s == "]" && rightCh == nil {
        currL.isClosed = true 
      }

      // one side is a close bracket and the other side has nothing
      if leftCh == nil && rightCh.s == "]" {
        currR.isClosed = true 
      }

      // one side is a close bracket and the other side a integer
      if leftCh.s == "]" && rightCh.isNumber {
        currL.isClosed = true 
        currR.values = append(currR.values, rightCh.i)
      }

      // one side is a close bracket and the other side a integer
      if leftCh.isNumber && rightCh.s == "]" {
        currR.isClosed = true 
        currL.values = append(currL.values, leftCh.i)
      }

      // one side is a close bracket and the other side is a open bracket 
      if leftCh.s == "]" && rightCh.s == "[" {
        currL.isClosed = true 
        currR = currR.NewNode()
      }
      
      // one side is a close bracket and the other side is a open bracket 
      if leftCh.s == "[" && rightCh.s == "]" {
        currR.isClosed = true 
        currL = currL.NewNode()
      }
    }
    if sigTree != nil {
      sigTree = sigTree.next
    }
    leftSide = leftSide.next
    rightSide = rightSide.next
  }
  Compare(sigTree)
}

func main() {
  Part1()
}
