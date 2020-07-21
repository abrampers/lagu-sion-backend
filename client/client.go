package main

import (
	"context"
	"flag"
	"log"

	pb "github.com/abrampers/lagu-sion-backend/lagusion"
	empty "github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/testdata"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	if *tls {
		if *caFile == "" {
			*caFile = testdata.Path("ca.pem")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	opts = append(opts, grpc.WithBlock())
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	ctx := context.Background()
	client := pb.NewLaguSionServiceClient(conn)
	resp, err := client.ListAllSongs(ctx, &empty.Empty{})
	if err != nil {
		log.Fatalf("SayHello failed: %v", err)
	}
	log.Printf("Response: %+v", resp)

	resp, err = client.ListSongs(ctx, &pb.ListSongsRequest{SongBook: pb.SongBook_LAGU_SION_EDISI_LENGKAP})
	if err != nil {
		log.Fatalf("SayHello failed: %v", err)
	}
	log.Printf("Response: %+v", resp)
}
