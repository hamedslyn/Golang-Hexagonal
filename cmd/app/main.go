package main

import (
	"github.com/hamedslyn/heli-todo/pkg/config"
	"github.com/hamedslyn/heli-todo/pkg/server"
)

func main() {
	cfg := config.MustLoad()
	s := server.NewServer(cfg)
	s.Run()
}
