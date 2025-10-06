package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "grpc_demo/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to grpc server %v", err)
	}
	defer conn.Close()

	my_client := pb.NewCoffeeShopClient(conn)

	// creating custom context with 7 second timeout (6 second for order streaming and extra second for order placement)
	// anything less than 7 second would fail here
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	menuStream, err := my_client.GetMenu(ctx, &pb.MenuRequest{}) // sending a blank request
	if err != nil {
		log.Fatalf("error while calling GetMenu %v", err)
	}

	done := make(chan bool)

	var items []*pb.Item

	// running a go routine bcoz streams are blocking calls
	go func() {
		for {
			resp, err := menuStream.Recv()
			if err != nil && err != io.EOF {
				log.Printf("error while receiving menu %v", err)
				done <- true
				return
			}

			if err == io.EOF {
				done <- true
				return
			}

			items = resp.Items
			log.Printf("received menu item: %v", resp.Items)
		}
	}()

	<-done

	// receipt, err := my_client.PlaceOrder(ctx, &pb.Order{Items: items})
	receipt, err := my_client.PlaceOrder(ctx, &pb.Order{Items: items[:2]}) // ordering only first 2 items
	if err != nil {
		log.Fatalf("error while calling PlaceOrder %v", err)
	}

	status, err := my_client.GetOrderStatus(ctx, receipt)
	if err != nil {
		log.Fatalf("error while calling GetOrderStatus %v", err)
	}

	log.Printf("order status: %v", status)
}
