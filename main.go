package main

import (
	"backEnd_Gin/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	fmt.Println("1")
	gin_server := gin.Default()
	gin_server = routes.CollectRoute(gin_server)
	err := gin_server.Run(":8401")
	if err != nil {
		log.Fatalln(err)
	}

}
