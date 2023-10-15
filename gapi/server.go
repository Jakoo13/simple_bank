package gapi

import (
	"fmt"

	db "github.com/jakoo13/simplebank/db/sqlc"
	"github.com/jakoo13/simplebank/pb"
	"github.com/jakoo13/simplebank/token"
	"github.com/jakoo13/simplebank/util"
)

// server serves gRPC requests for our banking service.
type Server struct {
	pb.SimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// Creates a new gRPC server.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}
