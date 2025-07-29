package grpc

import (
	"context"
	"crypto/tls"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	*grpc.ClientConn
}

func (c *Client) Dial(ctx context.Context, cfg ConfigClient) error {
	var creds credentials.TransportCredentials

	if cfg.Cert != "" && cfg.Key != "" {
		cert, err := tls.LoadX509KeyPair(cfg.Cert, cfg.Key)
		if err != nil {
			return err
		}

		creds = credentials.NewServerTLSFromCert(&cert)
	} else {
		creds = insecure.NewCredentials()
	}

	secureOption := grpc.WithTransportCredentials(creds)

	host := strings.Join([]string{cfg.Hostname, cfg.Port}, ":")

	var err error

	if c.ClientConn, err = grpc.NewClient(
		host,
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(int(cfg.MaxSendMsgSize))),
		secureOption,
	); err != nil {
		return err
	}

	return nil
}

func (c *Client) Close(ctx context.Context) error {
	return c.ClientConn.Close()
}
