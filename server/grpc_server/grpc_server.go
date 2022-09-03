package grpc_server

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/programzheng/games/internal/grpc/proto"
	"github.com/programzheng/games/internal/service"
	"github.com/programzheng/games/pkg/helper"
	"google.golang.org/grpc"
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

// AssignRandomIssuedTicketToThirdPartyUser implements message.GreeterServer
func (s *Server) AssignRandomIssuedTicketToThirdPartyUser(ctx context.Context, in *pb.AssignRandomIssuedTicketToThirdPartyUserRequest) (*pb.AssignRandomIssuedTicketToThirdPartyUserResponse, error) {
	agentCode := in.GetCode()
	thirdPartyID := in.GetThirdPartyID()

	response, err := service.PlayAssignTicketForThirdPartyUser(agentCode, thirdPartyID)
	if err != nil {
		return nil, err
	}

	return &pb.AssignRandomIssuedTicketToThirdPartyUserResponse{
		UserTicket: &pb.UserTicket{
			Code: response.Code,
			Name: response.Name,
		},
	}, nil
}

func (s *Server) GetIssuedUserTicketsByAgentCode(ctx context.Context, in *pb.GetIssuedUserTicketsByAgentCodeRequest) (*pb.GetIssuedUserTicketsByAgentCodeResponse, error) {
	agentCode := in.GetCode()

	users, err := service.GetUserByAgentCode(agentCode)
	if err != nil {
		return nil, err
	}
	userIDs := []int{}
	for _, user := range users {
		userIDs = append(userIDs, int(user.ID))
	}

	if len(userIDs) == 0 {
		return nil, fmt.Errorf("no users")
	}

	userTickets, tickets, err := service.GetUserTicketsAndTicketsByUserIDs(userIDs)
	if err != nil {
		return nil, err
	}

	grpcUserTickets := []*pb.UserTicket{}
	for _, userTicket := range userTickets {
		ticketName := ""
		for _, ticket := range tickets {
			if userTicket.TicketID == ticket.ID {
				ticketName = ticket.Name
			}
		}
		grpcUserTicket := &pb.UserTicket{
			Code: userTicket.Code,
			Name: ticketName,
		}
		grpcUserTickets = append(grpcUserTickets, grpcUserTicket)
	}
	return &pb.GetIssuedUserTicketsByAgentCodeResponse{
		UserTickets: grpcUserTickets,
	}, nil
}

func Run() {
	port := flag.Int("port", helper.ConvertToInt(os.Getenv("GRPC_SERVER_PORT")), "The server port")
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
