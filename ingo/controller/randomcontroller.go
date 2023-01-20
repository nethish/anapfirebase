package controller

import (
  "github.com/gin-gonic/gin"
  "anapfirebase/ingo/service"
  "net/http"
  "anapfirebase/ingo/util"
)

func Randomcontr(context *gin.Context) {
  n := util.GetN(context)
  ans := service.Random(int(n))
  context.JSON(http.StatusOK, ans)
}
