package main

import (
	"context"
	pb "github.com/abrampers/lagu-sion-backend/lagusion"
	empty "github.com/golang/protobuf/ptypes/empty"
)

type laguSionServer struct {
	pb.UnimplementedLaguSionServiceServer
}

func (s *laguSionServer) ListAllSongs(ctx context.Context, empty *empty.Empty) (*pb.ListSongResponse, error) {
	return &pb.ListSongResponse{}, nil
}

func (s *laguSionServer) ListSongs(ctx context.Context, ad *pb.ListSongsRequest) (*pb.ListSongResponse, error) {
	return &pb.ListSongResponse{}, nil
}

func main() {
}
