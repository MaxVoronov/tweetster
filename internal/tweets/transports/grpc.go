package transports

import (
	"context"

	transport "github.com/go-kit/kit/transport/grpc"
	"github.com/golang/protobuf/ptypes"

	"github.com/maxvoronov/tweetster/internal/tweets/endpoints"
	"github.com/maxvoronov/tweetster/internal/tweets/pb"
)

type grpcServer struct {
	postsGetList transport.Handler
	postsGetById transport.Handler
}

func NewGRPCServer(tweetsEndpoints endpoints.Endpoints) pb.TweetsServiceServer {
	return &grpcServer{
		postsGetList: transport.NewServer(
			tweetsEndpoints.PostsGetListEndpoint,
			postsGetListDecode,
			postsGetListEncode,
		),
		postsGetById: transport.NewServer(
			tweetsEndpoints.PostsGetByIdEndpoint,
			postsGetByIdDecode,
			postsGetByIdEncode,
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

func (g grpcServer) PostsGetById(ctx context.Context, request *pb.PostsGetByIdRequest) (*pb.PostsGetByIdResponse, error) {
	_, response, err := g.postsGetById.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
	}
	return response.(*pb.PostsGetByIdResponse), nil
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
			Id:        post.Id,
			AuthorId:  post.AuthorId,
			Content:   post.Content,
			CreatedAt: createdAt,
		})
	}

	return &pb.PostsGetListResponse{Posts: posts}, nil
}

func postsGetByIdDecode(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.PostsGetByIdRequest)
	return &endpoints.PostsGetByIdRequest{Id: req.Id}, nil
}

func postsGetByIdEncode(_ context.Context, response interface{}) (interface{}, error) {
	res := response.(*endpoints.PostsGetByIdResponse)
	createdAt, err := ptypes.TimestampProto(res.Post.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &pb.PostsGetByIdResponse{
		Post: &pb.Post{
			Id:        res.Post.Id,
			AuthorId:  res.Post.AuthorId,
			Content:   res.Post.Content,
			CreatedAt: createdAt,
		},
	}, nil
}
