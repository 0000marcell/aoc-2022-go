package main

import (
	"fmt"
   "reflect"
)

func ParseList(list []rune) ([]any, []rune) {
  var packets []any
  var rpackets []any

  out:
  for {
    switch list[0] {
    case '[':
      rpackets, list = ParseList(list[1:])
      packets = append(packets, rpackets)
    case ',':
    case ']':
      break out
    default:
      packets = append(packets, int(list[0]))
    }
    list = list[1:]
    if len(list) == 0 {
      break
    }
  }
  return packets, list
}

func isSlice(x interface{}) bool {
  t := reflect.TypeOf(x)

  if t.Kind() == reflect.Slice {
    return true
  }
  return false
}

func CompareWithRight(lelem interface{}, right interface{}, leftIsNumber bool) {
  switch right := right.(type) {
  case []interface{}:
    if leftIsNumber {
      arr := []any{}
      arr = append(arr, relem)
      Compare(lelem, arr)
    } else {
      Compare(lelem, arr)
    }
  case 2:
  default:
  }
}

func Compare(left interface{}, right interface{}) bool {
  switch left := left.(type) {
  // left is a slice
  case []interface{}:
    for _, lelem := range left {
      switch lelem := lelem.(type)  {
      case []interface:
        CompareWithRight(lelem, right, false)
      default:
        CompareWithRight(lelem, right, true)
      }
    }
  default:
    CompareWithRight(left, right, true)
  }

  return true
}

func Parse() {
  //list1 := []rune{'[','1','1','3','1','1',']'}
  list1 := []rune{'[','1','[','2','[','3','[','4','[','5','6','7',']',']',']',']','8','9',']'}
  //list2 := []rune{'[','1','[','2','[','3','[','4','[','5','6','7',']',']',']',']','8','9',']'}
  left, _ := ParseList(list1[1:])
  //right, _ := ParseList(list2[1:])
  Compare(left)
}

func main() {
  Parse()
}
