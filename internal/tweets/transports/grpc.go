package transports

import (
	"context"

	transport "github.com/go-kit/kit/transport/grpc"
	"github.com/golang/protobuf/ptypes"

	"github.com/maxvoronov/tweetster/internal/pb"
	"github.com/maxvoronov/tweetster/internal/tweets/endpoints"
)

type grpcServer struct {
	postsGetList transport.Handler
	postsGetByID transport.Handler
}

func NewGRPCServer(tweetsEndpoints endpoints.Endpoints) pb.TweetsServiceServer {
	return &grpcServer{
		postsGetList: transport.NewServer(
			tweetsEndpoints.PostsGetListEndpoint,
			postsGetListDecode,
			postsGetListEncode,
		),
		postsGetByID: transport.NewServer(
			tweetsEndpoints.PostsGetByIDEndpoint,
			postsGetByIDDecode,
			postsGetByIDEncode,
		),
	}
}

func (g grpcServer) PostsGetList(ctx context.Context, request *pb.PostsGetListRequest) (*pb.PostsGetListResponse, error) {
	_, response, err := g.postsGetList.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(*pb.PostsGetListResponse), nil
}

func (g grpcServer) PostsGetByID(ctx context.Context, request *pb.PostsGetByIDRequest) (*pb.PostsGetByIDResponse, error) {
	_, response, err := g.postsGetByID.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(*pb.PostsGetByIDResponse), nil
}

func postsGetListDecode(_ context.Context, _ interface{}) (interface{}, error) {
	return &endpoints.PostsGetListRequest{}, nil
}

func postsGetListEncode(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(*endpoints.PostsGetListResponse)
	posts := make([]*pb.Post, 0, len(res.Posts))
	for _, post := range res.Posts {
		createdAt, err := ptypes.TimestampProto(post.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &pb.Post{
			Id:        post.ID,
			AuthorId:  post.AuthorID,
			Content:   post.Content,
			CreatedAt: createdAt,
		})
	}

	return &pb.PostsGetListResponse{Posts: posts}, nil
}

func postsGetByIDDecode(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.PostsGetByIDRequest)
	return &endpoints.PostsGetByIDRequest{ID: req.Id}, nil
}

func postsGetByIDEncode(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(*endpoints.PostsGetByIDResponse)
	createdAt, err := ptypes.TimestampProto(res.Post.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &pb.PostsGetByIDResponse{
		Post: &pb.Post{
			Id:        res.Post.ID,
			AuthorId:  res.Post.AuthorID,
			Content:   res.Post.Content,
			CreatedAt: createdAt,
		},
	}, nil
}
