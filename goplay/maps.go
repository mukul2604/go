package main
import "fmt"

func main() {
  // s := make([]string, 3) slice of length 3
  m := make(map[string]int)

  m["k1"] = 7
  m["k2"] = 10
  fmt.Println("map:", m)

  v1 := m["k1"]
  fmt.Println("v1: ", v1)
  fmt.Println("len:", len(m))

  delete(m, "k2")
  fmt.Println("map:", m)

  b,a := m["k2"]
  fmt.Println("prs:", b,a)

  n := map[string]int{"foo":1, "bar":2}
  fmt.Println("map:", n)
}
