package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	pb "grpc_demo/proto"

	"google.golang.org/grpc"
)

type ServerImplementation struct {
	pb.UnimplementedCoffeeShopServer
}

func (s *ServerImplementation) GetMenu(menuRequest *pb.MenuRequest, svr pb.CoffeeShop_GetMenuServer) error {
	items := []*pb.Item{
		{Name: "Espresso", Id: "1"},
		{Name: "Cappuccino", Id: "2"},
		{Name: "Latte", Id: "3"},
	}

	for i, _ := range items {
		// svr.Send(&pb.Menu{Items: items[i:i+1]})
		svr.Send(&pb.Menu{Items: items[0 : i+1]})
		time.Sleep(2 * time.Second) // simulate some delay
	}

	return nil
}

func (s *ServerImplementation) PlaceOrder(context.Context, *pb.Order) (*pb.Receipt, error) {
	return &pb.Receipt{
		Id: "some dummy id",
	}, nil
}

func (s *ServerImplementation) GetOrderStatus(context.Context, *pb.Receipt) (*pb.OrderStatus, error) {
	return &pb.OrderStatus{
		OrderId: "ABC123",
		Status:  "IN PROGRESS",
	}, nil
}

func main() {
	// setup listener on PORT 9000
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Fail to listen on port 9000 %v", err)
	}

	grpcServer := grpc.NewServer()
	// Now bind our server to grpcServer
	pb.RegisterCoffeeShopServer(grpcServer, &ServerImplementation{})

	go func() {
		log.Println("gRPC server is running on port 9000")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()

	// Wait for Ctrl+C or SIGTERM
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down gRPC server...")
	grpcServer.GracefulStop()
	log.Println("gRPC server stopped gracefully")

}
