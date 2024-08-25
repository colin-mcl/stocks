package client_test

import (
	"os"
	"testing"

	"github.com/colin-mcl/stocks/pkg/v1/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var testClient client.StocksClient

func TestMain(m *testing.M) {
	serverURL := os.Getenv("STOCKS_URL")

	creds, err := credentials.NewClientTLSFromFile(os.Getenv("CERT_PATH"), "")
	if err != nil {
		panic(err)
	}

	conn, err := grpc.NewClient(serverURL, grpc.WithTransportCredentials(creds))

	testClient = *client.NewStocksClient(conn)

	os.Exit(m.Run())
}
