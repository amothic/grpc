package main

import (
	"github.com/amothic/grpc/service"
	"google.golang.org/grpc"
	pb "github.com/amothic/grpc/pb/cat"
	"log"
	"net"
)

func main() {
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	catService := &service.MyCatService{}
	pb.RegisterCatServer(server, catService)
	server.Serve(listenPort)
}
