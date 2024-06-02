package main

import (
	"log"
	"net"
	"os"

	"github.com/colin-mcl/stocks/controllers"
	"github.com/colin-mcl/stocks/gapi"
	"github.com/colin-mcl/stocks/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// launches the server to handle http or grpc requests for the stocks service
// can either use the GRPC server or gin server
func main() {

	errorLog := log.New(os.Stderr, "[ERROR]", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "[INFO]", log.Ldate|log.Ltime)

	runGrpcServer(errorLog, infoLog)
}

func runGrpcServer(errorLog *log.Logger, infoLog *log.Logger) {
	server, err := gapi.NewServer(errorLog, infoLog)
	if err != nil {
		errorLog.Fatal("cannot create server:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterStocksServer(grpcServer, server)

	// Allows a grpc client to explore which rpc calls are available
	reflection.Register(grpcServer)

	// creates a listener to listen for requests on localhost:9090
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		errorLog.Fatal("cannot create listener:", err)
	}

	infoLog.Printf("starting GRPC server at %s", listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil {
		errorLog.Fatal("cannot start grpc server:", err)
	}
}

func runGinServer() {
	router := gin.Default()
	router.GET("/tickers/:symbol", controllers.GetTicker)

	// Runs the server on localhost:8080 by default
	router.Run()
}
