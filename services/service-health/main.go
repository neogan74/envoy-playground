package main 

func main() {
  r := git.Default()

  r.GET("/health", func(c *gin.Context){
    c.JSON(http.StatusOK, gin.H{
      "status":"ok",
    })
  })

  r.RUN(":8081")
}
