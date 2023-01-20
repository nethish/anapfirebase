package service

func Sum(number int) int {
  if number < 0 || number > 100 {
    return 0
  }

  ans := 0

  for i := 0; i <= number; i++ {
    ans += i
  }

  return ans
}
