package main

import (
	"log"
	"net"
	"net/rpc"

	"./../shared"
)

func registerGitAPI(server *rpc.Server, git shared.Git) {
	server.RegisterName("GitAPI", git)
}

func main() {
	git := new(Git)

	server := rpc.NewServer()
	registerGitAPI(server, git)

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	server.Accept(l)
}
