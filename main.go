package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/VadimGossip/grpsProductsServer/gen/products"
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
	fetchResponse, err := client.Fetch(ctx, &products.FetchRequest{Url: url})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"handler": "fetch",
		}).Error(err)
	}
	fmt.Println(fetchResponse.Status)

	listResponse, err := client.List(ctx, &products.ListRequest{
		SortField:    products.SortingField_product_name,
		SortType:     products.SortingType_asc,
		PagingOffset: 30,
		PagingLimit:  20,
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"handler": "list",
		}).Error(err)
	}

	for idx := range listResponse.Product {
		fmt.Println(listResponse.Product[idx])
	}
}
