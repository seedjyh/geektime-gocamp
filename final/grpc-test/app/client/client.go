package client

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-test/api"
	"time"
)

func Run(serverAddr string) error {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		return errors.Wrap(err, "fail to dial")
	}
	defer conn.Close()
	client := pb.NewAuthClient(conn)

	// add some userToken
	if err := setUserToken(client, "id_123", "token_abc123abc"); err != nil {
		return err
	}
	if err := setUserToken(client, "id_456", "token_def456def"); err != nil {
		return err
	}

	// query some userToken
	if err := printUserToken(client, "id_456"); err != nil {
		return err
	}
	if err := printUserToken(client, "id_123"); err != nil {
		return err
	}
	return nil
}

func printUserToken(client pb.AuthClient, userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if idAndToken, err := client.GetUserToken(ctx, &pb.UserId{Id: userId}); err != nil {
		return errors.Wrap(err, "fail to getUserToken")
	} else {
		fmt.Println("Token for userId", userId, "is:", idAndToken.Token)
		return nil
	}
}

func setUserToken(client pb.AuthClient, userId string, userToken string) error {
	idAndToken := &pb.UserIdAndToken{
		UserId: &pb.UserId{Id: userId},
		Token:  userToken,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if _, err := client.SetUserToken(ctx, idAndToken); err != nil {
		return errors.Wrap(err, "fail to setUserToken")
	}
	return nil
}
