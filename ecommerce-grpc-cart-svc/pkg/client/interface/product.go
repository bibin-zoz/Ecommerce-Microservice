package interfaces

type NewProductClient interface {
	CheckProduct(product_id int) (bool, error)
	GetQuantityFromProductID(id int) (int, error)
	GetPriceofProductFromID(prodcut_id int) (float64, error)
	ProductStockMinus(productID, stock int) error
}
