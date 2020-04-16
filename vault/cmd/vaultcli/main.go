package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/namtx/go-blueprints/vault"
	grpcclient "github.com/namtx/go-blueprints/vault/client/grpc"
	"google.golang.org/grpc"
)

func main() {
	var gRPCAddr = flag.String("addr", ":8081", "gRPC Address")

	flag.Parse()

	ctx := context.Background()
	conn, err := grpc.Dial(*gRPCAddr, grpc.WithInsecure(), grpc.WithTimeout(1*time.Second))
	if err != nil {
		log.Fatalln("gRPC dial:", err)
	}

	defer conn.Close()
	vaultService := grpcclient.New(conn)
	args := flag.Args()
	var cmd string
	cmd, args = pop(args)
	switch cmd {
	case "hash":
		var password string
		password, args = pop(args)
		hash(ctx, vaultService, password)
	case "validate":
		var password, hash string
		password, args = pop(args)
		hash, args = pop(args)
		validate(ctx, vaultService, password, hash)
	default:
		log.Fatalln("Unknown command", cmd)
	}
}

func hash(ctx context.Context, service vault.Service, password string) {
	h, err := service.Hash(ctx, password)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(h)
}

func validate(ctx context.Context, service vault.Service, password, hash string) {
	valid, err := service.Validate(ctx, password, hash)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(valid)
}

func pop(s []string) (string, []string) {
	if len(s) == 0 {
		return "", s
	}

	return s[0], s[1:]
}
