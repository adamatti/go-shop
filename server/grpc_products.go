package main

import (
	"context"

	pb "adamatti.github.io/go-shop/grpc"
	"adamatti.github.io/go-shop/repo"
	"go.mongodb.org/mongo-driver/v2/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type shopServer struct {
	pb.UnimplementedShopServer
}

func (s *shopServer) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsResponse, error) {
	// TODO: it shall call a service layer, not repo
	products, err := repo.ListProducts()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list products: %v", err)
	}

	respProducts := make([]*pb.Product, 0, len(products))
	for _, p := range products {
		var id string
		if !p.ID.IsZero() {
			id = p.ID.Hex()
		}
		respProducts = append(respProducts, &pb.Product{
			Id:                id,
			Name:              p.Name,
			Description:       p.Description,
			Price:             p.Price,
			AvailableQuantity: int32(p.AvailableQuantity),
		})
	}

	return &pb.ListProductsResponse{
		Products: respProducts,
	}, nil
}

func (s *shopServer) InsertProduct(ctx context.Context, req *pb.InsertProductRequest) (*pb.InsertProductResponse, error) {
	p := req.GetProduct()
	if p == nil {
		return nil, status.Error(codes.InvalidArgument, "product is required")
	}

	var productID bson.ObjectID
	var err error
	if p.GetId() != "" {
		productID, err = bson.ObjectIDFromHex(p.GetId())
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid product id: %v", err)
		}
	} else {
		productID = bson.NewObjectID()
	}

	product := repo.Product{
		ID:                productID,
		Name:              p.GetName(),
		Description:       p.GetDescription(),
		Price:             p.GetPrice(),
		AvailableQuantity: int(p.GetAvailableQuantity()),
	}

	// TODO: it shall call a service layer, not repo
	if err := repo.InsertProduct(product); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to insert product: %v", err)
	}

	return &pb.InsertProductResponse{
		Id: productID.Hex(),
	}, nil
}
