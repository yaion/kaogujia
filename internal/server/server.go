package server

import (
	"goWebBasicProject/pkg/config"
)

type Server struct {
}

func NewServer(cfg *config.AppConfig) *Server {
	serve := new(Server)
	//todo 使用那个web框架

	return serve
}

func (serve *Server) Run() {

}
