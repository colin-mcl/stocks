package main

import (
	"log"
	"net"

	"github.com/colin-mcl/stocks/controllers"
	"github.com/colin-mcl/stocks/gapi"
	"github.com/colin-mcl/stocks/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	runGrpcServer()
}

func runGrpcServer() {
	server, err := gapi.NewServer()
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterStocksServer(grpcServer, server)

	// Allows a grpc client to explore which rpc calls are available
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", "localhost:9090")
	if err != nil {
		log.Fatal("cannot create listener:", err)
	}

	log.Printf("startGRPC server at %s", listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start grpc server:", err)
	}
}

func runGinServer() {
	router := gin.Default()
	router.GET("/tickers/:symbol", controllers.GetTicker)

	// Runs the server on localhost:8080 by default
	router.Run()
}
