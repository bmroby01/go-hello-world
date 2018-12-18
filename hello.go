package main

import "fmt"
import "math"

const s string = "constant"

func main() {
  fmt.Println("Hello, World!")

  fmt.Println("Bri" + "anna")

  fmt.Println("1+1=", 1+1)

  fmt.Println(true && false)
  fmt.Println(true || false)

  var a = "initial"
  fmt.Println(a)

  var b, c int = 1, 2
  fmt.Println(b, c)

  var d = true
  fmt.Println(d)

  var e int
  fmt.Println(e)

  f := "short"
  fmt.Println(f)

  fmt.Println(s)

  const n = 500000000

  const number = 3e20 / n
  fmt.Println(number)

  fmt.Println(int64(number))

  fmt.Println(math.Sin(n))
}
