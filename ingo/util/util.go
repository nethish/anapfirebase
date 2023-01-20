package util

import (
  "github.com/gin-gonic/gin"
  "strconv"
  "log"
)

func GetN(context *gin.Context) int {
  log.Println("Query params:", context.Request.URL.Query())
  query := context.Query("n")
  n, err := strconv.Atoi(query)
  if err != nil {
    log.Println("Cannot parse integer from: " + query)
  }
  return int(n)
}
