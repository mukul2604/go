package main
import (
  "fmt"
  "strconv"
  "math"
)

const s string = "constant"
// ~A
const MaxUint = ^uint(0)
const MaxUint32 = ^uint32(0)

//typedef similar to c
type Season uint8

// c's enumerator by GO's iota
const (
  Spring = Season(iota)
  Summer
  Autumn
  Winter
)

// func(var type) return var type
func (s Season) String() string {
  // somple array initialization
  name := []string{"spring", "summer","autumn","winter"}
  i := uint8(s)

  switch {
    //no need of break for each case
    case i <= uint8(Winter):
        return name[i]
    default:
        return strconv.Itoa(int(i))
  }
}

func main() {
  fmt.Println(s)
  const n =  50434343233
  const d = 3e20 / n
  fmt.Println(d)
  // numeric constant has no type until it is given one.
  //fmt.Println(int64(d))

  //math lib
  fmt.Println(math.Sin(d))

  s := Summer
  fmt.Println(s)
  s = Season(9)
  fmt.Println(s)
  fmt.Println(MaxUint)
  fmt.Println(MaxUint32)
}
