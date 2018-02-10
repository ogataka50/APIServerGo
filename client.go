package main

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
	"github.com/ogataka50/APIServerGo/pb/proto"
	"golang.org/x/net/context"
)

func main() {
	//sampleなのでwithInsecure
	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := ping.NewPingClient(conn)
	req := ping.HelloReqest{ToMessage: "oooooops"}
	res, err := client.Hello(context.TODO(), &req)
	fmt.Printf("result:%#v \n", res)
	fmt.Printf("error::%#v \n", err)
}
