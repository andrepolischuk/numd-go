package main

import (
  "fmt"
  "math/rand"
  ".."
)

func main() {
  for i := 0; i < 25; i++ {
    fmt.Println(numd.Decline(rand.Intn(100), "рубль", "рубля", "рублей"))
  }
}
