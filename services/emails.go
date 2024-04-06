package services

import (
	"context"
	"fmt"
	"log"

	proto "github.com/clintvidler/identity-go/gen/proto/client"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type EmailClient struct {
	host string
	port string
}

func NewEmailClient(host, port string) *EmailClient {
	return &EmailClient{host: host, port: port}
}

func (c *EmailClient) Send(to, from, subject, body string) {
	conn, err := grpc.Dial(c.host+":"+c.port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println(err)
	}

	client := proto.NewEmailClient(conn)

	sendResponse, err := client.Send(context.Background(),
		&proto.EmailSendRequest{
			To:      to,
			From:    from,
			Subject: subject,
			Body:    body,
		})
	if err != nil {
		log.Println(err)
	}

	fmt.Println(sendResponse)
}
