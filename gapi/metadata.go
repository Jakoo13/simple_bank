package gapi

import (
	"context"
	"log"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
) 

const (
	grpcUserAgentHeader = "user-agent"
	xFowardedForHeader = "x-forwarded-for"
)

type Metadata struct {
	UserAgent string
	ClientIP string
}

func (server *Server) extractMetadata(ctx context.Context) *Metadata{
	mtdt := &Metadata{}

	if md, ok := metadata.FromIncomingContext(ctx); ok {
		log.Printf("metadata: %v", md)
		if userAgents := md.Get(grpcUserAgentHeader); len(userAgents) > 0 {
			mtdt.UserAgent = userAgents[0]
		}
	}

	if peerInfo, ok := peer.FromContext(ctx); ok {
		mtdt.ClientIP = peerInfo.Addr.String()
	}

	return mtdt
}