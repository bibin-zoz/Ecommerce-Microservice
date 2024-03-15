package usecase

import (
	clinetinterface "cart/service/pkg/client/interface"
	interfaces "cart/service/pkg/repository/interface"
	services "cart/service/pkg/usecase/interface"
	"cart/service/pkg/utils/models"
	"errors"
)

type cartUseCase struct {
	cartRepository    interfaces.CartRepository
	productRepository clinetinterface.NewProductClient
}

func NewCartUseCase(repository interfaces.CartRepository, productRepo clinetinterface.NewProductClient) services.CartUseCase {

	return &cartUseCase{
		cartRepository:    repository,
		productRepository: productRepo,
	}

}
func (cr *cartUseCase) AddToCart(product_id, user_id, quantity int) (models.CartResponse, error) {
	ok, err := cr.productRepository.CheckProduct(product_id)
	if err != nil {
		return models.CartResponse{}, err
	}
	if !ok {
		return models.CartResponse{}, errors.New("product Does not exist")
	}
	QuantityOfProductInCart, err := cr.cartRepository.QuantityOfProductInCart(user_id, product_id)
	if err != nil {

		return models.CartResponse{}, err
	}
	quantityOfProduct, err := cr.productRepository.GetQuantityFromProductID(product_id)
	if err != nil {

		return models.CartResponse{}, err
	}
	if quantityOfProduct <= 0 {
		return models.CartResponse{}, errors.New("out of stock")
	}
	if quantityOfProduct == QuantityOfProductInCart {
		return models.CartResponse{}, errors.New("stock limit exceeded")
	}
	ok, err = cr.cartRepository.ProductExist(user_id, product_id)
	if err != nil {
		return models.CartResponse{}, err
	}
	if !ok {
		return models.CartResponse{}, err
	}
	productPrice, err := cr.productRepository.GetPriceofProductFromID(product_id)
	if err != nil {

		return models.CartResponse{}, err
	}
	if !ok {
		return models.CartResponse{}, err
	}

	FinalPrice := productPrice * float64(quantity)
	if QuantityOfProductInCart == 0 {
		err := cr.cartRepository.AddItemIntoCart(user_id, product_id, quantity, FinalPrice)
		if err != nil {

			return models.CartResponse{}, err
		}

	} else {
		currentTotal, err := cr.cartRepository.TotalPriceForProductInCart(user_id, product_id)
		if err != nil {
			return models.CartResponse{}, err
		}
		err = cr.cartRepository.UpdateCart(QuantityOfProductInCart+quantity, currentTotal+productPrice, user_id, product_id)
		if err != nil {
			return models.CartResponse{}, err
		}
	}
	cartDetails, err := cr.cartRepository.DisplayCart(user_id)
	if err != nil {
		return models.CartResponse{}, err
	}
	cartTotal, err := cr.cartRepository.GetTotalPrice(user_id)
	if err != nil {

		return models.CartResponse{}, err
	}
	err = cr.productRepository.ProductStockMinus(product_id, QuantityOfProductInCart)
	if err != nil {
		return models.CartResponse{}, err
	}
	return models.CartResponse{
		TotalPrice: cartTotal.TotalPrice,
		Cart:       cartDetails,
	}, nil

}
func (cr *cartUseCase) DisplayCart(user_id int) (models.CartResponse, error) {
	cart, err := cr.cartRepository.DisplayCart(user_id)
	if err != nil {
		return models.CartResponse{}, err
	}
	cartTotal, err := cr.cartRepository.GetTotalPrice(user_id)
	if err != nil {
		return models.CartResponse{}, err
	}
	return models.CartResponse{
		TotalPrice: cartTotal.TotalPrice,
		Cart:       cart,
	}, nil
}

func (cr *cartUseCase) GetAllItemsFromCart(userID int) ([]models.Cart, error) {
	res, err := cr.cartRepository.GetAllItemsFromCart(userID)
	if err != nil {
		return []models.Cart{}, err
	}
	return res, nil
}
func (cr *cartUseCase) DoesCartExist(userID int) (bool, error) {
	res, err := cr.cartRepository.DoesCartExist(userID)
	if err != nil {
		return false, err
	}
	return res, nil
}
func (cr *cartUseCase) TotalAmountInCart(userID int) (float64, error) {
	res, err := cr.cartRepository.TotalAmountInCart(userID)
	if err != nil {
		return 0.0, err
	}
	return res, nil
}
func (cr *cartUseCase) UpdateCartAfterOrder(userID, productID int, quantity float64) error {
	err := cr.cartRepository.UpdateCartAfterOrder(userID, productID, quantity)
	if err != nil {
		return err
	}
	return nil
}
