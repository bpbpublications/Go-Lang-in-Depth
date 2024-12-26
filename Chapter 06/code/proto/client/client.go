package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"math/rand"
	cat "proto/catalog"
	"strconv"
	"time"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("addr", "localhost:50051", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
)

func printProduct(client cat.CatalogClient, point *cat.ProductSpec) {
	log.Printf("Getting feature for product (%d, %d)", point.Name)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	product, err := client.GetProduct(ctx, point)
	if err != nil {
		log.Fatalf("client.GetProduct failed: %v", err)
	}
	log.Println("Client: Received from server - " + "product name " + product.Name + " ,Image " + product.Image)
}

func printProducts(client cat.CatalogClient, rect *cat.Category) {
	log.Printf("Client: Looking for products on the server within %v", rect)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.ProductList(ctx, rect)
	if err != nil {
		log.Fatalf("client.ProductList failed: %v", err)
	}
	for {
		product, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("client.ProductList failed: %v", err)
		}
		log.Printf("Product: name: %q, product:(%v)", product.GetName(),
			product.GetImage())
	}
}

func RecordProduct(client cat.CatalogClient) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	pointCount := int(r.Int31n(100)) + 2
	var points []*cat.Product
	for i := 0; i < pointCount; i++ {
		points = append(points, randomProduct(r))
	}
	log.Printf("client: catalog has %d products created.", len(points))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.RecordProduct(ctx)

	if err != nil {
		log.Fatalf("client.RecordProduct failed: %v", err)
	}
	for _, point := range points {
		log.Printf("client sending product" + point.Name)
		if err := stream.Send(&cat.ProductSpec{Name: point.Name}); err != nil {
			log.Fatalf("client.RecordProduct: stream.Send(%v) failed: %v", point, err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("client.RecordPoduct failed: %v", err)
	}
	log.Printf("Produuct summary: %v", reply)
}

func ProductChat(client cat.CatalogClient) {
	notes := []*cat.ProductSpecNote{
		{Cat: &cat.Category{Name: "Athletics"},
			Spec: &cat.ProductSpec{Name: "pump"}},
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.ProductChat(ctx)
	if err != nil {
		log.Fatalf("client.RouteChat failed: %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("client.ProductChat failed: %v", err)
			}
			log.Printf("Client Got message %s at product(%d)", in.Cat.Name, in.Spec.Name)
		}
	}()
	for _, note := range notes {
		if err := stream.Send(note); err != nil {
			log.Fatalf("client.ProductChat: stream.Send(%v) failed: %v", note, err)
		}
	}
	stream.CloseSend()
	<-waitc
}

func randomProduct(r *rand.Rand) *cat.Product {
	name := "Product" + strconv.Itoa(r.Intn(180))
	img := name + ".jpg"
	categ := "General"
	return &cat.Product{Name: name, Image: img, Catname: categ}
}

func main() {
	flag.Parse()
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := cat.NewCatalogClient(conn)

	printProduct(client, &cat.ProductSpec{Name: "cycle"})

	printProduct(client, &cat.ProductSpec{Name: "pump"})

	printProducts(client, &cat.Category{Name: "Electronics"})

	RecordProduct(client)

	ProductChat(client)
}
