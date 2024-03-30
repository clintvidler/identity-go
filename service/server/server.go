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

func streamInterceptor(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	return handler(srv, ss)
}

func (s *Server) Serve() {
	// Create the gRPC server
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(rpc.IsAuthInterceptor),
		grpc.StreamInterceptor(streamInterceptor),
	)
	proto.RegisterIdentityServiceServer(grpcServer, rpc.NewIdentityService(s.ds))

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
	mux.Handle("/", addCORSHeaders(rmux))

	// Mount the docs
	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./gen/proto/service.swagger.json")
	})
	mux.Handle("/docs", openapiMW.Redoc(openapiMW.RedocOpts{SpecURL: "/swagger.json", Path: "docs"}, nil))

	// Start serving RESTful on port 8080
	log.Println("HTTP server ready on :8080")
	if err := http.ListenAndServe(":8080", addCORSHeaders(mux)); err != nil {
		log.Fatal(err)
	}
}

// addCORSHeaders is a middleware function to add CORS headers to responses.
func addCORSHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set specific origin instead of '*'
		origin := r.Header.Get("Origin")
		w.Header().Set("Access-Control-Allow-Origin", origin)

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, X-User-Agent, X-Grpc-Web")
		// Note: "*" doesn't work for withCredentials requests
		w.Header().Set("Access-Control-Expose-Headers", "Grpc-Metadata-Access-Token, Grpc-Metadata-Refresh-Token")

		// Allow credentials
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Stop here if its Preflighted OPTIONS request
		if r.Method == "OPTIONS" {
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
