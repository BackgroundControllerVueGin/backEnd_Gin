package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	fmt.Println("1")
	gin_server := gin.Default()
	err := gin_server.Run(":8401")
	if err != nil {
		log.Fatalln(err)
	}

}
