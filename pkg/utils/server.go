package utils

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type APIServer struct {
	httpServer *http.Server
}

func NewServer(addr string, router *gin.Engine) *APIServer {
	return &APIServer{
		httpServer: &http.Server{
			Addr:    ":" + addr,
			Handler: router,
		},
	}
}
func Start(s *APIServer) {
	if err := s.httpServer.ListenAndServe(); err != nil {
		log.Fatal("Server not started: ", err)
	}
}
