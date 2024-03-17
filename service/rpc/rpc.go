package rpc

import (
	"log"
	"os"

	"github.com/clintvidler/identity-go/gen/proto/go/proto"
	"github.com/clintvidler/identity-go/service/data"
)

type IdentityService struct {
	*proto.UnimplementedIdentityServiceServer
	data   *data.Store
	prvKey []byte
	pubKey []byte
}

func NewIdentityService(ds *data.Store) *IdentityService {
	prvKey, err := os.ReadFile("cert/id_rsa")

	if err != nil {
		log.Printf("Error: %s", err)
	}

	pubKey, err := os.ReadFile("cert/id_rsa.pub")

	if err != nil {
		log.Printf("Error: %s", err)
	}

	return &IdentityService{data: ds, prvKey: prvKey, pubKey: pubKey}
}
