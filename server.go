package main

import (
	"log"
	"net"

	"github.com/ogataka50/APIServerGo/service"
	"google.golang.org/grpc"
	"github.com/ogataka50/come-on-proto/go/protos"
)

func main() {

	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	service := service.PingService{}
	// 実行したい実処理をseverに登録する
	ping.RegisterPingServer(server, &service)
	server.Serve(listenPort)
}
