package main
import ("fmt")

func main() {
  a := []int{2, 3, 5, 8, 6}
  for i:=0; i < 5; i++ {
    if a[i] < a[i + 1] {

    }
    fmt.Println(a[i])
  }
}
