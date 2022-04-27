package grpcwrapper

import (
	"fmt"
	"log"
	"net"

	flatbuffers "github.com/google/flatbuffers/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type Registry interface {
	RegisterServer(grpcServer *grpc.Server)
}

type server struct {
	grpcServer *grpc.Server
}

func New(registry Registry, kasp keepalive.ServerParameters, kaep keepalive.EnforcementPolicy) server {
	codec := &flatbuffers.FlatbuffersCodec{}
	s := server{
		// grpc.NewServer(grpc.KeepaliveEnforcementPolicy(kaep), grpc.KeepaliveParams(kasp), grpc.ForceServerCodec(codec)),
		grpc.NewServer(grpc.ForceServerCodec(codec)),
	}
	registry.RegisterServer(s.grpcServer)

	return s
}

func (s server) Run(server string, port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", server, port))
	if err != nil {
		log.Fatalf("Falied to listen: %v", err)
	}

	fmt.Println(fmt.Sprintf("listening on %s:%d ...", server, port))
	if err := s.grpcServer.Serve(lis); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
