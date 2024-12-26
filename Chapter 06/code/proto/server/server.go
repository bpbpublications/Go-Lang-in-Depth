package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"

	"google.golang.org/grpc"

	cat "proto/catalog"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 50051, "The server port")
)

type catalogServer struct {
	cat.UnimplementedCatalogServer
	savedProducts []*cat.Product // read-only after initialized

	mu               sync.Mutex // protects routeNotes
	productSpecNotes map[string][]*cat.ProductSpecNote
}

func (s *catalogServer) GetProduct(ctx context.Context, point *cat.ProductSpec) (*cat.Product, error) {
	for _, feature := range s.savedProducts {
		if feature.Name == point.Name {
			return feature, nil
		}
	}
	return &cat.Product{Name: point.Name, Image: point.Name + ".jpg"}, nil
}

func (s *catalogServer) ProductList(rect *cat.Category, stream cat.Catalog_ProductListServer) error {
	for _, feature := range s.savedProducts {
		if feature.Catname == rect.Name {
			if err := stream.Send(feature); err != nil {
				return err
			}
		}
	}
	return nil
}

func (s *catalogServer) RecordProduct(stream cat.Catalog_RecordProductServer) error {
	var pointCount, featureCount int32

	for {
		point, err := stream.Recv()
		//log.Printf("error " + err.Error())
		if err == io.EOF {
			//log.Printf(point.Name)
			return stream.SendAndClose(&cat.Product{
				Name:    "Completed",
				Image:   ".jpg",
				Catname: "Messages",
			})
		}
		if err != nil {
			return err
		}
		pointCount++
		for _, feature := range s.savedProducts {
			log.Printf("server: recorded " + point.Name)
			if feature.Name == point.Name {
				featureCount++

			}
		}
	}
}

func (s *catalogServer) ProductChat(stream cat.Catalog_ProductChatServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		key := serializeProductSpec(in)

		s.mu.Lock()
		s.productSpecNotes[key] = append(s.productSpecNotes[key], in)
		rn := make([]*cat.ProductSpecNote, len(s.productSpecNotes[key]))
		copy(rn, s.productSpecNotes[key])
		s.mu.Unlock()

		for _, note := range rn {
			if err := stream.Send(note); err != nil {
				return err
			}
		}
	}
}

func (s *catalogServer) loadProductSpecs(filePath string) {
	var data []byte
	if filePath != "" {
		var err error
		data, err = os.ReadFile(filePath)
		if err != nil {
			log.Fatalf("Failed to load default features: %v", err)
		}
	} else {
		data = exampleData
	}
	s.savedProducts = make([]*cat.Product, 0)
	if err := json.Unmarshal(data, &s.savedProducts); err != nil {
		log.Fatalf("Failed to load default features: %v", err)
	}
}

func serializeProductSpec(point *cat.ProductSpecNote) string {
	return fmt.Sprintf("%d", point.Spec.Name)
}
func serialize(point *cat.Product) string {
	return fmt.Sprintf("%d %d", point.Name, point.Image)
}

func newServer() *catalogServer {
	s := &catalogServer{productSpecNotes: make(map[string][]*cat.ProductSpecNote)}
	s.loadProductSpecs(*jsonDBFile)
	return s
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	cat.RegisterCatalogServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}

var exampleData = []byte(`[{"name": "cycle", "image" : "cycle.jpg", "catname" : "Athletics"}, {"name": "pump", "image" : "pump.jpg","catname" : "Athletics"},  {"name": "Product3", "image" : "product3.jpg","catname" : "Electronics"}]`)
