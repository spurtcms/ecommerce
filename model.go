package ecommerce

import (
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
	TenantId         int
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
	Stock              int
	IsActive           int
	PricingId          int
	ProductSlug        string
	TenantId           int
}

type CreateSettingReq struct {
	Id              int
	StoreName       string
	CurrencyDefault int
	PaymentDefault  int
	StatusDefault   int
	DisplayStock    int
	StockWarning    int
	StockCheckout   int
	CreatedBy       int
	ModifiedBy      int
	TenantId        int
}

type CreateCurrencyReq struct {
	Id              int
	CurrencyName    string
	CurrencyType    string
	CurrencySymbol  string
	CurrencyDefault int
	CreatedBy       int
	ModifiedBy      int
	IsActive        int
	TenantId        int
}

type CreateStatusReq struct {
	Id          int
	Status      string
	Description string
	Priority    int
	ColorCode   string
	CreatedBy   int
	ModifiedBy  int
	IsActive    int
	TenantId    int
}

type CustomerDetails struct {
	Id               int
	FirstName        string
	LastName         string
	MobileNo         string
	Email            string
	Username         string
	Password         string
	IsActive         int
	ProfileImage     string
	ProfileImagePath string
	CreatedOn        time.Time
	CreatedBy        int
	ModifiedOn       time.Time
	IsDeleted        int
	DeletedOn        time.Time
	ModifiedBy       int
	HouseNo          string
	Area             string
	City             string
	State            string
	Country          string
	ZipCode          string
	StreetAddress    string
	MemberID         int
	TenantId         int
}

type CreatePaymentReq struct {
	Id           int
	PaymentName  string
	Description  string
	PaymentImage string
	CreatedBy    int
	ModifiedBy   int
	IsActive     int
	TenantId     int
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
	OrderStatus  int
	PriceRange   string
}

type EcommerceCart struct {
	ID         int
	ProductID  int
	CustomerID int
	Quantity   int
	CreatedOn  time.Time
	ModifiedOn *time.Time
	IsDeleted  int
	DeletedOn  *time.Time
	TenantId   int
}

type ProductFilter struct {
	Status, SearchKeyword, OrderId, StartingDate, EndingDate string
	StartingPrice, EndingPrice, OrderHistory, UpcomingOrders int
}

type ProductSort struct {
	Price int
	Date  int
}

type OrderStatusNames struct {
	Id          int
	Status      string
	Priority    int
	ColorCode   string
	Description string
	IsActive    int
	CreatedBy   int
	CreatedOn   time.Time
	ModifiedBy  int
	ModifiedOn  time.Time
	IsDeleted   int
}

type EcommerceOrder struct {
	Id              int
	OrderId         string
	CustomerId      int
	Status          string
	ShippingAddress string
	IsDeleted       int
	CreatedOn       time.Time
	ModifiedOn      time.Time
	Price           int
	Tax             int
	TotalCost       int
}

type OrderProduct struct {
	Id        int
	OrderId   int
	ProductId int
	Quantity  int
	Price     int
	Tax       int
}

type OrderStatus struct {
	Id          int
	OrderId     int
	OrderStatus int
	CreatedBy   int
	CreatedOn   time.Time
}

type OrderPayment struct {
	Id          int
	OrderId     int
	PaymentMode string
}

type EcommerceProduct struct {
	ID                 int
	CategoriesID       int
	ProductName        string
	ProductSlug        string
	ProductDescription string
	ProductImagePath   string
	ProductYoutubePath string
	ProductVimeoPath   string
	Sku                string
	Tax                int
	Totalcost          int
	IsActive           int
	CreatedOn          time.Time
	CreatedBy          int
	ModifiedOn         time.Time
	ModifiedBy         int
	IsDeleted          int
	DeletedBy          int
	DeletedOn          time.Time
	ViewCount          int
	DefaultPrice       int
	DiscountPrice      int
	SpecialPrice       int
	ProductImageArray  []string
	EcommerceCart      EcommerceCart
	OrderID            int
	OrderUniqueID      string
	OrderQuantity      int
	OrderPrice         int
	OrderTax           int
	OrderStatus        string
	OrderCustomer      int
	OrderTime          time.Time
	PaymentMode        string
	ShippingDetails    string
}



type TblEcomProductOrderDetail struct {
	Id         int `gorm:"primaryKey;auto_increment;type:serial"`
	Order_id   int `gorm:"type:integer"`
	Product_id int `gorm:"type:integer"`
	Quantity   int `gorm:"type:integer"`
	Price      int `gorm:"type:integer"`
	Tax        int `gorm:"type:int"`
	TenantId   int `gorm:"type:integer"`
}

type TblEcomCurrency struct {
	Id              int    `gorm:"primaryKey;auto_increment;type:serial"`
	CurrencyName    string `gorm:"type:character varying"`
	CurrencyType    string `gorm:"type:character varying"`
	CurrencySymbol  string `gorm:"type:character varying"`
	IsActive        int    `gorm:"type:integer"`
	CurrencyDefault int    `gorm:"type:integer"`
	CreatedOn       time.Time
	CreatedBy       int
	ModifiedOn      time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy      int       `gorm:"DEFAULT:NULL"`
	IsDeleted       int       `gorm:"type:integer"`
	DeletedOn       time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy       int       `gorm:"type:integer"`
	DateString      string    `gorm:"-"`
	TenantId        int       `gorm:"type:integer"`
}

type TblEcomStatus struct {
	Id          int    `gorm:"primaryKey;auto_increment;type:serial"`
	Status      string `gorm:"type:character varying"`
	Description string `gorm:"type:character varying"`
	IsActive    int    `gorm:"type:integer"`
	Priority    int    `gorm:"type:integer"`
	ColorCode   string `gorm:"type:character varying"`
	CreatedOn   time.Time
	CreatedBy   int
	ModifiedOn  time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy  int       `gorm:"DEFAULT:NULL"`
	IsDeleted   int       `gorm:"type:integer"`
	DeletedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy   int       `gorm:"type:integer"`
	TenantId    int       `gorm:"type:integer"`
}

type TblEcomPayment struct {
	Id           int    `gorm:"primaryKey;auto_increment;type:serial"`
	PaymentName  string `gorm:"type:character varying"`
	Description  string `gorm:"type:character varying"`
	PaymentImage string `gorm:"type:character varying"`
	IsActive     int    `gorm:"type:integer"`
	CreatedOn    time.Time
	CreatedBy    int
	ModifiedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy   int       `gorm:"DEFAULT:NULL"`
	IsDeleted    int       `gorm:"type:integer"`
	DeletedOn    time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy    int       `gorm:"type:integer"`
	TenantId     int       `gorm:"type:integer"`
}

type TblEcomSettings struct {
	Id              int       `gorm:"primaryKey;auto_increment;type:serial"`
	StoreName       string    `gorm:"type:character varying"`
	DisplayStock    int       `gorm:"type:integer"`
	StockWarning    int       `gorm:"type:integer"`
	StockCheckout   int       `gorm:"type:integer"`
	CurrencyDefault int       `gorm:"type:integer"`
	PaymentDefault  int       `gorm:"type:integer"`
	StatusDefault   int       `gorm:"type:integer"`
	CreatedOn       time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedBy       int
	ModifiedOn      time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedBy      int       `gorm:"DEFAULT:NULL"`
	TenantId        int       `gorm:"type:integer"`
}

type TblEcomOrderPayment struct {
	Id          int    `gorm:"primaryKey;auto_increment;type:serial"`
	OrderId     int    `gorm:"type:integer"`
	PaymentMode string `gorm:"type:character varying"`
	TenantId    int    `gorm:"type:integer"`
}

type TblEcomProductPricing struct {
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
	TenantId  int       `gorm:"type:integer"`
}

// pass limit , offset get customerlist
func (ecommerceModel EcommerceModel) CustomersList(offset int, limit int, filter Filter, DB *gorm.DB, tenantid int) (customer []TblEcomCustomer, totalcustomer int64, err error) {

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

	query := DB.Table("tbl_ecom_customers").Select("tbl_ecom_customers.*,count( tbl_ecom_product_orders.customer_id)").Joins("left join tbl_ecom_product_orders on tbl_ecom_customers.id = tbl_ecom_product_orders.customer_id AND tbl_ecom_product_orders.is_deleted = 0").Where("tbl_ecom_customers.is_deleted=? and tbl_ecom_customers.tenant_id=?", 0, tenantid).Group("tbl_ecom_customers.id, tbl_ecom_product_orders.customer_id").Order("tbl_ecom_customers.id desc")

	if filter.Keyword != "" {

		if filter.Keyword == "0" || filter.Keyword == "1" {

			query = query.Where(" tbl_ecom_customers.is_active=?", filter.Keyword)

		} else {

			query = query.Where("(LOWER(TRIM(tbl_ecom_customers.first_name)) LIKE LOWER(TRIM(?)) OR LOWER(TRIM(tbl_ecom_customers.email)) LIKE LOWER(TRIM(?)))", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%")

		}
	}

	if filter.IsActive != "" {

		query = query.Where("tbl_ecom_customers.is_active=?", filter.IsActive)

	}
	if filter.MemberId != 0 {

		query = query.Where("tbl_ecom_customers.member_id=?", filter.MemberId)

	}

	if filter.FirstName != "" {

		query = query.Where("LOWER(TRIM(first_name)) LIKE LOWER(TRIM(?))", "%"+filter.FirstName+"%")

	}

	if ecommerceModel.DataAccess == 1 {

		query = query.Where("tbl_ecom_customers.created_by=?", ecommerceModel.UserId)
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

func (ecommerceModel EcommerceModel) CustomerCreate(customer TblEcomCustomer, DB *gorm.DB) error {

	if err := DB.Model(TblEcomCustomer{}).Create(&customer).Error; err != nil {
		return err
	}
	return nil

}

// pass customer id  get particular customer details

func (ecommerceModel EcommerceModel) CustomerEdit(id int, DB *gorm.DB, tenantid int) (customer TblEcomCustomer, err error) {

	if err := DB.Model(TblEcomCustomer{}).Select("tbl_ecom_customers.*, count( tbl_ecom_product_orders.customer_id)").Joins("left join tbl_ecom_product_orders on tbl_ecom_customers.id = tbl_ecom_product_orders.customer_id").Group("tbl_ecom_customers.id, tbl_ecom_product_orders.customer_id").Where(" tbl_ecom_customers.is_deleted = 0 AND tbl_ecom_customers.id=? and tbl_ecom_customers.tenant_id=?", id, tenantid).First(&customer).Error; err != nil {

		return TblEcomCustomer{}, err
	}
	return customer, nil
}

// pass customer id and pass update customer details

func (ecommerceModel EcommerceModel) CustomerUpdate(customer TblEcomCustomer, DB *gorm.DB) error {

	query := DB.Model(TblEcomCustomer{}).Where("member_id=?", customer.MemberId)

	if customer.Password == "" && customer.ProfileImage == "" && customer.ProfileImagePath == "" {

		query.Omit("password, profile_image , profile_image_path").UpdateColumns(map[string]interface{}{"first_name": customer.FirstName, "last_name": customer.LastName, "email": customer.Email, "username": customer.Username, "mobile_no": customer.MobileNo, "is_active": customer.IsActive, "modified_on": customer.ModifiedOn, "modified_by": customer.ModifiedBy, "street_address": customer.StreetAddress, "city": customer.City, "country": customer.Country, "state": customer.State, "zip_code": customer.ZipCode})

	} else {

		query.UpdateColumns(map[string]interface{}{"first_name": customer.FirstName, "last_name": customer.LastName, "email": customer.Email, "username": customer.Username, "mobile_no": customer.MobileNo, "is_active": customer.IsActive, "modified_on": customer.ModifiedOn, "modified_by": customer.ModifiedBy, "street_address": customer.StreetAddress, "city": customer.City, "country": customer.Country, "state": customer.State, "zip_code": customer.ZipCode, "password": customer.Password, "profile_image": customer.ProfileImage, "profile_image_path": customer.ProfileImagePath})
	}
	return nil

}

// pass customer id soft delete the particular record

func (ecommerceModel EcommerceModel) CustomerDelete(customer TblEcomCustomer, DB *gorm.DB, tenantid int) error {

	if err := DB.Model(TblEcomCustomer{}).Where("member_id=? and tenant_id=?", customer.MemberId, tenantid).UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_by": customer.DeletedBy, "deleted_on": customer.DeletedOn}).Error; err != nil {

		return err
	}

	return nil

}

// Order list
func (ecommerceModel EcommerceModel) OrderList(offset int, limit int, filter Filter, DB *gorm.DB, tenantid int) (tblorders []TblEcomProductOrder, ordercount int64, err error) {

	query := DB.Table("tbl_ecom_product_orders").Select("tbl_ecom_statuses.status as status_value,tbl_ecom_statuses.color_code as status_color,tbl_ecom_product_orders.*,tbl_ecom_customers.username").Joins("inner join tbl_ecom_customers on tbl_ecom_product_orders.customer_id = tbl_ecom_customers.id").Joins("inner join tbl_ecom_statuses on tbl_ecom_product_orders.order_status = tbl_ecom_statuses.id").Where("tbl_ecom_product_orders.is_deleted = 0 and tbl_ecom_product_orders.tenant_id=?", tenantid)

	if filter.Keyword != "" {

		query = query.Where("LOWER(TRIM(tbl_ecom_customers.username)) LIKE LOWER(TRIM(?))"+" OR tbl_ecom_product_orders.uuid =?"+" OR LOWER(TRIM(tbl_ecom_statuses.status::text)) LIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%", filter.Keyword, "%"+filter.Keyword+"%")

	}

	if filter.Customername != "" {

		query = query.Where("LOWER(TRIM(tbl_ecom_customers.username)) LIKE LOWER(TRIM(?))", "%"+filter.Customername+"%")

	}

	if filter.Orderid != "" {

		query = query.Where("LOWER(TRIM(tbl_ecom_product_orders.uuid)) LIKE LOWER(TRIM(?))", "%"+filter.Orderid+"%")

	}

	if filter.OrderStatus != 0 {

		query = query.Where("tbl_ecom_product_orders.order_status = ?", filter.OrderStatus)

	}

	if ecommerceModel.DataAccess == 1 {

		query = query.Where("tbl_ecom_product_orders.created_by = ?", ecommerceModel.UserId)

	}

	if limit != 0 {

		query.Limit(limit).Offset(offset).Order("tbl_ecom_product_orders.id desc").Find(&tblorders)

		return tblorders, 0, err

	}

	query.Find(&tblorders).Count(&ordercount)

	return tblorders, ordercount, nil
}

// Delete the  order
func (ecommerceModel EcommerceModel) OrderDelete(productorder TblEcomProductOrder, DB *gorm.DB, tenantid int) error {

	if err := DB.Model(TblEcomProductOrder{}).Where("id=? and tenant_id=?", productorder.Id, tenantid).UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_by": productorder.DeletedBy, "deleted_on": productorder.DeletedOn}).Error; err != nil {

		return err

	}

	return nil
}

// Get product details by pass order id

func (ecommerceModel EcommerceModel) GetProductdetailsByOrderId(orderid int, DB *gorm.DB, tenantid int) (orders []TblEcomProductOrderDetail, err error) {

	if err := DB.Model(TblEcomProductOrderDetail{}).Where("order_id=? and tenant_id=?", orderid, tenantid).Find(&orders).Error; err != nil {

		return []TblEcomProductOrderDetail{}, err
	}

	return orders, nil
}

// Get Product details pass produt id

func (ecommerceModel EcommerceModel) GetProductdetailsByProductId(productid []int, DB *gorm.DB, tenantid int) (product []TblEcomProduct, err error) {

	if err := DB.Model(TblEcomProduct{}).Where("is_deleted=0 and id in (?)n and tenant_id=?", productid, tenantid).Find(&product).Error; err != nil {

		return []TblEcomProduct{}, err
	}

	return product, nil

}

// Create order status

func (ecommerceModel EcommerceModel) CreateOrderStatus(status TblEcomOrderStatus, DB *gorm.DB) error {

	if err := DB.Table("tbl_ecom_order_statuses").Create(&status).Error; err != nil {
		return err
	}
	return nil

}

// update order status in product table

func (ecommerceModel EcommerceModel) OrderStatusUpdate(orderstatus TblEcomProductOrder, DB *gorm.DB, tenantid int) error {

	if err := DB.Model(TblEcomProductOrder{}).Where("id=? and tenant_id=?", orderstatus.Id, tenantid).Updates(TblEcomProductOrder{OrderStatus: orderstatus.OrderStatus, ModifiedOn: orderstatus.ModifiedOn}).Error; err != nil {

		return err
	}

	return nil
}

// edit the order details

func (ecommerceModel EcommerceModel) OrderEdit(id string, DB *gorm.DB, tenantid int) (productord TblEcomProductOrder, err error) {

	if err := DB.Table("tbl_ecom_product_orders").Select("tbl_ecom_statuses.status as status_value,tbl_ecom_statuses.priority as status_priority,tbl_ecom_statuses.color_code as status_color,tbl_ecom_product_orders.*,tbl_ecom_customers.username,tbl_ecom_customers.mobile_no,tbl_ecom_customers.email,tbl_ecom_customers.street_address,tbl_ecom_customers.city,tbl_ecom_customers.state,tbl_ecom_customers.country,tbl_ecom_customers.zip_code").Joins("inner join tbl_ecom_customers on tbl_ecom_product_orders.customer_id = tbl_ecom_customers.id").Joins("inner join tbl_ecom_statuses on tbl_ecom_product_orders.order_status = tbl_ecom_statuses.id").Where("uuid=? and tenant_id=?", id, tenantid).First(&productord).Error; err != nil {

		return TblEcomProductOrder{}, err
	}
	return productord, nil

}

// Get order status details in particular

func (ecommerceModel EcommerceModel) OrderStatusDetails(id string, DB *gorm.DB, tenantid int) (status []TblEcomOrderStatus, err error) {

	if err := DB.Table("tbl_ecom_order_statuses").Select("tbl_ecom_order_statuses.order_status,tbl_ecom_order_statuses.order_id,tbl_ecom_statuses.priority as status_priority,tbl_ecom_order_statuses.created_on").Joins("inner join tbl_ecom_product_orders on tbl_ecom_product_orders.id=tbl_ecom_order_statuses.order_id").Joins("inner join tbl_ecom_statuses on tbl_ecom_statuses.id = tbl_ecom_order_statuses.order_status").Where("tbl_ecom_product_orders.uuid=? and tbl_ecom_product_orders.tenant_id=?", id, tenantid).Find(&status).Error; err != nil {

		return []TblEcomOrderStatus{}, err
	}
	return status, nil

}

// MULTISELECT ORDER DELETE
func (ecommerceModel EcommerceModel) MultiSelectDeleteOrder(order TblEcomProductOrder, id []int, DB *gorm.DB, tenantid int) error {

	if err := DB.Model(TblEcomProductOrder{}).Where("id in (?) and tenant_id=?", id, tenantid).UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_on": order.DeletedOn, "deleted_by": order.DeletedBy}).Error; err != nil {

		return err

	}

	return nil
}

// Product list

func (ecommerceModel EcommerceModel) ProductList(offset int, limit int, filter Filter, DB *gorm.DB, tenantid int) (tblproducts []TblEcomProduct, productcount int64, err error) {

	query := DB.Model(&TblEcomProduct{}).
		Table("tbl_ecom_products AS p").
		Select("p.*", "COALESCE((SELECT price FROM tbl_ecom_product_pricings WHERE product_id = p.id AND type = 'special' AND is_deleted=0 AND (end_date > CURRENT_DATE OR end_date IS NULL) AND (start_date <= CURRENT_DATE OR start_date IS NULL) ORDER BY CASE WHEN start_date <= CURRENT_DATE THEN 1 WHEN priority = 1 THEN 2 ELSE 3 END, priority LIMIT 1), (SELECT price FROM tbl_ecom_product_pricings WHERE product_id = p.id AND type = 'discount' AND is_deleted=0 AND (end_date > CURRENT_DATE OR end_date IS NULL) AND (start_date <= CURRENT_DATE OR start_date IS NULL) ORDER BY CASE WHEN start_date <= CURRENT_DATE THEN 1 WHEN priority = 1 THEN 2 ELSE 3 END, priority LIMIT 1), p.totalcost) AS price").Where("is_deleted = 0 and tenant_id=?", tenantid).Limit(1)

	if ecommerceModel.DataAccess == 1 {

		query = query.Where("p.created_by = ? and p.tenant_id=?", ecommerceModel.UserId, tenantid)
	}

	if filter.Status == "InActive" {

		filter.Status = "0"

	} else if filter.Status == "Active" {

		filter.Status = "1"

	}

	if strings.ToLower(filter.Keyword) == "active" {

		filter.Keyword = "1"

	} else if strings.ToLower(filter.Keyword) == "inactive" {

		filter.Keyword = "0"
	}

	if filter.Keyword != "" {

		if filter.Keyword == "1" || filter.Keyword == "0" {

			query = query.Where("p.is_active=?", filter.Keyword)

		} else {

			query = query.Where("LOWER(TRIM(p.product_name)) LIKE LOWER(TRIM(?)) OR LOWER(TRIM(p.product_description)) LIKE LOWER(TRIM(?))", "%"+filter.Keyword+"%", "%"+filter.Keyword+"%")

		}

	}

	if filter.Status != "" {

		query = query.Where("p.is_active=?", filter.Status)

	}

	if filter.PriceRange != "" {

		if filter.PriceRange == "High to Low" {

			query = query.Order("price desc")
		}

		if filter.PriceRange == "Low to High" {

			query = query.Order("price asc")
		}

	}

	if filter.ProductName != "" {

		query = query.Where("LOWER(TRIM(p.product_name)) LIKE LOWER(TRIM(?))", "%"+filter.ProductName+"%")

	}

	if limit != 0 {

		query.Limit(limit).Offset(offset).Order("p.id desc").Find(&tblproducts)

	}

	query.Find(&tblproducts).Count(&productcount)

	return tblproducts, productcount, nil
}

// Create Product
func (ecommerceModel EcommerceModel) ProductCreate(product TblEcomProduct, DB *gorm.DB) (products TblEcomProduct, err error) {

	if err := DB.Table("tbl_ecom_products").Create(&product).Error; err != nil {

		return TblEcomProduct{}, err
	}

	return product, nil
}

// Create product price

func (ecommerceModel EcommerceModel) CreateProductPricing(pricing TblEcomProductPricing, DB *gorm.DB) error {

	if err := DB.Table("tbl_ecom_product_pricings").Create(&pricing).Error; err != nil {

		return err
	}

	return nil
}

// Get Product details pass product id
func (ecommerceModel EcommerceModel) ProductDetailsByProductId(productid int, DB *gorm.DB, tenantid int) (product TblEcomProduct, err error) {

	if err := DB.Table("tbl_ecom_products").Where("id=? and tenant_id=?", productid, tenantid).First(&product).Error; err != nil {

		return TblEcomProduct{}, err
	}
	return product, nil
}

// Get product price pass product id
func (ecommerceModel EcommerceModel) ProductPricingByProductId(productid int, DB *gorm.DB, tenantid int) (pricings []TblEcomProductPricing, err error) {

	if err := DB.Table("tbl_ecom_product_pricings").Where("tbl_ecom_product_pricings.product_id =? AND tbl_ecom_product_pricings.is_deleted = 0 and tbl_ecom_product_pricings.tenant_id=?", productid, tenantid).Find(&pricings).Error; err != nil {

		return []TblEcomProductPricing{}, err
	}

	return pricings, nil
}

// Update Product
func (ecommerceModel EcommerceModel) UpdateProducts(product TblEcomProduct, DB *gorm.DB, tenantid int) error {

	query := DB.Table("tbl_ecom_products").Where("id=? and tenant_id=?", product.Id, tenantid)

	if err := query.UpdateColumns(map[string]interface{}{"categories_id": product.CategoriesId, "stock": product.Stock, "product_slug": product.ProductSlug, "product_name": product.ProductName, "product_description": product.ProductDescription, "product_image_path": product.ProductImagePath, "product_vimeo_path": product.ProductVimeoPath, "sku": product.Sku, "product_youtube_path": product.ProductYoutubePath, "product_price": product.ProductPrice, "tax": product.Tax, "totalcost": product.Totalcost, "modified_on": product.ModifiedOn, "is_active": product.IsActive, "modified_by": product.ModifiedBy}).Error; err != nil {

		return err
	}

	return nil
}

// Update product price

func (ecommerceModel EcommerceModel) UpdateProductPricing(pricing TblEcomProductPricing, DB *gorm.DB, tenantid int) error {

	if err := DB.Table("tbl_ecom_product_pricings").Where("id=? and tenant_id=?", pricing.Id, tenantid).UpdateColumns(map[string]interface{}{"product_id": pricing.ProductId, "priority": pricing.Priority, "price": pricing.Price, "start_date": pricing.StartDate, "end_date": pricing.EndDate, "type": pricing.Type}).Error; err != nil {

		return err
	}

	return nil
}

// Delete offers
func (ecommerceModel EcommerceModel) RemoveOffers(price TblEcomProductPricing, deloffers []int, DB *gorm.DB, tenantid int) error {

	if err := DB.Model(TblEcomProductPricing{}).Where("id IN (?) and tenant_id=?", deloffers, tenantid).UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_by": price.DeletedBy, "deleted_on": price.DeletedOn}).Error; err != nil {

		return err
	}

	return nil

}

// Delete Selected Products

func (ecommerceModel EcommerceModel) DeleteSelectedProducts(product TblEcomProduct, productIds []int, DB *gorm.DB, tenantid int) error {

	if err := DB.Model(TblEcomProduct{}).Where("id IN (?) and tenant_id=?", productIds, tenantid).UpdateColumns(map[string]interface{}{"deleted_on": product.DeletedOn, "deleted_by": product.DeletedBy, "is_deleted": 1}).Error; err != nil {

		return err
	}

	return nil
}

// Delete single Products

func (ecommerceModel EcommerceModel) DeleteSingleProducts(product TblEcomProduct, DB *gorm.DB, tenantid int) error {

	if err := DB.Model(TblEcomProduct{}).Where("id = ? and tenant_id=?", product.Id, tenantid).UpdateColumns(map[string]interface{}{"deleted_on": product.DeletedOn, "deleted_by": product.DeletedBy, "is_deleted": 1}).Error; err != nil {

		return err
	}

	return nil
}

// Check sku name already exists

func (ecommerceModel EcommerceModel) SkuNameCheck(product TblEcomProduct, skuname string, productid int, DB *gorm.DB, tenantid int) (bool, error) {

	if productid == 0 {
		if err := DB.Model(TblEcomProduct{}).Where("LOWER(TRIM(sku))=LOWER(TRIM(?)) and is_deleted=0 and tenant_id=?", skuname, tenantid).First(&product).Error; err != nil {

			return false, err
		}
	} else {
		if err := DB.Model(TblEcomProduct{}).Where("LOWER(TRIM(sku))=LOWER(TRIM(?)) and id not in (?) and is_deleted = 0 and tenant_id=?", skuname, productid, tenantid).First(&product).Error; err != nil {

			return false, err
		}
	}

	return true, nil

}

// selected product status change

func (ecommerceModel EcommerceModel) SelectProductsChangeStatus(productIds []int, product TblEcomProduct, DB *gorm.DB, tenantid int) error {

	if err := DB.Model(TblEcomProduct{}).Where("id IN (?) and tenant_id=?", productIds, tenantid).UpdateColumns(map[string]interface{}{"is_active": product.IsActive, "modified_on": product.ModifiedOn}).Error; err != nil {

		return err
	}

	return nil
}

// Change active state in multiple customer

func (ecommerceModel EcommerceModel) MultiSelectCustomerIsactive(customer TblEcomCustomer, customerid []int, DB *gorm.DB, tenantid int) error {

	if err := DB.Table("tbl_ecom_customers").Where("member_id in (?) and tenant_id=?", customerid, tenantid).UpdateColumns(map[string]interface{}{"is_active": customer.IsActive, "modified_on": customer.ModifiedOn}).Error; err != nil {

		return err
	}

	return nil

}

// Get order details pass uuid

func (ecommerceModel EcommerceModel) GetOrderDetailsbyuuid(uuid string, DB *gorm.DB, tenantid int) (order TblEcomProductOrder, err error) {

	if err := DB.Model(TblEcomProductOrder{}).Select("tbl_ecom_customers.*,tbl_ecom_product_orders.*,tbl_ecom_statuses.status as status_value,tbl_ecom_statuses.color_code as status_color,tbl_ecom_statuses.priority as status_priority").Joins("left join tbl_ecom_customers on tbl_ecom_product_orders.customer_id = tbl_ecom_customers.id").Joins("inner join tbl_ecom_statuses on tbl_ecom_product_orders.order_status = tbl_ecom_statuses.id").Where("tbl_ecom_product_orders.is_deleted = 0 and uuid=? and tbl_ecom_product_orders.tenant_id=?", uuid, tenantid).First(&order).Error; err != nil {

		return TblEcomProductOrder{}, err
	}
	return order, nil

}

// Get Customer details pass customerid
func (ecommerceModel EcommerceModel) GetCustomerDetails(id int, DB *gorm.DB, tenantid int) (customer TblEcomCustomer, err error) {

	if err := DB.Table("tbl_ecom_customers").Select("tbl_ecom_customers.*, count( tbl_ecom_product_orders.customer_id)").Joins("left join tbl_ecom_product_orders on tbl_ecom_customers.id = tbl_ecom_product_orders.customer_id").Group("tbl_ecom_customers.id, tbl_ecom_product_orders.customer_id").Where(" tbl_ecom_customers.is_deleted = 0 AND tbl_ecom_customers.id=? and tbl_ecom_customers.tenant_id=?", id, tenantid).First(&customer).Error; err != nil {

		return TblEcomCustomer{}, err
	}
	return customer, nil
}

// Get order details pass customer id
func (ecommerceModel EcommerceModel) GetOrderDetailsbyCustomerId(limit, offset int, customerid int, DB *gorm.DB, tenantid int) (order []TblEcomProductOrder, totalorder int64, err error) {

	query := DB.Table("tbl_ecom_product_orders").Select("tbl_ecom_product_orders.*,tbl_ecom_statuses.status as status_value,tbl_ecom_statuses.color_code as status_color").Joins("inner join tbl_ecom_customers on tbl_ecom_product_orders.customer_id = tbl_ecom_customers.id").Joins("inner join tbl_ecom_statuses on tbl_ecom_product_orders.order_status = tbl_ecom_statuses.id").Where("tbl_ecom_product_orders.is_deleted = 0 and tbl_ecom_product_orders.customer_id=? and tbl_ecom_product_orders.tenant_id=?", customerid, tenantid).Find(&order)

	if limit != 0 {

		query.Offset(offset).Limit(limit).Order("tbl_ecom_product_orders.id desc").Find(&order)

		return order, 0, err

	} else {

		query.Find(&order).Count(&totalorder)

		return order, totalorder, err
	}

}

// Multiple delete
func (ecommerceModel EcommerceModel) MultiSelectDeleteCustomers(customer TblEcomCustomer, customerid []int, DB *gorm.DB, tenantid int) (flg bool, err error) {

	if err := DB.Table("tbl_ecom_customers").Where("member_id in (?) and tenant_id=?", customerid, tenantid).UpdateColumns(map[string]interface{}{"is_deleted": customer.IsDeleted, "deleted_by": customer.DeletedBy, "deleted_on": customer.DeletedOn}).Error; err != nil {

		return false, err
	}

	return true, nil

}

// Get Product details
func (ecommerceModel EcommerceModel) GetProduct(productId int, productSlug string, DB *gorm.DB, tenantid int) (product EcommerceProduct, err error) {

	if productId != 0 {

		if err := DB.Table("tbl_ecom_products").Where("is_deleted = 0 and id = ? and tenant_id=?", productId, tenantid).Find(&product).Error; err != nil {

			return EcommerceProduct{}, err
		}

	}

	if productSlug != "" {

		if err := DB.Table("tbl_ecom_products").Where("is_deleted = 0 and product_slug = ? and tenant_id=?", productSlug, tenantid).Find(&product).Error; err != nil {

			return EcommerceProduct{}, err
		}

	}

	return product, nil
}

// Get Customer details
func (ecommerceModel EcommerceModel) GetCustomer(memberId int, DB *gorm.DB, tenantid int) (customer TblEcomCustomer, err error) {

	if err := DB.Table("tbl_ecom_customers").Where("tbl_ecom_customers.is_deleted = 0 and tbl_ecom_customers.member_id = ? and tbl_ecom_customers.tenant_id=?", memberId, tenantid).Find(&customer).Error; err != nil {

		return TblEcomCustomer{}, err
	}

	return customer, nil
}

// Get Cart Count
func (ecommerce EcommerceModel) GetCartCount(customerId int, productId int, DB *gorm.DB, tenantid int) (count int64, err error) {

	result := DB.Table("tbl_ecom_carts").Where("is_deleted = 0 and customer_id = ? and product_id = ? and tenant_id=?", customerId, productId, tenantid).Count(&count)
	if result.Error != nil {
		return -1, result.Error
	}

	return count, nil
}

// Add to product in cart
func (ecommerceModel EcommerceModel) AddToCart(count int64, cart EcommerceCart, DB *gorm.DB, tenantid int) error {

	query := DB.Table("tbl_ecom_carts")

	if count > 0 {
		query = query.Where("is_deleted = 0 and customer_id = ? and product_id = ? and tenant_id=?", cart.CustomerID, cart.ProductID, tenantid).UpdateColumns(map[string]interface{}{"quantity": gorm.Expr("quantity + ?", cart.Quantity), "modified_on": cart.ModifiedOn})

	} else {

		query = query.Create(&cart)
	}

	if query.Error != nil {
		return query.Error
	}

	return nil
}

// Get Cart list
func (ecommerce EcommerceModel) GetCartListById(customerId int, limit int, offset int, DB *gorm.DB, tenantid int) (cartList []EcommerceProduct, err error) {

	result := DB.Table("tbl_ecom_products").Select("tbl_ecom_products.*,rp.price AS discount_price ,rs.price AS special_price,tbl_ecom_carts.*").Joins("inner join tbl_ecom_carts on tbl_ecom_carts.product_id =  tbl_ecom_products.id ").Joins("left join (select *, ROW_NUMBER() OVER (PARTITION BY tbl_ecom_product_pricings.id, tbl_ecom_product_pricings.type ORDER BY tbl_ecom_product_pricings.priority,tbl_ecom_product_pricings.start_date desc) AS rn from tbl_ecom_product_pricings where tbl_ecom_product_pricings.type ='discount' and tbl_ecom_product_pricings.start_date <= now() and tbl_ecom_product_pricings.end_date >= now()) rp on rp.product_id = tbl_ecom_products.id").Joins("left join (select *, ROW_NUMBER() OVER (PARTITION BY tbl_ecom_product_pricings.id, tbl_ecom_product_pricings.type ORDER BY tbl_ecom_product_pricings.priority,tbl_ecom_product_pricings.start_date desc) AS rn from tbl_ecom_product_pricings where tbl_ecom_product_pricings.type ='special' and tbl_ecom_product_pricings.start_date <= now() and tbl_ecom_product_pricings.end_date >= now()) rs on rs.product_id = tbl_ecom_products.id").Joins("inner join tbl_ecom_customers on tbl_ecom_customers.id = tbl_ecom_carts.customer_id").
		Where("tbl_ecom_carts.is_deleted = 0 and tbl_ecom_products.is_deleted = 0 and tbl_ecom_customers.is_deleted = 0 and tbl_ecom_products.is_active = 1 and tbl_ecom_customers.id = ? and tbl_ecom_products.tenant_id = ? ", customerId, tenantid).Preload("EcommerceCart").Limit(limit).Offset(offset).Order("tbl_ecom_carts.id desc").Find(&cartList)

	if result.Error != nil {
		return []EcommerceProduct{}, result.Error
	}

	return cartList, nil
}

// Get Cart list count by id

func (ecommerce EcommerceModel) GetCartListCountById(customerId int, DB *gorm.DB, tenantid int) (count int64, err error) {

	result := DB.Table("tbl_ecom_carts").Joins("inner join tbl_ecom_products on tbl_ecom_products.id = tbl_ecom_carts.product_id").Joins("left join (select *, ROW_NUMBER() OVER (PARTITION BY tbl_ecom_product_pricings.id, tbl_ecom_product_pricings.type ORDER BY tbl_ecom_product_pricings.priority,tbl_ecom_product_pricings.start_date desc) AS rn from tbl_ecom_product_pricings where tbl_ecom_product_pricings.type ='discount' and tbl_ecom_product_pricings.start_date <= now() and tbl_ecom_product_pricings.end_date >= now()) rp on rp.product_id = tbl_ecom_products.id").Joins("left join (select *, ROW_NUMBER() OVER (PARTITION BY tbl_ecom_product_pricings.id, tbl_ecom_product_pricings.type ORDER BY tbl_ecom_product_pricings.priority,tbl_ecom_product_pricings.start_date desc) AS rn from tbl_ecom_product_pricings where tbl_ecom_product_pricings.type ='special' and tbl_ecom_product_pricings.start_date <= now() and tbl_ecom_product_pricings.end_date >= now()) rs on rs.product_id = tbl_ecom_products.id").Joins("inner join tbl_ecom_customers on tbl_ecom_customers.id = tbl_ecom_carts.customer_id").
		Where("tbl_ecom_carts.is_deleted = 0 and tbl_ecom_products.is_deleted = 0 and tbl_ecom_customers.is_deleted = 0 and tbl_ecom_products.is_active = 1 and tbl_ecom_customers.id = ? tbl_ecom_carts.tenant_id = ? ", customerId, tenantid).Count(&count)

	if result.Error != nil {
		return -1, result.Error
	}

	return count, nil
}

// Remove product from cart list

func (ecommerce EcommerceModel) RemoveProductFromCartlist(productId int, memberId int, DB *gorm.DB, tenantid int) (err error) {

	currentTime, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	subquery := DB.Table("tbl_ecom_customers").Select("id").Where("is_deleted = 0 and member_id = ? and tenant_id=?", memberId, tenantid)

	result := DB.Table("tbl_ecom_carts").Where("tbl_ecom_carts.is_deleted = 0 and tbl_ecom_carts.product_id = ? and tbl_ecom_carts.customer_id = (?) and tbl_ecom_carts.tenant_id=?", productId, subquery, tenantid).UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_on": currentTime})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Get Product order list
func (ecommerce EcommerceModel) GetProductOrdersList(filter ProductFilter, sort ProductSort, customerId int, limit int, offset int, DB *gorm.DB, tenantid int) (productOrdersList []EcommerceProduct, count int64, err error) {

	query := DB.Table("tbl_ecom_products as p").Joins("inner join tbl_ecom_product_order_details d on d.product_id = p.id").Joins("inner join tbl_ecom_product_orders o on o.id = d.order_id").Joins("inner join tbl_ecom_order_payments op on op.order_id = o.id").Where("p.is_deleted = 0 and o.is_deleted = 0 and o.customer_id = ? and p.tenant_id=?", customerId, tenantid)

	if filter.UpcomingOrders == 1 {

		query = query.Where("o.status in (?)", []string{"placed", "outofdelivery", "shipped"})

	} else if filter.OrderHistory == 1 {

		query = query.Where("o.status in (?)", []string{"delivered", "cancelled"})

	} else if filter.Status != "" {

		query = query.Where("o.status = ?", filter.Status)
	}

	if filter.StartingPrice != 0 && filter.EndingPrice != 0 {

		query = query.Where("d.price between ? and ?", filter.StartingPrice, filter.EndingPrice)

	} else if filter.StartingPrice != 0 {

		query = query.Where("d.price >= ?", filter.StartingPrice)

	} else if filter.EndingPrice != 0 {

		query = query.Where("d.price <= ?", filter.EndingPrice)

	}

	if filter.SearchKeyword != "" {

		query = query.Where("LOWER(TRIM(p.product_name)) LIKE LOWER(TRIM(?))", "%"+filter.SearchKeyword+"%")
	}

	if filter.StartingDate != "" && filter.EndingDate != "" {

		query = query.Where("o.created_on between ? and ?", filter.StartingDate, filter.EndingDate)

	} else if filter.StartingDate != "" {

		query = query.Where("o.created_on >= ?", filter.StartingDate)

	} else if filter.EndingDate != "" {

		query = query.Where("o.created_on <= ?", filter.EndingDate)
	}

	if filter.OrderId != "" {

		query = query.Where("o.uuid = ?", filter.OrderId)
	}

	if err := query.Count(&count).Error; err != nil {

		return []EcommerceProduct{}, -1, err
	}

	if sort.Date != -1 {

		if sort.Date == 1 {

			query = query.Order("o.id desc")

		} else if sort.Date == 0 {

			query = query.Order("o.id")

		}

	} else if sort.Price != -1 {

		if sort.Price == 1 {

			query = query.Order("d.price desc")

		} else if sort.Price == 0 {

			query = query.Order("d.price")

		}

	} else {

		query = query.Order("o.id desc")
	}

	if err := query.Select("p.*,o.id,o.uuid,o.status,o.customer_id,o.created_on,o.shipping_address,d.quantity,d.price,d.tax,op.payment_mode").Limit(limit).Offset(offset).Find(&productOrdersList).Error; err != nil {

		return []EcommerceProduct{}, -1, err
	}

	return productOrdersList, count, nil
}

// update product view
func (ecommerce EcommerceModel) UpdateProductViewCount(productId int, productSlug string, DB *gorm.DB, tenantid int) (err error) {

	query := DB.Table("tbl_ecom_products").Where("is_deleted = 0 and is_active = 1 and tenant_id=?", tenantid)

	if productId != 0 && productId != -1 {

		query = query.Where("id = ?", productId)

	} else if productSlug != "" {

		query = query.Where("product_slug = ?", productSlug)
	}

	err = query.Update("view_count", gorm.Expr("view_count + 1")).Error
	if err != nil {

		return err
	}

	return nil
}

// place order

func (ecommerce EcommerceModel) PlaceOrder(orderPlaced EcommerceOrder, DB *gorm.DB) (err error) {

	if err := DB.Table("tbl_ecom_product_orders").Create(&orderPlaced).Error; err != nil {

		return err
	}

	return nil
}

// get order list
func (ecommerce EcommerceModel) GetOrderByOrderId(orderId string, DB *gorm.DB, tenantid int) (order EcommerceOrder, err error) {

	if err := DB.Table("tbl_ecom_product_orders").Where("uuid = ? and tenant_id=?", orderId, tenantid).First(&order).Error; err != nil {

		return EcommerceOrder{}, err
	}

	return order, nil
}

// create order
func (ecommerce EcommerceModel) CreateOrderDetails(orderDetails OrderProduct, DB *gorm.DB) (err error) {

	if err = DB.Table("tbl_ecom_product_order_details").Create(&orderDetails).Error; err != nil {

		return err
	}

	return nil
}

// update stock
func (ecommerce EcommerceModel) UpdateStock(productId int, quantity int, DB *gorm.DB, tenantid int) (err error) {

	if err = DB.Table("tbl_ecom_products").Where("is_deleted = 0 and is_active = 1 and id = ? and tenant_id=?", productId, tenantid).Update("stock", gorm.Expr("stock - ?", quantity)).Error; err != nil {

		return err
	}

	return nil

}

// create order payment
func (ecommerce EcommerceModel) CreateOrderPayment(orderpayment OrderPayment, DB *gorm.DB) (err error) {

	if err := DB.Table("tbl_ecom_order_payments").Create(&orderpayment).Error; err != nil {

		return err
	}

	return nil
}

// Delete form cart after order
func (ecommerce EcommerceModel) DeleteFromCartAfterOrder(orderedProductIds []int, customerId int, DB *gorm.DB, tenantid int) (err error) {

	currentTime, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	if err := DB.Table("tbl_ecom_carts").Where("is_deleted = 0 and product_id in (?) and customer_id = ? and tenant_id=?", orderedProductIds, customerId, tenantid).UpdateColumns(map[string]interface{}{"is_deleted": 1, "deleted_on": currentTime}).Error; err != nil {

		return err
	}

	return nil
}

// Get Status list
func (ecommerce EcommerceModel) GetOrderStatusList(DB *gorm.DB, tenantid int) (orderStatus []OrderStatusNames, err error) {

	if err := DB.Table("tbl_ecom_statuses").Where("is_deleted = 0 and is_active=1 and tenant_id=?", tenantid).Order("priority").Find(&orderStatus).Error; err != nil {

		return []OrderStatusNames{}, err
	}

	return orderStatus, nil
}

// Get Store list
func (ecommerce EcommerceModel) StoreLists(DB *gorm.DB) (setting TblEcomSettings, err error) {

	if err := DB.Table("tbl_ecom_settings").Find(&setting).Error; err != nil {

		return TblEcomSettings{}, err
	}
	return setting, nil
}

// Get payment list
func (ecommerce EcommerceModel) PaymentLists(offset, limit int, DB *gorm.DB, tenantid int) (pay []TblEcomPayment, err error) {

	query := DB.Table("tbl_ecom_payments").Where("is_deleted = 0 and tenant_id=?", tenantid).Order("id desc")

	if ecommerce.DataAccess == 1 {

		query = query.Where("tbl_ecom_payments.created_by = ?", ecommerce.UserId)
	}

	if limit != 0 {
		query.Limit(limit).Offset(offset).Find(&pay)
		return pay, nil
	}

	return pay, nil
}

// Get Order list
func (ecommerce EcommerceModel) CurrencyLists(offset, limit int, DB *gorm.DB, tenantid int) (money []TblEcomCurrency, err error) {

	query := DB.Table("tbl_ecom_currencies").Where("is_deleted = 0 and tenant_id=?", tenantid).Order("id desc")

	if ecommerce.DataAccess == 1 {

		query = query.Where("tbl_ecom_currencies.created_by = ?", ecommerce.UserId)
	}

	if limit != 0 {

		query.Limit(limit).Offset(offset).Find(&money)
		return money, nil
	}

	return money, nil
}

// Get Status list
func (ecommerce EcommerceModel) StatusLists(offset, limit int, DB *gorm.DB, tenantid int) (status []TblEcomStatus, err error) {

	query := DB.Table("tbl_ecom_statuses").Where("is_deleted =0 and tenant_id=?", tenantid).Order("priority")

	if ecommerce.DataAccess == 1 {

		query = query.Where("tbl_ecom_statuses.created_by = ?", ecommerce.UserId)
	}

	if limit != 0 {

		query.Limit(limit).Offset(offset).Find(&status)
		return status, nil
	}
	return status, nil
}

// Create Setting
func (ecommerce EcommerceModel) CreateSetting(setting TblEcomSettings, DB *gorm.DB) error {

	if err := DB.Table("tbl_ecom_settings").Create(&setting).Error; err != nil {
		return err
	}
	return nil
}

// Update Setting
func (ecommerce EcommerceModel) UpdateSetting(setting TblEcomSettings, DB *gorm.DB, tenantid int) error {

	if err := DB.Table("tbl_ecom_settings").Where("id=? and tenant_id=?", setting.Id, tenantid).UpdateColumns(map[string]interface{}{"store_name": setting.StoreName, "currency_default": setting.CurrencyDefault, "payment_default": setting.PaymentDefault, "status_default": setting.StatusDefault, "display_stock": setting.DisplayStock, "stock_warning": setting.StockWarning, "stock_checkout": setting.StockCheckout, "modified_on": setting.ModifiedOn, "modified_by": setting.ModifiedBy}).Error; err != nil {

		return err
	}
	return nil
}

// Create Currency
func (ecommerce EcommerceModel) CurrencyCreate(money TblEcomCurrency, DB *gorm.DB) error {

	if err := DB.Table("tbl_ecom_currencies").Create(&money).Error; err != nil {
		return err
	}
	return nil
}

// Create Currency
func (ecommerce EcommerceModel) FindDefault(DB *gorm.DB) (money TblEcomCurrency, err error) {

	if err := DB.Table("tbl_ecom_currencies").Where("currency_default = 1 and is_deleted=0").First(&money).Error; err != nil {

		return TblEcomCurrency{}, err
	}

	return money, nil
}

// Update Default values is 0

func (ecommerce EcommerceModel) ChangeDefaultValue(id int, DB *gorm.DB, tenantid int) error {

	if err := DB.Table("tbl_ecom_currencies").Where("id = ? and tenant_id=?", id, tenantid).UpdateColumns(map[string]interface{}{"currency_default": 0}).Error; err != nil {
		return err
	}
	return nil
}

// Update Currency

func (ecommerce EcommerceModel) UpdateCurrency(money TblEcomCurrency, DB *gorm.DB, tenantid int) error {

	if err := DB.Table("tbl_ecom_currencies").Where("id = ? and tenant_id=?", money.Id, tenantid).UpdateColumns(map[string]interface{}{"currency_name": money.CurrencyName, "currency_type": money.CurrencyType, "currency_symbol": money.CurrencySymbol, "currency_default": money.CurrencyDefault, "modified_on": money.ModifiedOn, "modified_by": money.ModifiedBy}).Error; err != nil {
		return err
	}

	return nil
}

// Delete Currency

func (ecommerce EcommerceModel) CurrencyDelete(money TblEcomCurrency, DB *gorm.DB, tenantid int) error {

	if err := DB.Table("tbl_ecom_currencies").Where("id = ? and tenant_id=?", money.Id, tenantid).UpdateColumns(map[string]interface{}{"is_deleted": money.IsDeleted, "deleted_on": money.DeletedOn, "deleted_by": money.DeletedBy}).Error; err != nil {
		return err
	}
	return nil

}

// Currency Isactive

func (ecommerce EcommerceModel) InActiveCurrency(money TblEcomCurrency, DB *gorm.DB, tenantid int) (bool, error) {

	if err := DB.Table("tbl_ecom_currencies").Where("id = ? and tenant_id=?", money.Id, tenantid).UpdateColumns(map[string]interface{}{"is_active": money.IsActive, "modified_on": money.ModifiedOn, "modified_by": money.ModifiedBy}).Error; err != nil {
		return false, err
	}
	return true, nil

}

// Create Status
func (ecommerce EcommerceModel) CreateStatus(status TblEcomStatus, DB *gorm.DB) error {

	if err := DB.Table("tbl_ecom_statuses").Create(&status).Error; err != nil {
		return err
	}
	return nil
}

// Update Currency

func (ecommerce EcommerceModel) UpdateStatus(status TblEcomStatus, DB *gorm.DB, tenantid int) error {

	if err := DB.Table("tbl_ecom_statuses").Where("id = ? and tenant_id=?", status.Id, tenantid).UpdateColumns(map[string]interface{}{"status": status.Status, "description": status.Description, "priority": status.Priority, "color_code": status.ColorCode, "modified_on": status.ModifiedOn, "modified_by": status.ModifiedBy}).Error; err != nil {
		return err
	}
	return nil
}

// Delete Status

func (ecommerce EcommerceModel) DeleteStatus(status TblEcomStatus, DB *gorm.DB, tenantid int) error {

	if err := DB.Table("tbl_ecom_statuses").Where("id = ? and tenant_id=?", status.Id, tenantid).UpdateColumns(map[string]interface{}{"is_deleted": status.IsDeleted, "deleted_on": status.DeletedOn, "deleted_by": status.DeletedBy}).Error; err != nil {
		return err
	}
	return nil

}

// Status Isactive

func (ecommerce EcommerceModel) OrderStatusIsActive(status TblEcomStatus, DB *gorm.DB, tenantid int) (bool, error) {

	if err := DB.Table("tbl_ecom_statuses").Where("id = ? and tenant_id=?", status.Id, tenantid).UpdateColumns(map[string]interface{}{"is_active": status.IsActive, "modified_on": status.ModifiedOn, "modified_by": status.ModifiedBy}).Error; err != nil {
		return false, err
	}
	return true, nil

}

// Create Payment
func (ecommerce EcommerceModel) PaymentCreate(pay TblEcomPayment, DB *gorm.DB) error {

	if err := DB.Table("tbl_ecom_payments").Create(&pay).Error; err != nil {
		return err
	}
	return nil
}

// Update Payment

func (ecommerce EcommerceModel) UpdatePayment(pay TblEcomPayment, DB *gorm.DB, tenantid int) error {

	if err := DB.Table("tbl_ecom_payments").Where("id = ? and tenant_id=?", pay.Id, tenantid).UpdateColumns(map[string]interface{}{"payment_name": pay.PaymentName, "payment_image": pay.PaymentImage, "description": pay.Description, "modified_on": pay.ModifiedOn, "modified_by": pay.ModifiedBy}).Error; err != nil {
		return err
	}
	return nil
}

// Delete payment

func (ecommerce EcommerceModel) DeletePayment(payment TblEcomPayment, DB *gorm.DB, tenantid int) error {

	if err := DB.Table("tbl_ecom_payments").Where("id = ? and tenant_id=?", payment.Id, tenantid).UpdateColumns(map[string]interface{}{"is_deleted": payment.IsDeleted, "deleted_on": payment.DeletedOn, "deleted_by": payment.DeletedBy}).Error; err != nil {
		return err
	}
	return nil

}

// Payment Isactive

func (ecommerce EcommerceModel) PaymentIsActive(pay TblEcomPayment, DB *gorm.DB, tenantid int) (bool, error) {

	if err := DB.Table("tbl_ecom_payments").Where("id = ? and tenant_id=?", pay.Id, tenantid).UpdateColumns(map[string]interface{}{"is_active": pay.IsActive, "modified_on": pay.ModifiedOn, "modified_by": pay.ModifiedBy}).Error; err != nil {

		return false, err
	}
	return true, nil

}

// Edit Currency

func (ecommerce EcommerceModel) CurrencyGet(id int, DB *gorm.DB, tenantid int) (currency TblEcomCurrency, err error) {

	if err := DB.Table("tbl_ecom_currencies").Where("id = ? and tenant_id=?", id, tenantid).Find(&currency).Error; err != nil {

		return TblEcomCurrency{}, err
	}
	return currency, nil
}

// Edit Payment
func (ecommerce EcommerceModel) PaymentGet(id int, DB *gorm.DB, tenantid int) (payment TblEcomPayment, err error) {

	if err := DB.Table("tbl_ecom_payments").Where("id = ? and tenant_id=?", id, tenantid).Find(&payment).Error; err != nil {

		return TblEcomPayment{}, err
	}
	return payment, nil
}

// Edit Status

func (ecommerce EcommerceModel) StatusGet(id int, DB *gorm.DB, tenantid int) (status TblEcomStatus, err error) {

	if err := DB.Table("tbl_ecom_statuses").Where("id = ? and tenant_id=?", id, tenantid).Find(&status).Error; err != nil {

		return TblEcomStatus{}, err
	}
	return status, nil
}

// Currency Name Already Exists

func (ecommerce EcommerceModel) CheckCurrencyName(id int, name string, currency TblEcomCurrency, DB *gorm.DB, tenantid int) (bool, error) {

	query := DB.Table("tbl_ecom_currencies").Where("LOWER(TRIM(currency_name))=LOWER(TRIM(?)) and is_deleted=0 and tenant_id=?", name, tenantid)

	if id != 0 {

		query = query.Where("LOWER(TRIM(currency_name))=LOWER(TRIM(?)) and id not in (?) and is_deleted=0", name, id)

	}

	if err := query.First(&currency).Error; err != nil {

		return false, err
	}

	return true, nil
}

// Currency Name Already Exists

func (ecommerce EcommerceModel) CheckCurrencyType(id int, ctype string, currency TblEcomCurrency, DB *gorm.DB, tenantid int) (bool, error) {

	query := DB.Table("tbl_ecom_currencies").Where("LOWER(TRIM(currency_type))=LOWER(TRIM(?)) and is_deleted=0 and tenant_id=?", ctype, tenantid)

	if id != 0 {

		query = query.Where("LOWER(TRIM(currency_type))=LOWER(TRIM(?)) and id not in (?) and is_deleted=0", ctype, id)

	}

	if err := query.First(&currency).Error; err != nil {
		return false, err
	}

	return true, nil
}

// Currency Name Already Exists

func (ecommerce EcommerceModel) CheckCurrencySymbol(id int, currencysymbol string, currency TblEcomCurrency, DB *gorm.DB, tenantid int) (bool, error) {

	query := DB.Table("tbl_ecom_currencies").Where("LOWER(TRIM(currency_symbol))=LOWER(TRIM(?)) and is_deleted=0 and tenant_id=?", currencysymbol, tenantid)

	if id != 0 {

		query = query.Where("LOWER(TRIM(currency_symbol))=LOWER(TRIM(?)) and id not in (?) and is_deleted=0", currencysymbol, id)

	}
	if err := query.First(&currency).Error; err != nil {

		return false, err
	}

	return true, nil
}

// Get Product Details By Id

func (ecommerce EcommerceModel) GetProductdetailsById(productId int, productSlug string, DB *gorm.DB, tenantid int) (product EcommerceProduct, err error) {

	query := DB.Table("tbl_ecom_products").Select("tbl_ecom_products.*,rp.price AS discount_price ,rs.price AS special_price").Joins("left join (select *, ROW_NUMBER() OVER (PARTITION BY tbl_ecom_product_pricings.id, tbl_ecom_product_pricings.type ORDER BY tbl_ecom_product_pricings.priority,tbl_ecom_product_pricings.start_date desc) AS rn from tbl_ecom_product_pricings where tbl_ecom_product_pricings.type ='discount' and tbl_ecom_product_pricings.start_date <= now() and tbl_ecom_product_pricings.end_date >= now()) rp on rp.product_id = tbl_ecom_products.id").Joins("left join (select *, ROW_NUMBER() OVER (PARTITION BY tbl_ecom_product_pricings.id, tbl_ecom_product_pricings.type ORDER BY tbl_ecom_product_pricings.priority,tbl_ecom_product_pricings.start_date desc) AS rn from tbl_ecom_product_pricings where tbl_ecom_product_pricings.type ='special' and tbl_ecom_product_pricings.start_date <= now() and tbl_ecom_product_pricings.end_date >= now()) rs on rs.product_id = tbl_ecom_products.id").Where("tbl_ecom_products.is_deleted = 0 and tbl_ecom_products.is_active = 1 and tbl_ecom_products.tenant_id=?", tenantid)

	if productId != 0 && productId != -1 {

		query = query.Where("tbl_ecom_products.id = ?", productId)

	} else if productSlug != "" {

		query = query.Where("tbl_ecom_products.product_slug = ?", productSlug)

	}

	if err := query.First(&product).Error; err != nil {

		return EcommerceProduct{}, err
	}

	return product, err
}

// Status Name Already Exists

func (ecommerce EcommerceModel) CheckStatusName(id int, name string, status TblEcomStatus, DB *gorm.DB, tenantid int) (bool, error) {

	query := DB.Table("tbl_ecom_statuses").Where("LOWER(TRIM(status))=LOWER(TRIM(?)) and is_deleted=0 and tenant_id=?", name, tenantid)

	if id != 0 {

		query = query.Where("LOWER(TRIM(status))=LOWER(TRIM(?)) and id not in (?) and is_deleted=0", name, id)
	}
	if err := query.First(&status).Error; err != nil {

		return false, err
	}

	return true, nil
}

// Payment Name Already Exists

func (ecommerce EcommerceModel) CheckPaymentName(id int, name string, payment TblEcomPayment, DB *gorm.DB, tenantid int) (bool, error) {

	query := DB.Table("tbl_ecom_payments").Where("LOWER(TRIM(payment_name))=LOWER(TRIM(?)) and is_deleted=0 and tenant_id=?", name, tenantid)

	if id != 0 {

		query = query.Where("LOWER(TRIM(payment_name))=LOWER(TRIM(?)) and id not in (?) and is_deleted=0", name, id)
	}

	if err := query.First(&payment).Error; err != nil {

		return false, err
	}

	return true, nil
}

// Check Status Priority Already Exists

func (ecommerce EcommerceModel) CheckStatusPriority(id int, priority int, status TblEcomStatus, DB *gorm.DB, tenantid int) (bool, error) {

	query := DB.Table("tbl_ecom_statuses").Where("priority = ? and is_deleted=0 and tenant_id=?", priority, tenantid)

	if id != 0 {

		query = query.Where("priority = ? and id not in (?) and is_deleted=0", priority, id)

	}

	if err := query.First(&status).Error; err != nil {

		return false, err
	}

	return true, nil
}

// Get Product Details and Order status by Id
func (ecommerce EcommerceModel) GetProductDetailsAndOrderStatus(productId int, productSlug string, customerId int, orderId int, DB *gorm.DB, tenantid int) (product EcommerceProduct, productOrderStatus []OrderStatus, err error) {

	query := DB.Table("tbl_ecom_products as p").Joins("inner join tbl_ecom_product_order_details d on d.product_id = p.id").Joins("inner join tbl_ecom_product_orders o on o.id = d.order_id").Joins("inner join tbl_ecom_order_payments op on op.order_id = o.id").Where("p.is_deleted = 0 and o.is_deleted = 0 and o.customer_id = ? and o.id = ? and p.tenant_id=?", customerId, orderId, tenantid)

	if productId != 0 && productId != -1 {

		query = query.Where("p.id = ?", productId)
	}

	if productSlug != "" {

		query = query.Where("p.product_slug = ?", productSlug)
	}

	if err := query.Select("p.*,o.id as order_id,o.uuid as order_unique_id,o.status as order_status,o.customer_id as order_customer,o.created_on as order_time,o.shipping_address as shipping_details,d.quantity as order_quantity,d.price as order_price,d.tax as order_tax,op.payment_mode as payment_mode").First(&product).Error; err != nil {

		return EcommerceProduct{}, []OrderStatus{}, err
	}

	if err := DB.Table("tbl_ecom_order_statuses").Where("order_id = ?", orderId).Find(&productOrderStatus).Error; err != nil {

		return EcommerceProduct{}, []OrderStatus{}, err
	}

	return product, productOrderStatus, nil
}

// get Customer Details by Id
func (ecommerce EcommerceModel) GetCustomerDetailsById(memberId int, DB *gorm.DB, tenantid int) (customer CustomerDetails, err error) {

	if err := DB.Table("tbl_ecom_customers").Where("is_deleted = 0 and member_id = ? and tenant_id=?", memberId, tenantid).First(&customer).Error; err != nil {

		return CustomerDetails{}, err
	}

	return customer, err
}

// Update Member details
func (ecommerce EcommerceModel) UpdateMemberDetails(memberId int, memberDetails map[string]interface{}, DB *gorm.DB, tenantid int) (err error) {

	if err := DB.Table("tbl_members").Where("is_deleted = 0 and id = ? and tenant_id=?", memberId, tenantid).UpdateColumns(&memberDetails).Error; err != nil {

		return err
	}

	return nil
}

// Update Customer Details
func (ecommerce EcommerceModel) UpdateCustomerDetails(memberId int, customerDetails map[string]interface{}, DB *gorm.DB, tenantid int) (err error) {

	if err := DB.Table("tbl_ecom_customers").Where("is_deleted = 0 and member_id = ? and tenant_id=?", memberId, tenantid).UpdateColumns(&customerDetails).Error; err != nil {

		return err
	}

	return nil
}
