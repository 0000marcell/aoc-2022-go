package main

import (
	"fmt"
)

func main() {
  // remove last item from matrix
  m := make([][]string, 10)
  m[0] = append(m[0], "")
  m[0][0] = "A"
  fmt.Println(m[0][0])
	////str1 := "this is a [sample] [[string]] with [SOME] special words"
  //str1 := "asbasdkjfsajfd bnaksdjf"

	//re := regexp.MustCompile(`\[`)
  //if re.Match([]byte(str1)) {
    //fmt.Println("it does match")
  //} else {
    //fmt.Println("it does not match")
  //}
}
