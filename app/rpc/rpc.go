package rpc

import (
	"log"
	"os"

	"github.com/clintvidler/identity-go/app/data"
	proto "github.com/clintvidler/identity-go/gen/proto/server"
	"github.com/clintvidler/identity-go/services"
)

type IdentityService struct {
	*proto.UnimplementedIdentityServiceServer
	data        *data.Store
	emailClient *services.EmailClient
	prvKey      []byte
	pubKey      []byte
}

func NewIdentityService(ds *data.Store, ec *services.EmailClient) *IdentityService {
	prvKey, err := os.ReadFile("cert/id_rsa")

	if err != nil {
		log.Printf("Error: %s", err)
	}

	pubKey, err := os.ReadFile("cert/id_rsa.pub")

	if err != nil {
		log.Printf("Error: %s", err)
	}

	return &IdentityService{data: ds, emailClient: ec, prvKey: prvKey, pubKey: pubKey}
}
