package main

import "github.com/gin-gonic/gin"

func main() {

  gin.SetMode(gin.ReleaseMode)
  router := gin.New()

  router.GET("/health", func(c *gin.Context) {
    c.JSON(200, "ok")
  })

  err := router.Run("127.0.0.1:8080")

  if err != nil {
    panic(err)
  }
}
