package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-messanger/config"
	"go-messanger/server/http/handler"
)

type Server struct {
	accountHandler handler.AccountHandler
	cfg            config.HttpConfig
	engine         *gin.Engine
}

func NewServer(cfg config.HttpConfig, accountHandler handler.AccountHandler) Server {
	return Server{
		accountHandler: accountHandler,
		cfg:            cfg,
	}
}

func (s *Server) Run() error {
	r := gin.Default()

	r.POST("/api/v1/register", s.accountHandler.HandleRegistration)

	if err := r.Run(fmt.Sprintf(":%d", s.cfg.Port)); err != nil {
		return err
	}

	s.engine = r
	return nil
}
