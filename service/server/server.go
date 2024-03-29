package server

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/clintvidler/identity-go/gen/proto/go/proto"
	"github.com/clintvidler/identity-go/service/data"
	"github.com/clintvidler/identity-go/service/rpc"

	openapiMW "github.com/go-openapi/runtime/middleware"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
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

	// Wrap the gRPC server with gRPC-Web handler
	wrappedGrpc := grpcweb.WrapServer(grpcServer)

	// Start serving on port 9090
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()
	log.Println("gRPC server ready on :9090")

	// Create a standard HTTP router
	mux := http.NewServeMux()

	// Mount the gRPC-Web handler
	mux.HandleFunc("/proto.IdentityService/", func(w http.ResponseWriter, req *http.Request) {
		if wrappedGrpc.IsGrpcWebRequest(req) {
			wrappedGrpc.ServeHTTP(w, req)
			return
		}
		http.Error(w, "Unexpected request", http.StatusBadRequest)
	})

	// Mount the RESTful handlers
	rmux := runtime.NewServeMux()
	err = proto.RegisterIdentityServiceHandlerServer(context.Background(), rmux, rpc.NewIdentityService(s.ds))
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/", rmux)

	// Mount the docs
	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./gen/proto/service.swagger.json")
	})
	mux.Handle("/docs", openapiMW.Redoc(openapiMW.RedocOpts{SpecURL: "/swagger.json", Path: "docs"}, nil))

	// Start serving RESTful and gRPC-Web on port 8080
	log.Println("HTTP server ready on :8080")
	if err := http.ListenAndServe(":8080", addCORSHeaders(mux)); err != nil {
		log.Fatal(err)
	}
}

// addCORSHeaders is a middleware function to add CORS headers to responses.
func addCORSHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, X-User-Agent, X-Grpc-Web")

		// Stop here if its Preflighted OPTIONS request
		if r.Method == "OPTIONS" {
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
