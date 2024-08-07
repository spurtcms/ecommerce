package ecommerce

import (
	"encoding/json"
	"fmt"
	"log"

	"strings"
	"time"
)

type OrderShippingAddress struct {
	FirstName     string `json:"firstname"`
	UserName      string `json:"username"`
	StreetAddress string `json:"streetaddress"`
	MobileNo      string `json:"mobileno"`
	Email         string `json:"email"`
	ZipCode       string `json:"zipcode"`
	City          string `json:"city"`
	Country       string `json:"country"`
	State         string `json:"state"`
	ProfileImage  string `json:"profileimage"`
	IsActive      int    `json:"isactive"`
}

// pass limit , offset get orderslist
func (ecommerce *Ecommerce) OrdersList(offset int, limit int, filter Filter, tenantid int) (order []TblEcomProductOrder, count int64, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return []TblEcomProductOrder{}, 0, AuthErr
	}
	orders, _, err := Ecommercemodel.OrderList(offset, limit, filter, ecommerce.DB, tenantid)

	if err != nil {
		log.Println(err)
	}

	_, totalcount, _ := Ecommercemodel.OrderList(0, 0, filter, ecommerce.DB, tenantid)

	return orders, totalcount, nil

}

// pass Order id  get particular Order details

func (ecommerce *Ecommerce) OrderInfo(id string, tenantid int) (orderlists TblEcomProductOrder, product []TblEcomProduct, Address OrderShippingAddress, count int, status []TblEcomOrderStatus, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return TblEcomProductOrder{}, []TblEcomProduct{}, OrderShippingAddress{}, 0, []TblEcomOrderStatus{}, AuthErr
	}

	orderlist, err := Ecommercemodel.OrderEdit(id, ecommerce.DB, tenantid)

	if err != nil {
		log.Println(err)
	}
	var shippingAddress OrderShippingAddress

	err4 := json.Unmarshal([]byte(orderlist.ShippingAddress), &shippingAddress)

	if err4 != nil {
		fmt.Println("Error unmarshalling JSON:", err4)
		return
	}

	orid := orderlist.Id

	var length int

	productinfo, err1 := Ecommercemodel.GetProductdetailsByOrderId(orid, ecommerce.DB, tenantid)

	length = len(productinfo)

	if err1 != nil {

		log.Println(err1)
	}

	var product_id []int

	for _, val := range productinfo {

		product_id = append(product_id, val.Product_id)

	}

	productdetails, err2 := Ecommercemodel.GetProductdetailsByProductId(product_id, ecommerce.DB, tenantid)

	if err2 != nil {

		log.Println(err2)
	}

	// To get order stauts is particular id
	statusdetails, err6 := Ecommercemodel.OrderStatusDetails(id, ecommerce.DB, tenantid)

	if err6 != nil {
		log.Println(err6)
	}

	var productList []TblEcomProduct

	for i, val := range productdetails {

		imgs := strings.Split(val.ProductImagePath, ",")

		if len(imgs) > 0 {

			val.ProductImagePath = imgs[0]

		}
		if i < len(productinfo) {

			quantity := productinfo[i].Quantity
			price := productinfo[i].Price
			quantityPrice := quantity * price

			productList = append(productList, TblEcomProduct{

				ProductImagePath:   val.ProductImagePath,
				ProductDescription: val.ProductDescription,
				ProductName:        val.ProductName,
				Quantity:           productinfo[i].Quantity,
				Price:              productinfo[i].Price,
				Quantityprice:      quantityPrice,
			})
		} else {

			productList = append(productList, TblEcomProduct{
				ProductImagePath:   val.ProductImagePath,
				ProductDescription: val.ProductDescription,
				ProductName:        val.ProductName,
				Quantity:           0,
				Price:              0,
				Quantityprice:      0,
			})
		}
	}

	return orderlist, productList, shippingAddress, length, statusdetails, nil

}

// Update Order status
func (ecommerce *Ecommerce) UpdateOrderStatus(orderid, status int, tenantid int) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	var order TblEcomProductOrder

	order.Id = orderid

	order.OrderStatus = status

	order.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Ecommercemodel.OrderStatusUpdate(order, ecommerce.DB, tenantid)

	if err != nil {
		return err
	}
	var orderstatus TblEcomOrderStatus

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

func (ecommerce *Ecommerce) DeleteOrder(id int, deletedby int, tenantid int) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	var order TblEcomProductOrder

	order.Id = id

	order.DeletedBy = deletedby

	order.DeletedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	err := Ecommercemodel.OrderDelete(order, ecommerce.DB, tenantid)

	if err != nil {
		return err
	}

	return nil

}

// multi delete order
func (ecommerce *Ecommerce) MultiSelectOrdersDelete(orderids []int, deletedby int, tenantid int) error {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return AuthErr
	}

	orderidint := make([]int, len(orderids))

	for i, id := range orderids {

		intId := id

		orderidint[i] = intId

	}

	var order TblEcomProductOrder

	order.ModifiedOn, _ = time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

	order.DeletedBy = deletedby

	err1 := Ecommercemodel.MultiSelectDeleteOrder(order, orderidint, ecommerce.DB, tenantid)

	if err1 != nil {

		return err1
	}

	return nil

}

// Get Order status list
func (Ecommerce *Ecommerce) GetOrderStatusList(tenantid int) (orderStatus []OrderStatusNames, err error) {

	if AuthErr := AuthandPermission(Ecommerce); AuthErr != nil {

		return []OrderStatusNames{}, AuthErr
	}

	orderStatus, err = Ecommercemodel.GetOrderStatusList(Ecommerce.DB, tenantid)
	if err != nil {

		return []OrderStatusNames{}, err
	}

	return orderStatus, nil
}

func (Ecommerce *Ecommerce) PlaceOrder(orderPlaced EcommerceOrder) (err error) {

	if AuthErr := AuthandPermission(Ecommerce); AuthErr != nil {

		return AuthErr
	}

	err = Ecommercemodel.PlaceOrder(orderPlaced, Ecommerce.DB)
	if err != nil {

		return err
	}

	return nil
}

func (Ecommerce *Ecommerce) GetOrderByOrderId(orderId string, tenantid int) (order EcommerceOrder, err error) {

	if AuthErr := AuthandPermission(Ecommerce); AuthErr != nil {

		return EcommerceOrder{}, AuthErr
	}

	order, err = Ecommercemodel.GetOrderByOrderId(orderId, Ecommerce.DB, tenantid)
	if err != nil {

		return EcommerceOrder{}, err
	}

	return order, err
}

func (Ecommerce *Ecommerce) CreateOrderDetails(orderDetails OrderProduct) (err error) {

	if AuthErr := AuthandPermission(Ecommerce); AuthErr != nil {

		return AuthErr
	}

	err = Ecommercemodel.CreateOrderDetails(orderDetails, Ecommerce.DB)
	if err != nil {

		return err
	}

	return nil
}

func (Ecommerce *Ecommerce) UpdateStock(productId, quantity int, tenantid int) (err error) {

	if AuthErr := AuthandPermission(Ecommerce); AuthErr != nil {

		return AuthErr
	}

	err = Ecommercemodel.UpdateStock(productId, quantity, Ecommerce.DB, tenantid)
	if err != nil {

		return err
	}

	return nil
}

func (Ecommerce *Ecommerce) CreateOrderStatus(orderStatus TblEcomOrderStatus) (err error) {

	if AuthErr := AuthandPermission(Ecommerce); AuthErr != nil {

		return AuthErr
	}

	err = Ecommercemodel.CreateOrderStatus(orderStatus, Ecommerce.DB)
	if err != nil {

		return err
	}

	return nil
}

func (Ecommerce *Ecommerce) CreateOrderPayment(orderPayment OrderPayment) (err error) {

	if AuthErr := AuthandPermission(Ecommerce); AuthErr != nil {

		return AuthErr
	}

	err = Ecommercemodel.CreateOrderPayment(orderPayment, Ecommerce.DB)
	if err != nil {

		return err
	}

	return nil
}

func (Ecommerce *Ecommerce) DeleteFromCartAfterOrder(orderedproductIds []int, customerId int, tenantid int) (err error) {

	if AuthErr := AuthandPermission(Ecommerce); AuthErr != nil {

		return AuthErr
	}

	err = Ecommercemodel.DeleteFromCartAfterOrder(orderedproductIds, customerId, Ecommerce.DB, tenantid)
	if err != nil {

		return err
	}

	return nil
}
