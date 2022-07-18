package grpc_server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/programzheng/games/internal/grpc/proto"
	"github.com/programzheng/games/internal/service"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement message.GreeterServer.
type Server struct {
	pb.UnimplementedGreeterServer
}

// RandomTicket implements message.GreeterServer
func (s *Server) RandomTicket(ctx context.Context, in *pb.RandomTicketRequest) (*pb.RandomTicketResponse, error) {
	count := int(in.GetCount())
	service.IssuedRandomTickets(count)
	return &pb.RandomTicketResponse{Message: "success"}, nil
}

func Run() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &Server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
