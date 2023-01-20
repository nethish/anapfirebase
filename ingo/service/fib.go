package service

import (
  "log"
)

func Fib(number int) int {
  log.Println("Computing fib for n:", number)
  if number > 100 || number < 1 {
    return 0
  }

  x := 1
  y := 1

  for i:= 3; i <= number; i++ {
    t := y
    y = x + y
    x = t
  }

  return y
}
