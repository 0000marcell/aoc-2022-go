package main

import (
	"fmt"
)

func main() {
  // remove last item from matrix
  image := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
  fmt.Println(len(image))
  var result int = len(image)/40 * 40
  fmt.Println(result)
	////str1 := "this is a [sample] [[string]] with [SOME] special words"
  //str1 := "asbasdkjfsajfd bnaksdjf"

	//re := regexp.MustCompile(`\[`)
  //if re.Match([]byte(str1)) {
    //fmt.Println("it does match")
  //} else {
    //fmt.Println("it does not match")
  //}
}
