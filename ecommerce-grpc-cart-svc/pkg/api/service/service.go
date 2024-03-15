package service

import (
	pb "cart/service/pkg/pb/cart"
	interfaces "cart/service/pkg/usecase/interface"
	"context"
)

type CartServer struct {
	CartUseCase interfaces.CartUseCase
	pb.UnimplementedCartServer
}

func NewCartServer(useCase interfaces.CartUseCase) pb.CartServer {
	return &CartServer{
		CartUseCase: useCase,
	}
}

func (c *CartServer) AddToCart(ctx context.Context, req *pb.AddToCartRequest) (*pb.AddToCartResponse, error) {
	productID := int(req.ProductID)
	userID := int(req.UserID)
	quantity := int(req.Quantity)
	res, err := c.CartUseCase.AddToCart(productID, userID, quantity)
	if err != nil {
		return &pb.AddToCartResponse{
			Error: err.Error(),
		}, err
	}
	var result pb.AddToCartResponse
	var cartDetails []*pb.CartDetails
	for _, v := range res.Cart {
		details := &pb.CartDetails{
			ProductID:  int64(v.ProductID),
			Quantity:   float32(v.Quantity),
			TotalPrice: float32(v.TotalPrice),
		}
		cartDetails = append(cartDetails, details)
	}
	result.Price = float32(res.TotalPrice)
	result.Cart = cartDetails
	return &result, nil
}

func (c *CartServer) GetCart(ctx context.Context, req *pb.GetCartRequest) (*pb.GetCartResponse, error) {
	userID := int(req.UserID)
	res, err := c.CartUseCase.DisplayCart(userID)
	if err != nil {
		return &pb.GetCartResponse{
			Error: err.Error(),
		}, err
	}
	var result pb.GetCartResponse
	var cartDetails []*pb.CartDetails
	for _, v := range res.Cart {
		details := &pb.CartDetails{
			ProductID:  int64(v.ProductID),
			Quantity:   float32(v.Quantity),
			TotalPrice: float32(v.TotalPrice),
		}
		cartDetails = append(cartDetails, details)
	}
	result.Price = float32(res.TotalPrice)
	result.Cart = cartDetails
	return &result, nil
}
func (c *CartServer) DoesCartExist(ctx context.Context, req *pb.DoesCartExistRequest) (*pb.DoesCartExistReponse, error) {
	userID := int(req.UserID)
	exists, err := c.CartUseCase.DoesCartExist(userID)
	if err != nil {
		return &pb.DoesCartExistReponse{
			Error: err.Error(),
		}, err
	}
	return &pb.DoesCartExistReponse{
		Data: exists,
	}, nil
}
func (c *CartServer) GetAllItemsFromCart(ctx context.Context, req *pb.GetAllItemsFromCartRequest) (*pb.GetAllItemsFromCartResponse, error) {
	userID := int(req.UserID)
	res, err := c.CartUseCase.GetAllItemsFromCart(userID)
	if err != nil {
		return &pb.GetAllItemsFromCartResponse{
			Error: err.Error(),
		}, err
	}

	var cartDetails []*pb.CartDetails
	for _, cartItem := range res {
		cartDetail := &pb.CartDetails{
			ProductID:  int64(cartItem.ProductID),
			Quantity:   float32(cartItem.Quantity),
			TotalPrice: float32(cartItem.TotalPrice),
		}
		cartDetails = append(cartDetails, cartDetail)
	}

	return &pb.GetAllItemsFromCartResponse{
		Cart: cartDetails,
	}, nil
}

func (c *CartServer) TotalAmountInCart(ctx context.Context, req *pb.TotalAmountInCartRequest) (*pb.TotalAmountInCartResponse, error) {
	userID := int(req.UserID)
	res, err := c.CartUseCase.TotalAmountInCart(userID)
	if err != nil {
		return &pb.TotalAmountInCartResponse{
			Error: err.Error(),
		}, err
	}
	return &pb.TotalAmountInCartResponse{
		Data: float32(res),
	}, nil
}

func (c *CartServer) UpdateCartAfterOrder(ctx context.Context, req *pb.UpdateCartAfterOrderRequest) (*pb.UpdateCartAfterOrderResponse, error) {
	userID := int(req.UserID)
	productID := int(req.ProductID)
	quantity := float64(req.Quantity)
	err := c.CartUseCase.UpdateCartAfterOrder(userID, productID, float64(quantity))
	if err != nil {
		return &pb.UpdateCartAfterOrderResponse{
			Error: err.Error(),
		}, err
	}

	return &pb.UpdateCartAfterOrderResponse{}, nil
}
