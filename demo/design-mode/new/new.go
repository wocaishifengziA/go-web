package new

import "fmt"

// 实现一个服务器对象

type Service interface {
	Start()
	Stop()
}

type Server struct{}

func (s *Server) Start() {
	fmt.Println("server start!")
}

func (s *Server) Stop() {
	fmt.Println("server stop!")
}

func NewService() Service {
	return &Server{}
}
