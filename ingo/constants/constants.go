package constants

import (
  "github.com/joho/godotenv"
  "anapfirebase/ingo/util"
  "os"
  "log"
)

var credential_path string = "CREDENTIALS_PATH"
var project_id string = "PROJECT_ID"
var anap_collection string = "ANAP_COLLECTION"

var (
  DatabaseURL string
  ProjectID string
  CredentialsFilePath string
  AnapCollection string
)

func init() {
  err := godotenv.Load()
  util.HandleError(err)

  CredentialsFilePath = os.Getenv(credential_path)
  ProjectID = os.Getenv(project_id)
  AnapCollection = os.Getenv(anap_collection)

  log.Println("CredentialsFilePath: " + CredentialsFilePath)
  log.Println("ProjectID: " + ProjectID)
}
