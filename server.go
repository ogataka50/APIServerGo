package main

import (
	"log"
	"net"

	"github.com/ogataka50/APIServerGo/service"
	"google.golang.org/grpc"
	"github.com/ogataka50/APIServerGo/pb/proto"
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
