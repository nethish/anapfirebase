package controller

import (
	"anapfirebase/ingo/service"
	"anapfirebase/ingo/util"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Allcontr(context *gin.Context) {
  n := util.GetN(context)

  cs := service.All(n)

  context.JSON(http.StatusOK, cs)
}
