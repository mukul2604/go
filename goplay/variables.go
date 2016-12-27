package main

import "fmt"

func main() {
  // var <varname>  <var type> = <value>
  var a string = "golang"
  fmt.Println(a)
  //multiple var declaration and initialization
  var b, c int = 1,2
  fmt.Println(b,c)
  // automatically infers the type
  var d = true
  fmt.Println(d)
  //uninitialized var is set to zero
  var e int
  fmt.Println(e)

  // equivalent to var f string = " shortcut"
  f := "shortcut"
  fmt.Println(f)
}
