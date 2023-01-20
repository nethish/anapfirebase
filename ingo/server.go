package main

import (
  "github.com/gin-gonic/gin"
  "log"
  "anapfirebase/ingo/controller"
)

func Serve() {
  log.Println("Starting the server")

  router := gin.Default()

  computeRoutes := router.Group("/compute")
  computeRoutes.GET("/fib", controller.Fibcontr)
  computeRoutes.GET("/parity", controller.Paritycontr)
  computeRoutes.GET("/sum", controller.Sumcontr)
  computeRoutes.GET("/random", controller.Randomcontr)
  computeRoutes.GET("/computeall", controller.Allcontr)

  databaseRoutes := router.Group("/database")
  databaseRoutes.GET("/put", controller.DatabasePut)
  databaseRoutes.GET("/get", controller.DatabaseGet)

  log.Println("Controllers setup done. Starting server at port 8000");

  router.Run(":8000")

  log.Println("Server started")

}
