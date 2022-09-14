package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/VadimGossip/grpsProductsServer/gen/products"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	host = "localhost"
	port = 9000
	url  = "http://164.92.251.245:8080/api/v1/products/"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatal(err)
	}
	defer conn.Close()
	ctx := context.Background()

	client := products.NewProductsServiceClient(conn)
	response, err := client.Fetch(ctx, &products.FetchRequest{Url: url})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"handler": "fetch",
		}).Error(err)
	}

	fmt.Println(response.Status)

}
