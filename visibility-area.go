// simple hello-world code

package main

import "fmt"

var visibleVar string = "I'm visible variable!"

func main() {
  fmt.Println(visibleVar)
  f()
}

func f()  {
  fmt.Println(visibleVar)
}
