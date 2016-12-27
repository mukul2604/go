package main

import "fmt"

func main() {
  res := mul(1,2)
  fmt.Println("1+2=", res)
  s,m := sumandmul(3,4)
  fmt.Println("sum:",s,"mul:",m)
  p,_ := sumandmul(3,4)
  fmt.Println("sum:",p)
  sum(1,2)
  sum(1,2,3)
  nums := []int{1,2,3,4}
  sum(nums...)
}

func mul(a int, b int) int {
  return a*b
}

// multiple return
func  sumandmul(a int, b int) (int, int) {
    return  a+b, a*b
}

//variadic function
func sum(nums ...int) {
  fmt.Print(nums," ")
  total := 0
  for _, num := range nums {
    total += num
  }
  fmt.Println(total)
}


 
