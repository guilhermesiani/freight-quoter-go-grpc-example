package freight_quoter

import (
	"golang.org/x/net/context"
	"log"
)

type Server struct {

}

func (s *Server) MountPackage(
	ctx context.Context,
	toPackage *ToPackage,
) (*Package, error) {
	log.Printf("Received items to package: %s", toPackage)
	return &Package{Price: 91, Quantity: 3, Volume: 21}, nil
}

func (s *Server) Quote(
	ctx context.Context,
	quotation *Quotation,
) (*Delivery, error) {
	log.Printf("Received quotation to calculate: %s", quotation)
	return &Delivery{Time: 3, Price: 32}, nil
}