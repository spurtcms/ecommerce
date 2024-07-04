package ecommerce

import (
	"fmt"

	"strings"
	"time"
)

// pass limit , offset get productslist
func (ecommerce *Ecommerce) ProductsList(offset int, limit int, filter Filter) (productlists []TblEcomProduct, totalcount int64, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return []TblEcomProduct{}, 0, AuthErr
	}

	Ecommercemodel.DataAccess = ecommerce.DataAccess

	Ecommercemodel.UserId = ecommerce.UserId

	productlist, _, err := Ecommercemodel.ProductList(offset, limit, filter, ecommerce.DB)

	_, count, _ := Ecommercemodel.ProductList(0, 0, filter, ecommerce.DB)

	if err != nil {

		fmt.Println(err)
	}
	var finalproduct []TblEcomProduct

	for _, val := range productlist {

		imgs := strings.Split(val.ProductImagePath, ",")

		if len(imgs) > 0 {

			val.ProductImagePath = imgs[0]

			finalproduct = append(finalproduct, val)
		}
	}

	return finalproduct, count, nil

}

// Create Product

func (ecommerce *Ecommerce) CreateProduct(Pc CreateProductReq) (crtproduct TblEcomProduct, errr error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return TblEcomProduct{}, AuthErr
	}

	var product TblEcomProduct

	product.CategoriesId = Pc.CategoriesId

	product.ProductPrice = Pc.ProductPrice

	product.ProductDescription = Pc.ProductDescription

	product.Tax = Pc.Tax

	product.Totalcost = Pc.Totalcost

	product.IsActive = Pc.IsActive

	product.Sku = Pc.Sku

	product.ProductSlug = Pc.ProductSlug

	product.ProductName = Pc.ProductName

	product.ProductVimeoPath = Pc.ProductVimeoPath

	product.ProductYoutubePath = Pc.ProductYoutubePath

	product.ProductImagePath = Pc.ProductImagePath

	product.Stock = Pc.Stock

	product.CreatedBy = Pc.CreatedBy

	product.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	cproduct, err := Ecommercemodel.ProductCreate(product, ecommerce.DB)

	if err != nil {
		return TblEcomProduct{}, err
	}

	return cproduct, nil
}

// pass product id  get particular product details

func (ecommerce *Ecommerce) EditProduct(productid int) (products TblEcomProduct, discountprice []TblEcomProductPricing, price []TblEcomProductPricing, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return TblEcomProduct{}, []TblEcomProductPricing{}, []TblEcomProductPricing{}, AuthErr
	}

	product, err := Ecommercemodel.ProductDetailsByProductId(productid, ecommerce.DB)

	product.Imgpath = strings.Split(product.ProductImagePath, ",")

	if err != nil {

		return TblEcomProduct{}, []TblEcomProductPricing{}, []TblEcomProductPricing{}, err
	}

	var discount []TblEcomProductPricing

	var special []TblEcomProductPricing

	offers, err1 := Ecommercemodel.ProductPricingByProductId(product.Id, ecommerce.DB)

	for _, val := range offers {

		layout := "2006-01-02"

		if !val.StartDate.IsZero() {

			val.Startdate = val.StartDate.Format(layout)

		}
		if !val.EndDate.IsZero() {

			val.Enddate = val.EndDate.Format(layout)

		}

		if val.Type == "discount" {

			discount = append(discount, val)
		}
		if val.Type == "special" {

			special = append(special, val)
		}
	}

	if err1 != nil {

		return TblEcomProduct{}, []TblEcomProductPricing{}, []TblEcomProductPricing{}, err
	}

	return product, discount, special, nil

}

// pass Product id and pass update Product details

func (ecommerce *Ecommerce) UpdateProduct(Pc CreateProductReq, removeoff []int) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	var product TblEcomProduct

	product.Id = Pc.ProductId

	product.CategoriesId = Pc.CategoriesId

	product.ProductPrice = Pc.ProductPrice

	product.ProductDescription = Pc.ProductDescription

	product.Totalcost = Pc.Totalcost

	product.Tax = Pc.Tax

	product.ProductSlug = Pc.ProductSlug

	product.IsActive = Pc.IsActive

	product.Sku = Pc.Sku

	product.ProductName = Pc.ProductName

	product.ProductVimeoPath = Pc.ProductVimeoPath

	product.ProductYoutubePath = Pc.ProductYoutubePath

	product.ProductImagePath = Pc.ProductImagePath

	product.Stock = Pc.Stock

	product.ModifiedBy = Pc.ModifiedBy

	product.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Ecommercemodel.UpdateProducts(product, ecommerce.DB)

	if err != nil {
		return err
	}

	var price TblEcomProductPricing

	price.DeletedBy = Pc.ModifiedBy

	price.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err3 := Ecommercemodel.RemoveOffers(price, removeoff, ecommerce.DB)

	if err3 != nil {

		return err3

	}

	return nil
}

// pass multiple Product id soft delete the particular record

func (ecommerce *Ecommerce) MultiDeleteProduct(productid []int, id int) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	var product TblEcomProduct

	product.DeletedBy = id

	product.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Ecommercemodel.DeleteSelectedProducts(product, productid, ecommerce.DB)

	if err != nil {

		return err
	}

	return nil

}

// pass Product id soft delete the particular record

func (ecommerce *Ecommerce) DeleteProduct(productid int, id int) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	var product TblEcomProduct

	product.Id = productid

	product.DeletedBy = id

	product.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Ecommercemodel.DeleteSingleProducts(product, ecommerce.DB)

	if err != nil {

		return err
	}

	return nil

}

// check sku name already exists
func (ecommerce *Ecommerce) CheckSkuName(sku string, id int) (bool, error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return false, AuthErr
	}

	var product TblEcomProduct

	flg, err := Ecommercemodel.SkuNameCheck(product, sku, id, ecommerce.DB)

	if err != nil {

		return false, err
	}

	return flg, nil

}

// product status change
func (ecommerce *Ecommerce) SelectProductsChangeStatus(status int, productid []int) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	var product TblEcomProduct

	product.IsActive = status

	product.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Ecommercemodel.SelectProductsChangeStatus(productid, product, ecommerce.DB)

	if err != nil {
		return err
	}

	return nil
}

// Add to product in cart

func (Ecommerce *Ecommerce) AddToCart(cart EcommerceCart) (boolean bool, err error) {

	var count int64

	count, err = EcommerceModel.GetCartCount(EcommerceModel{}, cart.CustomerID, cart.ProductID, Ecommerce.DB)
	if err != nil {
		return false, err
	}

	err = EcommerceModel.AddToCart(EcommerceModel{}, count, cart, Ecommerce.DB)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Get Product details
func (Ecommerce *Ecommerce) GetProduct(productId int, productSlug string) (product EcommerceProduct, err error) {

	if AuthErr := AuthandPermission(Ecommerce); AuthErr != nil {

		return EcommerceProduct{}, AuthErr
	}

	product, err = Ecommercemodel.GetProduct(productId, productSlug, Ecommerce.DB)
	if err != nil {
		return EcommerceProduct{}, err
	}

	return product, nil
}

// Get cart list
func (Ecommerce *Ecommerce) GetCartListById(customerId, limit, offset int) (cartList []EcommerceProduct, err error) {

	if AuthErr := AuthandPermission(Ecommerce); AuthErr != nil {

		return []EcommerceProduct{}, AuthErr
	}

	cartList, err = EcommerceModel.GetCartListById(EcommerceModel{}, customerId, limit, offset, Ecommerce.DB)
	if err != nil {
		return []EcommerceProduct{}, err
	}

	return cartList, err
}

// Remove product forom cart list
func (Ecommerce *Ecommerce) RemoveProductFromCartlist(productId int, memberId int) (err error) {

	if AuthErr := AuthandPermission(Ecommerce); AuthErr != nil {

		return AuthErr
	}

	err = EcommerceModel.RemoveProductFromCartlist(EcommerceModel{}, productId, memberId, Ecommerce.DB)
	if err != nil {
		return err
	}

	return nil
}

// Get product order list
func (Ecommerce *Ecommerce) GetProductOrdersList(filter ProductFilter, sort ProductSort, customerId int, limit int, offset int) (orderedProductList []EcommerceProduct, count int64, err error) {

	if AuthErr := AuthandPermission(Ecommerce); AuthErr != nil {

		return []EcommerceProduct{}, -1, AuthErr
	}

	orderedProductList, count, err = EcommerceModel.GetProductOrdersList(EcommerceModel{}, filter, sort, customerId, limit, offset, Ecommerce.DB)
	if err != nil {
		return []EcommerceProduct{}, -1, err
	}

	return orderedProductList, count, nil
}

// update product count
func (Ecommerce *Ecommerce) UpdateProductViewCount(productId int, productSlug string) (err error) {

	if AuthErr := AuthandPermission(Ecommerce); AuthErr != nil {

		return AuthErr
	}

	err = EcommerceModel.UpdateProductViewCount(EcommerceModel{}, productId, productSlug, Ecommerce.DB)
	if err != nil {

		return err
	}

	return nil
}

// Get Cart list cunt by id

func (Ecommerce *Ecommerce) GetCartListCountById(customerId int) (count int64, err error) {

	if AuthErr := AuthandPermission(Ecommerce); AuthErr != nil {

		return -1, AuthErr
	}

	count, err = EcommerceModel.GetCartListCountById(EcommerceModel{}, customerId, Ecommerce.DB)
	if err != nil {
		return -1, err
	}

	return count, nil

}

// Get Product Details By ID
func (Ecommerce *Ecommerce) GetProductdetailsById(productId int, productSlug string) (product EcommerceProduct, err error) {

	if AuthErr := AuthandPermission(Ecommerce); AuthErr != nil {

		return EcommerceProduct{}, AuthErr
	}

	product, err = EcommerceModel.GetProductdetailsById(EcommerceModel{}, productId, productSlug, Ecommerce.DB)
	if err != nil {

		return EcommerceProduct{}, err
	}

	return product, nil
}

// Get Product Order Details By Id
func (Ecommerce *Ecommerce) GetProductOrderDetailsById(productId int, productSlug string, customerId int, orderId int) (product EcommerceProduct, productOrderStatus []OrderStatus, err error) {

	if AuthErr := AuthandPermission(Ecommerce); AuthErr != nil {

		return EcommerceProduct{}, []OrderStatus{}, AuthErr
	}

	product, productOrderStatus, err = EcommerceModel.GetProductDetailsAndOrderStatus(EcommerceModel{}, productId, productSlug, customerId, orderId, Ecommerce.DB)
	if err != nil {
		return EcommerceProduct{}, []OrderStatus{}, err
	}

	return product, productOrderStatus, nil
}

func (ecommerce *Ecommerce) CreateProductPricing(pricing TblEcomProductPricing) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	err1 := Ecommercemodel.CreateProductPricing(pricing, ecommerce.DB)

	if err1 != nil {
		return err1
	}
	return nil

}
func (ecommerce *Ecommerce) UpdateProductPricing(pricing TblEcomProductPricing) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	err1 := Ecommercemodel.UpdateProductPricing(pricing, ecommerce.DB)

	if err1 != nil {
		return err1
	}
	return nil

}
