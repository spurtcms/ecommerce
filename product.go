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
func (ecommerce *Ecommerce) ProductsList(limit int, offset int, filter Filter) (productlists []TblEcomProducts, totalcount int64, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return []TblEcomProducts{}, 0, AuthErr
	}

	productlist, _, err := Ecommercemodel.ProductList(limit, offset, filter, ecommerce.DB)

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

	product.Status = Pc.Status

	product.Sku = Pc.Sku

	product.ProductName = Pc.ProductName

	product.ProductVimeoPath = Pc.ProductVimeoPath

	product.ProductYoutubePath = Pc.ProductYoutubePath

	product.ProductImagePath = Pc.ProductImagePath

	product.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	cproduct,err := Ecommercemodel.ProductCreate(product, ecommerce.DB)

	if err != nil {
		return err
	}

	var pricing TblEcomProductPricings

	pricing.ProductId = cproduct.Id

	pricing.Priority = Pc.Priority

	pricing.StartDate = Pc.StartDate

	pricing.EndDate = Pc.EndDate

	pricing.Type = Pc.Type

	err1 := Ecommercemodel.CreateProductPricing(pricing, ecommerce.DB)

	if err1 != nil {
		return err1
	}

	return nil
}

// pass product id  get particular product details

func (ecommerce *Ecommerce) EditProduct(productid int) (products TblEcomProducts, price []TblEcomProductPricings, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return TblEcomProducts{}, []TblEcomProductPricings{}, AuthErr
	}

	product, err := Ecommercemodel.ProductDetailsByProductId(productid, ecommerce.DB)

	if err != nil {

		return TblEcomProducts{}, []TblEcomProductPricings{}, err
	}

	var pricings []TblEcomProductPricings

	pricing, err1 := Ecommercemodel.ProductPricingByProductId(productid, ecommerce.DB)

	if err1 != nil {

		return TblEcomProducts{}, []TblEcomProductPricings{}, err1
	}

	for _, pricevalue := range pricing {

		layout := "2006-01-02T15:04"

		if !pricevalue.StartDate.IsZero() {

			pricevalue.Startdate = pricevalue.StartDate.Format(layout)

		}
		if !pricevalue.EndDate.IsZero() {

			pricevalue.Enddate = pricevalue.EndDate.Format(layout)

		}

		pricings = append(pricings, pricevalue)
	}

	return product, pricings, nil

}

// pass Product id and pass update Product details

func (ecommerce *Ecommerce) UpdateProduct(Productid int, offerid int, removeoff []int, Pc CreateProductReq) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	var product TblEcomProducts

	product.Id = Productid

	product.CategoriesId = Pc.CategoriesId

	product.ProductPrice = Pc.ProductPrice

	product.ProductDescription = Pc.ProductDescription

	product.Tax = Pc.Tax

	product.Status = Pc.Status

	product.Sku = Pc.Sku

	product.ProductName = Pc.ProductName

	product.ProductVimeoPath = Pc.ProductVimeoPath

	product.ProductYoutubePath = Pc.ProductYoutubePath

	product.ProductImagePath = Pc.ProductImagePath

	product.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Ecommercemodel.UpdateProducts(product, ecommerce.DB)

	if err != nil {
		return err
	}

	var pricing TblEcomProductPricings

	pricing.Id = offerid

	pricing.ProductId = Productid

	pricing.Priority = Pc.Price

	pricing.StartDate = Pc.StartDate

	pricing.EndDate = Pc.EndDate

	pricing.Type = Pc.Type

	if offerid != 0 {
		err1 := Ecommercemodel.UpdateProductPricing(pricing, ecommerce.DB)

		if err1 != nil {
			return err1
		}
	}

	if offerid == 0 {

		err2 := Ecommercemodel.CreateProductPricing(pricing, ecommerce.DB)

		if err2 != nil {
			return err2
		}

	}

	var price TblEcomProductPricings

	price.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err3 := Ecommercemodel.RemoveOffers(price, removeoff, ecommerce.DB)

	if err3 != nil {

		return err3

	}

	return nil
}

// pass Product id soft delete the particular record

func (ecommerce *Ecommerce) DeleteProduct(productid []int) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	var product TblEcomProducts

	product.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Ecommercemodel.DeleteSelectedProducts(product, productid, ecommerce.DB)

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

		return flg, err
	}

	return flg, nil

}

// product status change
func (ecommerce *Ecommerce) SelectProductsChangeStatus(status int, productid []int) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	var product TblEcomProducts

	product.Status = status

	product.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Ecommercemodel.SelectProductsChangeStatus(productid, product, ecommerce.DB)

	if err != nil {
		return err
	}

	return nil
}
