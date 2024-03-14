package server

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/clintvidler/identity-go/service"

	"github.com/clintvidler/identity-go/gen/proto/go/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	openapiMW "github.com/go-openapi/runtime/middleware"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Serve() {

	// Start the gRPC server

	grpcServer := grpc.NewServer()

	proto.RegisterIdentityServiceServer(grpcServer, &service.IdentityService{})

	log.Println("gRPC server ready on :9090")

	lis, err := net.Listen("tcp", ":9090")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go grpcServer.Serve(lis)

	// Dial the gRPC server to make a client connection

	conn, err := grpc.Dial(":9090", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	// Start the RESTful server

	s.serveMux(ctx, conn)
}

func (s *Server) serveMux(ctx context.Context, conn *grpc.ClientConn) {

	// Create and register the HTTP router with the gRPC service and connection

	rmux := runtime.NewServeMux()

	client := proto.NewIdentityServiceClient(conn)

	err := proto.RegisterIdentityServiceHandlerClient(ctx, rmux, client)

	if err != nil {
		log.Fatal(err)
	}

	// Create a standard HTTP router

	mux := http.NewServeMux()

	// Mount the gRPC HTTP gateway route

	mux.Handle("/", rmux)

	// Mount the generated OpenAPI specification route

	mux.HandleFunc("/swagger-ui/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./gen/proto/service.swagger.json")
	})

	// Mount the Swagger UI template

	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", http.FileServer(http.Dir("./swagger-ui"))))

	// Mount the Redoc template

	mux.Handle("/json", openapiMW.Redoc(openapiMW.RedocOpts{SpecURL: "/swagger-ui/swagger.json", Path: "json"}, nil))

	log.Println("HTTP server ready on :8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
