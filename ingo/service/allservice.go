package service

import (
	"anapfirebase/ingo/model"
)

func All(n int) model.ComplexStruct {
  fib := Fib(n)
  sum := Sum(n)
  parity := Parity(n)
  random := Random(n)

  var cs model.ComplexStruct
  cs.Fib = fib
  cs.Sum = sum
  cs.Parity = parity
  cs.Random = random
  cs.Number = n

  return cs
}
