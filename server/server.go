package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/abrampers/lagu-sion-backend/lagusion"
	"github.com/abrampers/lagu-sion-backend/models"
	"google.golang.org/grpc"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	// tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	// certFile   = flag.String("cert_file", "", "The TLS cert file")
	// keyFile    = flag.String("key_file", "", "The TLS key file")
	port = flag.Int("port", 10000, "The server port")
)

var (
	dbHost     = os.Getenv("DB_HOST")
	dbPort     = os.Getenv("DB_PORT")
	dbName     = os.Getenv("DB_NAME")
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
)

func convert(songs []models.Song) []*pb.Song {
	var convertedSongs []*pb.Song
	for _, song := range songs {
		convertedSongs = append(convertedSongs, song.LaguSionSong())
	}
	return convertedSongs
}

type laguSionServer struct {
	pb.UnimplementedLaguSionServiceServer
	db *gorm.DB
}

func (s *laguSionServer) ListSongs(ctx context.Context, request *pb.ListSongRequest) (*pb.ListSongResponse, error) {
	var songs []models.Song
	var db *gorm.DB

	if request.SortOptions == pb.SortOptions_NUMBER {
		db = s.db.Order("book_id, number")
	} else if request.SortOptions == pb.SortOptions_ALPHABET {
		db = s.db.Order("book_id, title")
	}

	if request.SongBook == pb.SongBook_ALL {
		db.Find(&songs)
	} else {
		db.Where(&models.Song{BookId: uint32(request.SongBook)}).Find(&songs)
	}
	return &pb.ListSongResponse{Songs: convert(songs)}, nil
}

func main() {
	flag.Parse()
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPassword))
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer db.Close()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterLaguSionServiceServer(grpcServer, &laguSionServer{db: db.Set("gorm:auto_preload", true)})
	grpcServer.Serve(lis)
}
