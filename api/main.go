package main

import "github.com/gin-gonic/gin"

type SampleServer struct {
	ServerId int
}

func main() {
	r := gin.Default()
	server := &SampleServer{
		ServerId: 1,
	}
	RegisterHandlers(r, server)
	r.Run() // listen and serve on 0.0.0.0:8080
}
