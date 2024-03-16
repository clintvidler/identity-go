package server

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/clintvidler/identity-go/gen/proto/go/proto"
	"github.com/clintvidler/identity-go/service/data"
	"github.com/clintvidler/identity-go/service/rpc"

	openapiMW "github.com/go-openapi/runtime/middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	ds *data.Store
}

func NewServer(ds *data.Store) *Server {
	return &Server{ds: ds}
}

func (s *Server) Serve() {
	// Create the gRPC server
	grpcServer := grpc.NewServer()
	proto.RegisterIdentityServiceServer(grpcServer, &rpc.IdentityService{})

	// Start serving gRPC on port 9090
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	go grpcServer.Serve(lis)
	log.Println("gRPC server ready on :9090")

	// Dial the gRPC server to make a client connection
	conn, err := grpc.Dial(":9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	// Create the RESTful server using the gRPC client and connection
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	rmux := runtime.NewServeMux()
	client := proto.NewIdentityServiceClient(conn)
	err = proto.RegisterIdentityServiceHandlerClient(ctx, rmux, client)
	if err != nil {
		log.Fatal(err)
	}

	// Create a standard HTTP router and mount the gRPC gateway
	mux := http.NewServeMux()
	mux.Handle("/", rmux)

	// Mount the docs
	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./gen/proto/service.swagger.json")
	})
	mux.Handle("/docs", openapiMW.Redoc(openapiMW.RedocOpts{SpecURL: "/swagger.json", Path: "docs"}, nil))

	// Start serving RESTful on port 8080
	log.Println("HTTP server ready on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
