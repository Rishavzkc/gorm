package routes

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	server *gin.Engine
}

func NewServer() Server {
	return Server{

		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := ConfigRoutes(s.server)

	log.Fatal(router.Run(":8080"))

}
