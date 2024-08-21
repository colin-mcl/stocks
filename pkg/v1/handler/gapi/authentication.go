package gapi

import (
	"context"
	"fmt"

	"github.com/colin-mcl/stocks/internal/token"
	"google.golang.org/grpc/metadata"
)

const (
	authenticationHeader = "authentication"
)

// authenticateUser authenticates a user's credentials by verifying the access
// token passed in the metadata of the context and returns its payload if it is
// valid
func (server *Server) authenticateUser(ctx context.Context) (*token.Payload, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}

	values := md.Get(authenticationHeader)
	if len(values) == 0 {
		return nil, fmt.Errorf("missing authentication header")
	}

	accessToken := values[0]

	payload, err := server.tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("invalid access token: %s", accessToken)
	}

	return payload, nil
}
