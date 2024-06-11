package ecommerce

import (
	"encoding/json"
	"fmt"
	"log"

	"strings"
	"time"
)

type tblecomProductOrders struct {
	Id              int                    `gorm:"primaryKey;auto_increment;type:serial"`
	Uuid            string                 `gorm:"type:character varying"`
	CustomerId      int                    `gorm:"type:integer"`
	Status          string                 `gorm:"type:character varying"`
	ShippingAddress string                 `gorm:"type:character varying"`
	IsDeleted       int                    `gorm:"type:integer"`
	Username        string                 `gorm:"-:migration;<-:false"`
	Email           string                 `gorm:"-:migration;<-:false"`
	MobileNo        string                 `gorm:"-:migration;<-:false"`
	StreetAddress   string                 `gorm:"-:migration;<-:false"`
	City            string                 `gorm:"-:migration;<-:false"`
	State           string                 `gorm:"-:migration;<-:false"`
	Country         string                 `gorm:"-:migration;<-:false"`
	ZipCode         string                 `gorm:"-:migration;<-:false"`
	CreatedOn       time.Time              `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedOn      time.Time              `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	ModifiedDate    string                 `gorm:"-:migration"`
	CreatedDate     string                 `gorm:"-:migration"`
	Price           int                    `gorm:"type:integer"`
	Tax             int                    `gorm:"type:integer"`
	TotalCost       int                    `gorm:"type:integer"`
	FirstName       string                 `gorm:"-:migration;<-:false"`
	LastName        string                 `gorm:"-:migration;<-:false"`
	NameString      string                 `gorm:"-:migration;<-:false"`
	Orders          []TblEcomOrderStatuses `gorm:"foreignKey:OrderId;references:Id"`
	DeletedOn       time.Time              `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	DeletedBy       int                    `gorm:"type:integer;DEFAULT:NULL"`
	CreatedBy       int                    `gorm:"type:integer;DEFAULT:NULL"`
	ModifiedBy      int                    `gorm:"type:integer;DEFAULT:NULL"`
}

type tblecomproductorderdetails struct {
	Id         int `gorm:"primaryKey;auto_increment;type:serial"`
	Order_id   int `gorm:"type:integer"`
	Product_id int `gorm:"type:integer"`
	Quantity   int `gorm:"type:integer"`
	Price      int `gorm:"type:integer"`
}

type tblEcomOrderStatuses struct {
	Id          int       `gorm:"primaryKey;auto_increment;type:serial"`
	OrderId     int       `gorm:"type:integer"`
	OrderStatus string    `gorm:"type:character varying"`
	CreatedBy   int       `gorm:"type:integer"`
	CreatedOn   time.Time `gorm:"type:timestamp without time zone;DEFAULT:NULL"`
	CreatedDate string    `gorm:"-:migration;<-:false"`
}

// pass limit , offset get orderslist
func (ecommerce *Ecommerce) OrdersList(offset int, limit int, filter Filter) (order []TblEcomProductOrders, count int64, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return []TblEcomProductOrders{}, 0, AuthErr
	}
	orders, _, err := Ecommercemodel.OrderList(offset, limit, filter, ecommerce.DB)

	if err != nil {
		log.Println(err)
	}

	_, totalcount, _ := Ecommercemodel.OrderList(0, 0, filter, ecommerce.DB)

	var finalorderlist []TblEcomProductOrders

	for _, orders := range orders {

		orders.CreatedDate = orders.CreatedOn.In(TZONE).Format("02 Jan 2006 03:04 PM")

		if !orders.ModifiedOn.IsZero() {

			orders.ModifiedDate = orders.ModifiedOn.In(TZONE).Format("02 Jan 2006 03:04 PM")

		} else {

			orders.ModifiedDate = orders.CreatedOn.In(TZONE).Format("02 Jan 2006 03:04 PM")

		}

		finalorderlist = append(finalorderlist, orders)
	}

	return finalorderlist, totalcount, nil

}

// pass Order id  get particular Order details

func (ecommerce *Ecommerce) OrderInfo(id string) (orderlists TblEcomProductOrders, product []tblEcomProducts, Address OrderShippingAddress, count int, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return TblEcomProductOrders{}, []tblEcomProducts{}, OrderShippingAddress{}, 0, AuthErr
	}

	orderlist, err := Ecommercemodel.OrderEdit(id, ecommerce.DB)

	if err != nil {
		log.Println(err)
	}
	var shippingAddress OrderShippingAddress

	err4 := json.Unmarshal([]byte(orderlist.ShippingAddress), &shippingAddress)

	if err4 != nil {
		fmt.Println("Error unmarshalling JSON:", err4)
		return
	}
	orderlist.CreatedDate = orderlist.CreatedOn.In(TZONE).Format("02 Jan 2006 03:04 PM")

	orderlist.ModifiedDate = orderlist.ModifiedOn.In(TZONE).Format("02 Jan 2006 03:04 PM")

	orid := orderlist.Id

	var length int

	productinfo, err1 := Ecommercemodel.GetProductdetailsByOrderId(orid, ecommerce.DB)

	length = len(productinfo)

	if err1 != nil {

		log.Println(err1)
	}

	var product_id []int

	for _, val := range productinfo {

		product_id = append(product_id, val.Product_id)

	}

	productdetails, err2 := Ecommercemodel.GetProductdetailsByProductId(product_id, ecommerce.DB)

	if err2 != nil {

		log.Println(err2)
	}

	var productList []tblEcomProducts

	for i, val := range productdetails {

		imgs := strings.Split(val.ProductImagePath, ",")

		if len(imgs) > 0 {

			val.ProductImagePath = imgs[0]

		}
		if i < len(productinfo) {
			fmt.Println("productinfo[i].Quantity", productinfo[i].Quantity)
			quantity := productinfo[i].Quantity
			price := productinfo[i].Price
			quantityPrice := quantity * price

			fmt.Println("quantityPrice", quantityPrice, price)
			productList = append(productList, tblEcomProducts{

				ProductImagePath:   val.ProductImagePath,
				ProductDescription: val.ProductDescription,
				ProductName:        val.ProductName,
				Quantity:           productinfo[i].Quantity,
				Price:              productinfo[i].Price,
				Quantityprice:      quantityPrice,
			})
		} else {

			productList = append(productList, tblEcomProducts{
				ProductImagePath:   val.ProductImagePath,
				ProductDescription: val.ProductDescription,
				ProductName:        val.ProductName,
				Quantity:           0,
				Price:              0,
				Quantityprice:      0,
			})
		}
	}

	return orderlist, productList, shippingAddress, length, nil

}

// Update Order status
func (ecommerce *Ecommerce) UpdateOrderStatus(orderid, status int) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	var order TblEcomProductOrders

	order.Id = orderid

	order.OrderStatus = status

	order.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Ecommercemodel.OrderStatusUpdate(order, ecommerce.DB)

	if err != nil {
		return err
	}
	var orderstatus TblEcomOrderStatuses

	orderstatus.OrderId = orderid
	orderstatus.OrderStatus = status
	orderstatus.CreatedBy = 1
	orderstatus.CreatedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err1 := Ecommercemodel.CreateOrderStatus(orderstatus, ecommerce.DB)

	if err1 != nil {
		return err
	}

	return nil
}

// pass Order id soft delete the particular record

func (ecommerce *Ecommerce) DeleteOrder(id int, deletedby int) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	var order TblEcomProductOrders

	order.Id = id

	order.DeletedBy = deletedby

	order.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Ecommercemodel.OrderDelete(order, ecommerce.DB)

	if err != nil {
		return err
	}

	return nil

}

// multi delete order
func (ecommerce *Ecommerce) MultiSelectOrdersDelete(orderids []int, deletedby int) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	orderidint := make([]int, len(orderids))

	for i, id := range orderids {

		intId := id

		orderidint[i] = intId

	}

	var order TblEcomProductOrders

	order.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	order.DeletedBy = deletedby

	err1 := Ecommercemodel.MultiSelectDeleteOrder(order, orderidint, ecommerce.DB)

	if err1 != nil {

		return err1
	}

	return nil

}

// Get Order status list
func (Ecommerce *Ecommerce) GetOrderStatusList() (orderStatus []OrderStatusNames, err error) {

	if AuthErr := AuthandPermission(Ecommerce); AuthErr != nil {

		return []OrderStatusNames{}, AuthErr
	}

	orderStatus, err = EcommerceModel.GetOrderStatusList(EcommerceModel{}, Ecommerce.DB)
	if err != nil {

		return []OrderStatusNames{}, err
	}

	return orderStatus, nil
}

func (Ecommerce *Ecommerce) PlaceOrder(orderPlaced EcommerceOrder) (err error) {

	if AuthErr := AuthandPermission(Ecommerce); AuthErr != nil {

		return AuthErr
	}

	err = EcommerceModel.PlaceOrder(EcommerceModel{}, orderPlaced, Ecommerce.DB)
	if err != nil {

		return err
	}

	return nil
}

func (Ecommerce *Ecommerce) GetOrderByOrderId(orderId string) (order EcommerceOrder, err error) {

	if AuthErr := AuthandPermission(Ecommerce); AuthErr != nil {

		return EcommerceOrder{}, AuthErr
	}

	order, err = EcommerceModel.GetOrderByOrderId(EcommerceModel{}, orderId, Ecommerce.DB)
	if err != nil {

		return EcommerceOrder{}, err
	}

	return order, err
}

func (Ecommerce *Ecommerce) CreateOrderDetails(orderDetails OrderProduct) (err error) {

	if AuthErr := AuthandPermission(Ecommerce); AuthErr != nil {

		return AuthErr
	}

	err = EcommerceModel.CreateOrderDetails(EcommerceModel{}, orderDetails, Ecommerce.DB)
	if err != nil {

		return err
	}

	return nil
}

func (Ecommerce *Ecommerce) UpdateStock(productId, quantity int) (err error) {

	if AuthErr := AuthandPermission(Ecommerce); AuthErr != nil {

		return AuthErr
	}

	err = EcommerceModel.UpdateStock(EcommerceModel{}, productId, quantity, Ecommerce.DB)
	if err != nil {

		return err
	}

	return nil
}

func (Ecommerce *Ecommerce) CreateOrderStatus(orderStatus TblEcomOrderStatuses) (err error) {

	if AuthErr := AuthandPermission(Ecommerce); AuthErr != nil {

		return AuthErr
	}

	err = EcommerceModel.CreateOrderStatus(Ecommercemodel, orderStatus, Ecommerce.DB)
	if err != nil {

		return err
	}

	return nil
}

func (Ecommerce *Ecommerce) CreateOrderPayment(orderPayment OrderPayment) (err error) {

	if AuthErr := AuthandPermission(Ecommerce); AuthErr != nil {

		return AuthErr
	}

	err = EcommerceModel.CreateOrderPayment(EcommerceModel{}, orderPayment, Ecommerce.DB)
	if err != nil {

		return err
	}

	return nil
}

func (Ecommerce *Ecommerce) DeleteFromCartAfterOrder(orderedproductIds []int, customerId int) (err error) {

	if AuthErr := AuthandPermission(Ecommerce); AuthErr != nil {

		return AuthErr
	}

	err = EcommerceModel.DeleteFromCartAfterOrder(EcommerceModel{}, orderedproductIds, customerId, Ecommerce.DB)
	if err != nil {

		return err
	}

	return nil
}
