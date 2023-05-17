package gapi

import (
	"fmt"

	db "github.com/Huy1996/simplebank/db/sqlc"
	"github.com/Huy1996/simplebank/pb"
	"github.com/Huy1996/simplebank/token"
	"github.com/Huy1996/simplebank/util"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config 		util.Config
	store 		db.Store
	tokenMaker 	token.Maker
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:		config,
		store: 		store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}


