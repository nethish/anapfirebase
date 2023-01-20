package controller

import (
	"anapfirebase/ingo/constants"
	"anapfirebase/ingo/firebase"
	"anapfirebase/ingo/service"
	"anapfirebase/ingo/util"
	"net/http"
	"github.com/gin-gonic/gin"
  "log"
)

func DatabasePut(context *gin.Context) {
  n := util.GetN(context)
  cs := service.All(int(n))

  client, ctx := firebase.NewFirestoreClient()

  states := client.Collection(constants.AnapCollection)
  ca := states.Doc(context.Query("n"))
  _, err := ca.Set(ctx, cs)
  util.HandleError(err)

  log.Println("Write success")

  context.JSON(http.StatusOK, cs)
}
