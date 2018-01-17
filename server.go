package main

import (
	"log"
	"net"

	pb "github.com/ogataka50/APIServerGo/pb/proto"
	"github.com/ogataka50/APIServerGo/service"
	"google.golang.org/grpc"
)

func main() {

	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	catService := &service.MyCatService{}
	// 実行したい実処理をseverに登録する
	pb.RegisterCatServer(server, catService)
	server.Serve(listenPort)
}
