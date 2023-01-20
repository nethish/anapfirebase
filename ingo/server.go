package main

import (
  "github.com/gin-gonic/gin"
  "log"
  "anapfirebase/ingo/controller"
)

func Serve() {
  log.Println("Starting the server")

  router := gin.Default()

  routes := router.Group("/compute")
  routes.GET("/fib", controller.Fibcontr)
  routes.GET("/parity", controller.Paritycontr)
  routes.GET("/sum", controller.Sumcontr)
  routes.GET("/random", controller.Randomcontr)
  routes.GET("/computeall", controller.Allcontr)

  log.Println("Controllers setup done. Starting server at port 8000");

  router.Run(":8000")

  log.Println("Server started")

}
