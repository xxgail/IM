package models

import (
	"errors"
	"fmt"
	"strings"
)

type Server struct {
	Ip   string `json:"ip"`
	Port string `json:"port"`
}

func NewServer(ip string, port string) *Server {
	return &Server{
		Ip:   ip,
		Port: port,
	}
}

func (s *Server) String() (str string) {
	if s == nil {
		return
	}

	str = fmt.Sprintf("%s:%s", s.Ip, s.Port)

	return
}

func StringToServer(str string) (server *Server, err error) {
	list := strings.Split(str, ":")
	if len(list) != 2 {
		return nil, errors.New("err")
	}

	server = &Server{
		Ip:   list[0],
		Port: list[1],
	}

	return
}
