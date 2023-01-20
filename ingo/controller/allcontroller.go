package controller

import (
	"anapfirebase/ingo/model"
	"anapfirebase/ingo/service"
	"anapfirebase/ingo/util"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Allcontr(context *gin.Context) {
  n := util.GetN(context)

  fib := service.Fib(n)
  sum := service.Sum(n)
  parity := service.Parity(n)
  random := service.Random(n)

  var cs model.ComplexStruct
  cs.Fib = fib
  cs.Sum = sum
  cs.Parity = parity
  cs.Random = random
  cs.Number = n

  context.JSON(http.StatusOK, cs)
}
