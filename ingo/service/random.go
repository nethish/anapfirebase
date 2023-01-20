package service

import (
  "math/rand"
)

func Random(number int) string {
  ans := ""

  const charset = "nethish"

  for i := 0; i < number; i++ {
    ans += string(charset[rand.Int() % len(charset)])
  }

  return ans
}
