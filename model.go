package ecommerce

import (
	"log"
	"strings"
	"time"

	"gorm.io/gorm"
)

type CreateCustomerReq struct {
	MemberId         int
	FirstName        string
	LastName         string
	Email            string
	MobileNo         string
	Username         string
	Password         string
	StreetAddress    string
	City             string
	State            string
	Country          string
	ZipCode          string
	IsActive         int
	ProfileImage     string
	ProfileImagePath string
	CreatedBy        int
	ModifiedBy       int
}

type CreateOrderReq struct {
	CustomerId      int
	Status          string
	ShippingAddress string
	Price           int
	Tax             int
	TotalCost       int
	CreatedBy       int
	ModifiedBy      int
	DeletedBy       int
}

type CreateProductReq struct {
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
	Status             int
	CreatedBy          int
	ModifiedBy         int
	DeletedBy          int
	ProductId          int
	Priority           int
	Price              int
	StartDate          time.Time
	EndDate            time.Time
	Type               string
}

type Filter struct {
	Keyword      string
	Status       string
	Orderid      string
	Customername string
	IsActive     string
	FirstName    string
	MemberId     int
	ProductName  string
	CategoryId   int
}

// pass limit , offset get customerlist
func (ecommerceModel EcommerceModel) CustomersList(limit int, offset int, filter Filter, DB *gorm.DB) (customer []TblEcomCustomers, totalcustomer int64, err error) {

	if filter.IsActive == "InActive" {

		filter.IsActive = "0"

	} else if filter.IsActive == "Active" {

		filter.IsActive = "1"

	}

	if strings.ToLower(filter.Keyword) == "active" {

		filter.Keyword = "1"

	} else if strings.ToLower(filter.Keyword) == "inactive" {

		filter.Keyword = "0"
	}

	query := DB.Debug().Table("tbl_ecom_customers").Select("tbl_ecom_customers.*,count( tbl_ecom_product_orders.customer_id)").Joins("left join tbl_ecom_product_orders on tbl_ecom_customers.id = tbl_ecom_product_orders.customer_id AND tbl_ecom_product_orders.is_deleted = 0").Where("tbl_ecom_customers.is_deleted=? ", 0).Group("tbl_ecom_customers.id, tbl_ecom_product_orders.customer_id").Order("tbl_ecom_customers.id desc")

	if filter.Keyword != "" {

		if filter.Keyword == "0" || filter.Keyword == "1" {

			query = query.Where(" tbl_ecom_customers.is_active=?", filter.Keyword)

		} else {

			query = query.Where("(LOWER(TRIM(tbl_ecom_customers.first_name)) ILIKE LOWER(TRIM(?)) OR LOWER(TRIM(tbl_ecom_customers.email)) ILIKE LOWER(TRIM(?)))", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%")

		}
	}

	if filter.IsActive != "" {

		query = query.Where("tbl_ecom_customers.is_active=?", filter.IsActive)

	}
	if filter.MemberId != 0 {

		query = query.Where("tbl_ecom_customers.member_id=?", filter.MemberId)

	}

	if filter.FirstName != "" {

		query = query.Where("LOWER(TRIM(first_name)) ILIKE LOWER(TRIM(?))", "%"+filter.FirstName+"%")

	}

	if limit != 0 {

		query.Offset(offset).Limit(limit).Order("id desc").Find(&customer)

		return customer, 0, err

	} else {

		query.Find(&customer).Count(&totalcustomer)

		return customer, totalcustomer, err
	}
}

// Create customer

func (ecommerceModel EcommerceModel) CustomerCreate(customer TblEcomCustomers, DB *gorm.DB) error {

	if err := DB.Model(TblEcomCustomers{}).Create(&customer).Error; err != nil {
		return err
	}
	return nil

}

// pass customer id  get particular customer details

func (ecommerceModel EcommerceModel) CustomerEdit(id int, DB *gorm.DB) (customer TblEcomCustomers, err error) {

	if err := DB.Model(TblEcomCustomers{}).Select("tbl_ecom_customers.*, count( tbl_ecom_product_orders.customer_id)").Joins("left join tbl_ecom_product_orders on tbl_ecom_customers.id = tbl_ecom_product_orders.customer_id").Group("tbl_ecom_customers.id, tbl_ecom_product_orders.customer_id").Where(" tbl_ecom_customers.is_deleted = 0 AND tbl_ecom_customers.id=?", id).First(&customer).Error; err != nil {

		// if err := db.Debug().Table("tbl_ecom_customers").Where("is_deleted = 0 and id=?", id).First(&customer).Error; err != nil {

		return TblEcomCustomers{}, err
	}
	return customer, nil
}

// pass customer id and pass update customer details

func (ecommerceModel EcommerceModel) CustomerUpdate(customer TblEcomCustomers, DB *gorm.DB) error {

	query := DB.Model(TblEcomCustomers{}).Where("member_id=?", customer.MemberId)

	if customer.Password == "" && customer.ProfileImage == "" && customer.ProfileImagePath == "" {

		query.Omit("password, profile_image , profile_image_path").UpdateColumns(map[string]interface{}{"first_name": customer.FirstName, "last_name": customer.LastName, "email": customer.Email, "username": customer.Username, "mobile_no": customer.MobileNo, "is_active": customer.IsActive, "modified_on": customer.ModifiedOn, "modified_by": customer.ModifiedBy, "street_address": customer.StreetAddress, "city": customer.City, "country": customer.Country, "state": customer.State, "zip_code": customer.ZipCode})

	} else {

		query.UpdateColumns(map[string]interface{}{"first_name": customer.FirstName, "last_name": customer.LastName, "email": customer.Email, "username": customer.Username, "mobile_no": customer.MobileNo, "is_active": customer.IsActive, "modified_on": customer.ModifiedOn, "modified_by": customer.ModifiedBy, "street_address": customer.StreetAddress, "city": customer.City, "country": customer.Country, "state": customer.State, "zip_code": customer.ZipCode, "password": customer.Password, "profile_image": customer.ProfileImage, "profile_image_path": customer.ProfileImagePath})
	}
	return nil

}

// pass customer id soft delete the particular record

func (ecommerceModel EcommerceModel) CustomerDelete(customer TblEcomCustomers, DB *gorm.DB) error {

	if err := DB.Model(TblEcomCustomers{}).Where("member_id=?", customer.MemberId).UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_by": customer.DeletedBy, "deleted_on": customer.DeletedOn}).Error; err != nil {

		return err
	}

	return nil

}

// Order list
func (ecommerceModel EcommerceModel) OrderList(offset int, limit int, filter Filter, DB *gorm.DB) (tblorders []TblEcomProductOrders, ordercount int64, err error) {

	query := DB.Table("tbl_ecom_product_orders").Select("tbl_ecom_product_orders.*,tbl_ecom_customers.username").Joins("inner join tbl_ecom_customers on tbl_ecom_product_orders.customer_id = tbl_ecom_customers.id").Where("tbl_ecom_product_orders.is_deleted = 0")

	if filter.Keyword != "" {

		query = query.Debug().Where("LOWER(TRIM(tbl_ecom_customers.username)) ILIKE LOWER(TRIM(?))"+" OR tbl_ecom_product_orders.uuid =?"+" OR LOWER(TRIM(tbl_ecom_order_statuses.order_status::text)) ILIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%", filter.Keyword, "%"+filter.Keyword+"%")

	}

	if filter.Customername != "" {

		query = query.Where("LOWER(TRIM(tbl_ecom_customers.username)) ILIKE LOWER(TRIM(?))", "%"+filter.Customername+"%")

	}

	if filter.Orderid != "" {

		query = query.Where("LOWER(TRIM(tbl_ecom_product_orders.uuid)) ILIKE LOWER(TRIM(?))", "%"+filter.Orderid+"%")

	}

	if filter.Status != "" {

		query = query.Where("tbl_ecom_product_orders.status = ?", filter.Status)

	}

	if limit != 0 {

		query.Limit(limit).Offset(offset).Order("tbl_ecom_product_orders.id desc").Find(&tblorders)

		return tblorders, 0, err

	}

	query.Find(&tblorders).Count(&ordercount)

	return tblorders, ordercount, nil
}

// Delete the  order
func (ecommerceModel EcommerceModel) OrderDelete(productorder TblEcomProductOrders, DB *gorm.DB) error {

	if err := DB.Model(TblEcomProductOrders{}).Where("id=?", productorder.Id).UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_by": productorder.DeletedBy, "deleted_on": productorder.DeletedOn}).Error; err != nil {

		return err

	}

	return nil
}

// Get product details by pass order id

func (ecommerceModel EcommerceModel) GetProductdetailsByOrderId(orderid int, DB *gorm.DB) (orders []TblEcomProductOrderDetails, err error) {

	if err := DB.Model(TblEcomProductOrderDetails{}).Where("order_id=?", orderid).Find(&orders).Error; err != nil {

		return []TblEcomProductOrderDetails{}, err
	}

	return orders, nil
}

// Get Product details pass produt id

func (ecommerceModel EcommerceModel) GetProductdetailsByProductId(productid []int, DB *gorm.DB) (product []TblEcomProducts, err error) {

	if err := DB.Model(TblEcomProducts{}).Where("is_deleted=0 and id in (?)", productid).Find(&product).Error; err != nil {

		return []TblEcomProducts{}, err
	}

	return product, nil

}

// Create order status

func (ecommerceModel EcommerceModel) CreateOrderStatus(status TblEcomOrderStatuses, DB *gorm.DB) error {

	if err := DB.Table("tbl_ecom_order_statuses").Create(&status).Error; err != nil {
		return err
	}
	return nil

}

// update order status in product table

func (ecommerceModel EcommerceModel) OrderStatusUpdate(orderstatus TblEcomProductOrders, DB *gorm.DB) error {

	if err := DB.Debug().Model(TblEcomProductOrders{}).Where("id=?", orderstatus.Id).Updates(TblEcomProductOrders{Status: orderstatus.Status, ModifiedOn: orderstatus.ModifiedOn}).Error; err != nil {

		return err
	}

	return nil
}

// edit the order details

func (ecommerceModel EcommerceModel) OrderEdit(id string, DB *gorm.DB) (productord TblEcomProductOrders, err error) {

	if err := DB.Preload("Orders", func(DB *gorm.DB) *gorm.DB {
		return DB.Order("id asc")
	}).Table("tbl_ecom_product_orders").Select("tbl_ecom_product_orders.*,tbl_ecom_customers.*").Joins("inner join tbl_ecom_customers on tbl_ecom_product_orders.customer_id = tbl_ecom_customers.id").Where("uuid=?", id).First(&productord).Error; err != nil {

		return TblEcomProductOrders{}, err
	}
	return productord, nil

}

// MULTISELECT ORDER DELETE
func (ecommerceModel EcommerceModel) MultiSelectDeleteOrder(order TblEcomProductOrders, id []int, DB *gorm.DB) error {

	if err := DB.Model(TblEcomProductOrders{}).Where("id in (?)", id).UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_on": order.DeletedOn, "deleted_by": order.DeletedBy}).Error; err != nil {

		return err

	}

	return nil
}

// Product list

func (ecommerceModel EcommerceModel) ProductList(limit int, offset int, filter Filter, DB *gorm.DB) (tblproducts []TblEcomProducts, productcount int64, err error) {

	if filter.Status == "InActive" {

		filter.Status = "2"

	} else if filter.Status == "Active" {

		filter.Status = "1"

	}

	if strings.ToLower(filter.Keyword) == "active" {

		filter.Keyword = "1"

	} else if strings.ToLower(filter.Keyword) == "inactive" {

		filter.Keyword = "2"
	}

	query := DB.Debug().Table("tbl_ecom_products").Where("tbl_ecom_products.is_deleted = 0")

	if filter.Keyword != "" {

		if filter.Keyword == "1" || filter.Keyword == "2" {

			query = query.Where(" tbl_ecom_products.status=?", filter.Keyword)

		} else {

			query = query.Debug().Where("LOWER(TRIM(tbl_ecom_products.product_name)) ILIKE LOWER(TRIM(?)) OR LOWER(TRIM(tbl_ecom_products.product_description)) ILIKE LOWER(TRIM(?)) ", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%")

		}

	}

	log.Println("offset", offset, limit)

	if filter.Status != "" {

		query = query.Where("tbl_ecom_products.status=?", filter.Status)

	}

	// if filter.MemberId != 0 {

	// 	query = query.Where("tbl_ecom_products.member_id=?", filter.MemberId)

	// }

	if filter.ProductName != "" {

		query = query.Where("LOWER(TRIM(product_name)) ILIKE LOWER(TRIM(?))", "%"+filter.ProductName+"%")

	}

	if limit != 0 {

		query.Limit(limit).Offset(offset).Order("tbl_ecom_products.id desc").Find(&tblproducts)

	}

	query.Find(&tblproducts).Count(&productcount)

	return tblproducts, productcount, nil
}

// Create Product
func (ecommerceModel EcommerceModel) ProductCreate(product TblEcomProducts, DB *gorm.DB) (products TblEcomProducts, err error) {

	if err := DB.Debug().Table("tbl_ecom_products").Create(&product).Error; err != nil {

		return TblEcomProducts{}, err
	}

	return product, nil
}

// Create product price

func (ecommerceModel EcommerceModel) CreateProductPricing(pricing TblEcomProductPricings, DB *gorm.DB) error {

	if err := DB.Table("tbl_ecom_product_pricings").Create(&pricing).Error; err != nil {

		return err
	}

	return nil
}

// Get Product details pass product id
func (ecommerceModel EcommerceModel) ProductDetailsByProductId(productid int, DB *gorm.DB) (product TblEcomProducts, err error) {

	if err := DB.Table("tbl_ecom_products").Where("id=?", productid).First(&product).Error; err != nil {

		return TblEcomProducts{}, err
	}
	return product, nil
}

// Get product price pass product id
func (ecommerceModel EcommerceModel) ProductPricingByProductId(productid int, DB *gorm.DB) (pricings []TblEcomProductPricings, err error) {

	if err := DB.Debug().Table("tbl_ecom_product_pricings").Where("tbl_ecom_product_pricings.product_id =? AND tbl_ecom_product_pricings.is_deleted = 0", productid).Find(&pricings).Error; err != nil {

		return []TblEcomProductPricings{}, err
	}

	return pricings, nil
}

// Update Product
func (ecommerceModel EcommerceModel) UpdateProducts(product TblEcomProducts, DB *gorm.DB) error {

	query := DB.Table("tbl_ecom_products").Where("id=?", product.Id)

	if err := query.UpdateColumns(map[string]interface{}{"categories_id": product.CategoriesId, "product_name": product.ProductName, "product_description": product.ProductDescription, "product_image_path": product.ProductImagePath, "product_vimeo_path": product.ProductVimeoPath, "sku": product.Sku, "product_youtube_path": product.ProductYoutubePath, "product_price": product.ProductPrice, "tax": product.Tax, "totalcost": product.Totalcost, "modified_on": product.ModifiedOn, "status": product.Status}).Error; err != nil {

		return err
	}

	return nil
}

// Update product price

func (ecommerceModel EcommerceModel) UpdateProductPricing(pricing TblEcomProductPricings, DB *gorm.DB) error {

	if err := DB.Debug().Table("tbl_ecom_product_pricings").Where("product_id=?", pricing.Id).UpdateColumns(map[string]interface{}{"product_id": pricing.ProductId, "priority": pricing.Priority, "price": pricing.Price, "start_date": pricing.StartDate, "end_date": pricing.EndDate, "type": pricing.Type}).Error; err != nil {

		return err
	}

	return nil
}

// Delete offers
func (ecommerceModel EcommerceModel) RemoveOffers(price TblEcomProductPricings, deloffers []int, DB *gorm.DB) error {

	if err := DB.Debug().Model(TblEcomProductPricings{}).Where("id IN (?)", deloffers).UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_on": price.DeletedOn}).Error; err != nil {

		return err
	}

	return nil

}

// Delete Selected Products

func (ecommerceModel EcommerceModel) DeleteSelectedProducts(product TblEcomProducts, productIds []int, DB *gorm.DB) error {

	if err := DB.Debug().Model(TblEcomProducts{}).Where("id IN (?)", productIds).UpdateColumns(map[string]interface{}{"deleted_on": product.DeletedOn, "is_deleted": 1}).Error; err != nil {

		return err
	}

	return nil
}

// Check sku name already exists

func (ecommerceModel EcommerceModel) SkuNameCheck(product TblEcomProducts, skuname string, productid int, DB *gorm.DB) (bool, error) {

	if productid == 0 {
		if err := DB.Debug().Model(TblEcomProducts{}).Where("LOWER(TRIM(sku))=LOWER(TRIM(?)) and is_deleted=0", skuname).First(&product).Error; err != nil {

			return false, err
		}
	} else {
		if err := DB.Model(TblEcomProducts{}).Where("LOWER(TRIM(sku))=LOWER(TRIM(?)) and id not in (?) and is_deleted = 0 ", skuname, productid).First(&product).Error; err != nil {

			return false, err
		}
	}

	return true, nil

}

// selected product status change

func (ecommerceModel EcommerceModel) SelectProductsChangeStatus(productIds []int, product TblEcomProducts, DB *gorm.DB) error {

	if err := DB.Debug().Model(TblEcomProducts{}).Where("id IN (?)", productIds).UpdateColumns(map[string]interface{}{"status": product.Status, "modified_on": product.ModifiedOn}).Error; err != nil {

		return err
	}

	return nil
}

// Change active state in multiple customer

func (ecommerceModel EcommerceModel) MultiSelectCustomerIsactive(customer TblEcomCustomers, customerid []int, DB *gorm.DB) error {

	if err := DB.Table("tbl_ecom_customers").Where("member_id in (?)", customerid).UpdateColumns(map[string]interface{}{"is_active": customer.IsActive, "modified_on": customer.ModifiedOn}).Error; err != nil {

		return err
	}

	return nil

}

// Get order details pass uuid

func (ecommerceModel EcommerceModel) GetOrderDetailsbyuuid(uuid string, DB *gorm.DB) (order TblEcomProductOrders, err error) {

	if err := DB.Preload("Orders", func(DB *gorm.DB) *gorm.DB {
		return DB.Order("id asc")
	}).Model(TblEcomProductOrders{}).Select("tbl_ecom_customers.*,tbl_ecom_product_orders.*").Joins("left join tbl_ecom_customers on tbl_ecom_product_orders.customer_id = tbl_ecom_customers.id").Where("tbl_ecom_product_orders.is_deleted = 0 and uuid=?", uuid).First(&order).Error; err != nil {

		return TblEcomProductOrders{}, err
	}
	return order, nil

}

// Get Customer details pass customerid
func (ecommerceModel EcommerceModel) GetCustomerDetails(id int, DB *gorm.DB) (customer TblEcomCustomers, err error) {

	if err := DB.Table("tbl_ecom_customers").Select("tbl_ecom_customers.*, count( tbl_ecom_product_orders.customer_id)").Joins("left join tbl_ecom_product_orders on tbl_ecom_customers.id = tbl_ecom_product_orders.customer_id").Group("tbl_ecom_customers.id, tbl_ecom_product_orders.customer_id").Where(" tbl_ecom_customers.is_deleted = 0 AND tbl_ecom_customers.id=?", id).First(&customer).Error; err != nil {

		return TblEcomCustomers{}, err
	}
	return customer, nil
}

// Get order details pass customer id
func (ecommerceModel EcommerceModel) GetOrderDetailsbyCustomerId(limit, offset int, customerid int, DB *gorm.DB) (order []TblEcomProductOrders, totalorder int64, err error) {

	query := DB.Preload("Orders", func(DB *gorm.DB) *gorm.DB {
		return DB.Order("id desc")
	}).Table("tbl_ecom_product_orders").Where("is_deleted = 0 and customer_id=?", customerid).Find(&order)

	if limit != 0 {

		query.Offset(offset).Limit(limit).Order("id desc").Find(&order)

		return order, 0, err

	} else {

		query.Find(&order).Count(&totalorder)

		return order, totalorder, err
	}

}

// Multiple delete
func (ecommerceModel EcommerceModel) MultiSelectDeleteCustomers(customer TblEcomCustomers, customerid []int, DB *gorm.DB) error {

	if err := DB.Table("tbl_ecom_customers").Where("member_id in (?)", customerid).UpdateColumns(map[string]interface{}{"is_deleted": customer.IsDeleted, "deleted_by": customer.DeletedBy, "deleted_on": customer.DeletedOn}).Error; err != nil {

		return err
	}

	return nil

}