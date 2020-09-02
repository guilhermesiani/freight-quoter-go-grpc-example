package main

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	freightQuoter "github.com/guilhermesiani/go-grpc/freight_quoter"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	defer conn.Close()

	c := freightQuoter.NewFreightQuoterClient(conn)

	var items []*freightQuoter.Item
	items = append(items, &freightQuoter.Item{
		Quantity: 1,
		Price: 2,
		Width: 3,
		Height: 4,
		Length: 5,
	})

	toPackage := freightQuoter.ToPackage{
		Items: items,
	}

	packg, err := c.MountPackage(context.Background(), &toPackage)
	if err != nil {
		log.Fatalf("Error when calling MountPackage: %s", err)
	}

	log.Printf("Package Response from Server: %s", packg)

	delivery, err := c.Quote(context.Background(), &freightQuoter.Quotation{
		Package: packg,
		ZipCode: "13607222",
	})
	if err != nil {
		log.Fatalf("Error when calling Quote: %s", err)
	}

	log.Printf("Delivery Response from Server: %s", delivery)
}
