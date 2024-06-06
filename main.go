package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/colin-mcl/stocks/controllers"
	"github.com/colin-mcl/stocks/gapi"
	"github.com/colin-mcl/stocks/internal/models"
	"github.com/colin-mcl/stocks/pb"
	"github.com/colin-mcl/stocks/util"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// launches the server to handle http or grpc requests for the stocks service
// can either use the GRPC server or gin server
func main() {

	errorLog := log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLog := log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime)

	// open mysql database connection and check for errors
	// TODO: change this to not be hard coded
	db, err := util.OpenDB("web:Amsterdam22!@/stocks?parseTime=true")

	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	err = runGrpcServer(db, errorLog, infoLog)

	if err != nil {
		errorLog.Fatal(err)
	}
}

func runGrpcServer(db *sql.DB, errorLog *log.Logger, infoLog *log.Logger) error {

	server, err := gapi.NewServer(&models.UserModel{DB: db}, errorLog, infoLog)
	if err != nil {
		return fmt.Errorf("failed to create server:%w", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterStocksServer(grpcServer, server)

	// Allows a grpc client to explore which rpc calls are available
	reflection.Register(grpcServer)

	// creates a listener to listen for requests on localhost:9090
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		return fmt.Errorf("failed to create listener: %w", err)
	}

	infoLog.Printf("starting GRPC server at %s", listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil {
		return fmt.Errorf("failed to start grpc server: %w", err)
	}

	return nil
}

func runGinServer() {
	router := gin.Default()
	router.GET("/tickers/:symbol", controllers.GetTicker)

	// Runs the server on localhost:8080 by default
	router.Run()
}
