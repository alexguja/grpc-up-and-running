package main

import (
	"context"
	"log"
	"time"

	pb "productinfo/client/ecommerce" // Import generated proto code

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	addr = "localhost:50051"
)

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect %v", err)
	}

	// Close the connection when everything is done
	defer conn.Close()

	// Setup a connection with the gRPC server
	c := pb.NewProductInfoClient(conn)

	name := "Apple iPhone 15"
	description := "Meet Apple iPhone 15. All-new dual-camera system with ultra Wide and night mode"
	price := float32(1000.00)

	// Create a request context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call addProduct method with product details
	r, err := c.AddProduct(ctx, &pb.Product{Name: name, Description: description, Price: price})
	if err != nil {
		log.Fatalf("Could not add product %v", err)
	}

	log.Printf("Product ID: %s added succesfully", r.Value)

	// Call getProduct with the product ID.
	product, err := c.GetProduct(ctx, &pb.ProductID{Value: r.Value})
	if err != nil {
		log.Fatalf("Could not get product %v", err)
	}

	log.Printf("Product %v", product.String())
}
