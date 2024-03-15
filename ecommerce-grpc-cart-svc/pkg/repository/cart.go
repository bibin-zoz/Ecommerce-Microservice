package repository

import (
	interfaces "cart/service/pkg/repository/interface"
	"cart/service/pkg/utils/models"
	"errors"

	"gorm.io/gorm"
)

type cartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(DB *gorm.DB) interfaces.CartRepository {
	return &cartRepository{
		DB: DB,
	}
}

func (cr *cartRepository) QuantityOfProductInCart(userId int, productId int) (int, error) {
	var productQty int
	err := cr.DB.Raw("SELECT quantity FROM carts WHERE user_id = ? AND product_id = ?", userId, productId).Scan(&productQty).Error
	if err != nil {
		return 0, err
	}
	return productQty, nil
}
func (cr *cartRepository) AddItemIntoCart(userId int, productId int, Quantity int, productprice float64) error {
	if err := cr.DB.Exec("INSERT INTO carts (user_id,product_id,quantity,total_price) values(?,?,?,?)", userId, productId, Quantity, productprice).Error; err != nil {
		return err
	}
	return nil

}
func (cr *cartRepository) TotalPriceForProductInCart(userID int, productID int) (float64, error) {

	var totalPrice float64
	if err := cr.DB.Raw("SELECT SUM(total_price) as total_price FROM carts  WHERE user_id = ? AND product_id = ?", userID, productID).Scan(&totalPrice).Error; err != nil {
		return 0.0, err
	}
	return totalPrice, nil
}
func (cr *cartRepository) UpdateCart(quantity int, price float64, userID int, product_id int) error {
	if err := cr.DB.Exec(`UPDATE carts
	SET quantity = ?, total_price = ? 
	WHERE user_id = ? and product_id = ?`, quantity, price, product_id, userID).Error; err != nil {
		return err
	}

	return nil

}
func (cr *cartRepository) DisplayCart(userID int) ([]models.Cart, error) {

	var count int
	if err := cr.DB.Raw("SELECT COUNT(*) FROM carts WHERE user_id = ? ", userID).Scan(&count).Error; err != nil {
		return []models.Cart{}, err
	}

	if count == 0 {
		return []models.Cart{}, nil
	}

	var cartResponse []models.Cart

	if err := cr.DB.Raw("SELECT * from carts WHERE user_id = ?", userID).Scan(&cartResponse).Error; err != nil {
		return []models.Cart{}, err
	}
	return cartResponse, nil

}
func (cr *cartRepository) GetTotalPrice(userID int) (models.CartTotal, error) {

	var cartTotal models.CartTotal
	err := cr.DB.Raw("SELECT COALESCE(SUM(total_price), 0) FROM carts WHERE user_id = ?", userID).Scan(&cartTotal.TotalPrice).Error
	if err != nil {
		return models.CartTotal{}, err
	}
	err = cr.DB.Raw("SELECT COALESCE(SUM(total_price), 0) FROM carts WHERE user_id = ?", userID).Scan(&cartTotal.FinalPrice).Error
	if err != nil {
		return models.CartTotal{}, err
	}
	return cartTotal, nil

}

func (cr *cartRepository) EmptyCart(userID int) error {
	if err := cr.DB.Exec("DELETE FROM carts WHERE  user_id = ?", userID).Error; err != nil {
		return err
	}
	return nil
}

func (cr *cartRepository) ProductExist(userID int, productID int) (bool, error) {
	var count int
	if err := cr.DB.Raw("SELECT count(*) FROM carts  WHERE user_id = ? AND product_id = ?", userID, productID).Scan(&count).Error; err != nil {
		return false, err
	}
	if count == 0 {
		return true, nil
	} else {
		return false, nil
	}
}
func (cr *cartRepository) GetAllItemsFromCart(userID int) ([]models.Cart, error) {
	var count int
	var cartResponse []models.Cart
	err := cr.DB.Raw("SELECT COUNT(*) FROM carts WHERE user_id = ?", userID).Scan(&count).Error
	if err != nil {
		return []models.Cart{}, err
	}
	if count == 0 {
		return []models.Cart{}, nil
	}
	err = cr.DB.Raw("SELECT carts.user_id,carts.product_id,carts.quantity,carts.total_price from carts WHERE user_id = ?", userID).First(&cartResponse).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if len(cartResponse) == 0 {
				return []models.Cart{}, nil
			}
			return []models.Cart{}, err
		}
		return []models.Cart{}, err
	}
	return cartResponse, nil
}
func (cr *cartRepository) DoesCartExist(userID int) (bool, error) {
	var exist bool
	err := cr.DB.Raw("SELECT exists(SELECT 1 FROM carts WHERE user_id = ?)", userID).Scan(&exist).Error
	if err != nil {
		return false, err
	}
	return exist, nil
}
func (cr *cartRepository) TotalAmountInCart(userID int) (float64, error) {
	var price float64
	if err := cr.DB.Raw("SELECT SUM(total_price) FROM carts WHERE  user_id= $1", userID).Scan(&price).Error; err != nil {
		return 0, err
	}
	return price, nil
}
func (cr *cartRepository) UpdateCartAfterOrder(userID, productID int, quantity float64) error {
	err := cr.DB.Exec("DELETE FROM carts WHERE user_id = ? and product_id = ?", userID, productID).Error
	if err != nil {
		return err
	}
	return nil
}
