package controller

import (
	"anapfirebase/ingo/model"
	"anapfirebase/ingo/constants"
	"anapfirebase/ingo/util"
	"anapfirebase/ingo/firebase"
	"net/http"
	"github.com/gin-gonic/gin"
  "log"
)

func DatabaseGet(context *gin.Context) {
  n := context.Query("n")

  client, ctx := firebase.NewFirestoreClient()

  states := client.Collection(constants.AnapCollection)
  doc := states.Doc(n)

  docsnap, err := doc.Get(ctx)

  if !(docsnap.Exists()) {
    context.JSON(http.StatusNotFound, "Document not found")
    return;
  }
  util.HandleError(err)
  
  log.Println("Successfully got the doc")
  var cs model.ComplexStruct

  docsnap.DataTo(&cs)

  context.JSON(http.StatusOK, cs)
}
