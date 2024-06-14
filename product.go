package ecommerce

import (
	"fmt"

	"strings"
	"time"
)

type tblEcomProducts struct {
	Id                 int `gorm:"primaryKey;auto_increment;type:serial"`
	CategoriesId       string
	ProductName        string
	ProductDescription string
	ProductImagePath   string
	ProductYoutubePath string
	ProductVimeoPath   string
	Sku                string
	ProductPrice       int
	Tax                int
	Totalcost          int
	Priority           int       `gorm:"-:migration;<-:false"`
	Price              int       `gorm:"-:migration;<-:false"`
	StartDate          time.Time `gorm:"-:migration;<-:false"`
	EndDate            time.Time `gorm:"-:migration;<-:false"`
	Type               string    `gorm:"-:migration;<-:false"`
	Quantity           int       `gorm:"-:migration;<-:false"`
	Order_id           int       `gorm:"-:migration;<-:false"`
	Product_id         int       `gorm:"-:migration;<-:false"`
	Quantityprice      int       `gorm:"-:migration;<-:false"`
	Status             int
	CreatedOn          time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy          int       `gorm:"DEFAULT:NULL"`
	ModifiedOn         time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy         int       `gorm:"DEFAULT:NULL"`
	IsDeleted          int       `gorm:"DEFAULT:0"`
	DeletedOn          time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy          int       `gorm:"type:integer;DEFAULT:NULL"`
	Imgpath            []string  `gorm:"-"`
}

type tblEcomProductPricings struct {
	Id        int `gorm:"primaryKey;auto_increment;type:serial"`
	ProductId int
	Priority  int
	Price     int
	StartDate time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	EndDate   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	Type      string
	IsDeleted int       `gorm:"DEFAULT:0"`
	Startdate string    `gorm:"-:migration;<-:false"`
	Enddate   string    `gorm:"-:migration;<-:false"`
	DeletedOn time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy int       `gorm:"type:integer;DEFAULT:NULL"`
}

// pass limit , offset get productslist
func (ecommerce *Ecommerce) ProductsList(offset int, limit int, filter Filter) (productlists []TblEcomProducts, totalcount int64, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return []TblEcomProducts{}, 0, AuthErr
	}

	productlist, _, err := Ecommercemodel.ProductList(offset, limit, filter, ecommerce.DB)

	_, count, _ := Ecommercemodel.ProductList(0, 0, filter, ecommerce.DB)

	if err != nil {

		fmt.Println(err)
	}
	var finalproduct []TblEcomProducts

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

func (ecommerce *Ecommerce) CreateProduct(Pc CreateProductReq) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	var product TblEcomProducts

	product.CategoriesId = Pc.CategoriesId

	product.ProductPrice = Pc.ProductPrice

	product.ProductDescription = Pc.ProductDescription

	product.Tax = Pc.Tax

	product.Totalcost = Pc.Totalcost

	product.IsActive = Pc.IsActive

	product.Sku = Pc.Sku

	product.ProductName = Pc.ProductName

	product.ProductVimeoPath = Pc.ProductVimeoPath

	product.ProductYoutubePath = Pc.ProductYoutubePath

	product.ProductImagePath = Pc.ProductImagePath

	product.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	cproduct, err := Ecommercemodel.ProductCreate(product, ecommerce.DB)

	if err != nil {
		return err
	}

	if Pc.Priority != 0 {

		var pricing TblEcomProductPricings

		pricing.ProductId = cproduct.Id

		pricing.Priority = Pc.Priority

		pricing.StartDate = Pc.StartDate

		pricing.EndDate = Pc.EndDate

		pricing.Type = Pc.Type

		pricing.Price = Pc.Price

		err1 := Ecommercemodel.CreateProductPricing(pricing, ecommerce.DB)

		if err1 != nil {
			return err1
		}
	}

	return nil
}

// pass product id  get particular product details

func (ecommerce *Ecommerce) EditProduct(productid int) (products TblEcomProducts, discountprice []TblEcomProductPricings, price []TblEcomProductPricings, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return TblEcomProducts{}, []TblEcomProductPricings{}, []TblEcomProductPricings{}, AuthErr
	}

	product, err := Ecommercemodel.ProductDetailsByProductId(productid, ecommerce.DB)

	product.Imgpath = strings.Split(product.ProductImagePath, ",")

	if err != nil {

		return TblEcomProducts{}, []TblEcomProductPricings{}, []TblEcomProductPricings{}, err
	}

	var discount []TblEcomProductPricings

	var special []TblEcomProductPricings

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

		return TblEcomProducts{}, []TblEcomProductPricings{}, []TblEcomProductPricings{}, err
	}
	fmt.Println("product", product)
	return product, discount, special, nil

}

// pass Product id and pass update Product details

func (ecommerce *Ecommerce) UpdateProduct(Pc CreateProductReq, removeoff []int, userid int) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	var product TblEcomProducts

	product.Id = Pc.ProductId

	product.CategoriesId = Pc.CategoriesId

	product.ProductPrice = Pc.ProductPrice

	product.ProductDescription = Pc.ProductDescription

	product.Tax = Pc.Tax

	product.IsActive = Pc.IsActive

	product.Sku = Pc.Sku

	product.ProductName = Pc.ProductName

	product.ProductVimeoPath = Pc.ProductVimeoPath

	product.ProductYoutubePath = Pc.ProductYoutubePath

	product.ProductImagePath = Pc.ProductImagePath

	product.ModifiedBy = userid

	product.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Ecommercemodel.UpdateProducts(product, ecommerce.DB)

	if err != nil {
		return err
	}

	var pricing TblEcomProductPricings

	pricing.Id = Pc.PricingId

	pricing.ProductId = Pc.ProductId

	pricing.Priority = Pc.Price

	pricing.StartDate = Pc.StartDate

	pricing.EndDate = Pc.EndDate

	pricing.Type = Pc.Type

	if Pc.PricingId != 0 {
		err1 := Ecommercemodel.UpdateProductPricing(pricing, ecommerce.DB)

		if err1 != nil {
			return err1
		}
	}

	if Pc.PricingId == 0 {

		err2 := Ecommercemodel.CreateProductPricing(pricing, ecommerce.DB)

		if err2 != nil {
			return err2
		}

	}

	var price TblEcomProductPricings

	price.DeletedBy = userid

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

	var product TblEcomProducts

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

	var product TblEcomProducts

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

	var product TblEcomProducts

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

	var product TblEcomProducts

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
