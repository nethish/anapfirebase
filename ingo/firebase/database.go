package firebase

import (
	"anapfirebase/ingo/constants"
	"anapfirebase/ingo/util"
	"context"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
  "log"
)

func NewFirestoreClient() (*firestore.Client, context.Context) {
  log.Println("Getting client")

  ctx := context.Background()
  opt := option.WithCredentialsFile(constants.CredentialsFilePath)
  client, err := firestore.NewClient(ctx, constants.ProjectID, opt)
  util.HandleError(err)

  log.Println("Firestore client connection success")

  return client, ctx
}
