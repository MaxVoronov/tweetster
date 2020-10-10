package transports

import (
	"context"

	transport "github.com/go-kit/kit/transport/grpc"

	"github.com/maxvoronov/tweetster/internal/pb"
	"github.com/maxvoronov/tweetster/internal/users/endpoints"
)

type grpcServer struct {
	usersGetByID transport.Handler
}

func NewGRPCServer(usersEndpoints endpoints.Endpoints) pb.UsersServiceServer {
	return &grpcServer{
		usersGetByID: transport.NewServer(
			usersEndpoints.UsersGetByIDEndpoint,
			usersGetByIDDecode,
			usersGetByIDEncode,
		),
	}
}

func (g grpcServer) UserGetByID(ctx context.Context, request *pb.UserGetByIDRequest) (*pb.UserGetByIDResponse, error) {
	_, response, err := g.usersGetByID.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(*pb.UserGetByIDResponse), nil
}

func usersGetByIDDecode(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.UserGetByIDRequest)
	return &endpoints.UsersGetByIDRequest{ID: req.Id}, nil
}

func usersGetByIDEncode(_ context.Context, response interface{}) (interface{}, error) {
	user := response.(*endpoints.UsersGetByIDResponse).User
	return &pb.UserGetByIDResponse{
		User: &pb.User{
			Id:    user.ID,
			Login: user.Login,
			Email: user.Email,
			Name:  user.Name,
		},
	}, nil
}
