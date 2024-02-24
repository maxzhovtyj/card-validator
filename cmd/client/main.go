package main

import (
	"context"
	"fmt"
	"github.com/maxzhovtyj/card-validator/pkg/log/applogger"
	pb "github.com/maxzhovtyj/card-validator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"sync"
)

func main() {
	logger := applogger.New()

	conn, err := grpc.Dial(":7799", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal(err)
	}

	defer conn.Close()

	client := pb.NewCardServiceClient(conn)

	testTable := []struct {
		number          string
		expirationMonth string
		expirationYear  int64
	}{
		{"4111111111111111", "12", 2028},
		{"4111111111111111", "01", 2021},
		{"4111111111111111", "40", 2021},
		{"1111111111111", "10", 2028},
		{"411111ABC1111111", "10", 2028},
	}

	var wg sync.WaitGroup

	for _, testCase := range testTable {
		wg.Add(1)
		req := &pb.ValidateCardRequest{Card: &pb.Card{Number: testCase.number, ExpirationMonth: testCase.expirationMonth, ExpirationYear: testCase.expirationYear}}

		go func(wg *sync.WaitGroup, req *pb.ValidateCardRequest) {
			defer wg.Done()
			resp, err := client.Validate(context.Background(), req)
			if err != nil {
				errStatus, _ := status.FromError(err)
				fmt.Println(errStatus.Message())
			}

			logger.Infof("%s %s/%d - card validation result: '%s'", req.Card.Number, req.Card.ExpirationMonth, req.Card.ExpirationYear, resp.String())
		}(&wg, req)
	}

	wg.Wait()
}
