package grpc

import (
	"context"

	pb "github.com/hatena/Hatena-Intern-2021/services/fetcher/pb/fetcher"
)

func (s *Server) Fetch(ctx context.Context, in *pb.FetchRequest) (*pb.FetchReply, error) {
	title := "title"
	return &pb.FetchReply{Title: title}, nil
}
