package ecommerce

import (
	"encoding/json"
	"fmt"
	"log"

	"strings"
	"time"
)

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

	return orders, totalcount, nil

}

// pass Order id  get particular Order details

func (ecommerce *Ecommerce) OrderInfo(id string) (orderlists TblEcomProductOrders, product []tblEcomProducts, Address OrderShippingAddress, count int, status []TblEcomOrderStatus, err error) {

	if AuthErr := AuthandPermission(ecommerce); AuthErr != nil {

		return TblEcomProductOrders{}, []tblEcomProducts{}, OrderShippingAddress{}, 0, []TblEcomOrderStatus{}, AuthErr
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

	// To get order stauts is particular id
	statusdetails, err6 := Ecommercemodel.OrderStatusDetails(id, ecommerce.DB)

	if err6 != nil {
		log.Println(err6)
	}

	var productList []tblEcomProducts

	for i, val := range productdetails {

		imgs := strings.Split(val.ProductImagePath, ",")

		if len(imgs) > 0 {

			val.ProductImagePath = imgs[0]

		}
		if i < len(productinfo) {
			
			quantity := productinfo[i].Quantity
			price := productinfo[i].Price
			quantityPrice := quantity * price

			
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

	return orderlist, productList, shippingAddress, length, statusdetails, nil

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

func (Ecommerce *Ecommerce) CreateOrderStatus(orderStatus TblEcomOrderStatus) (err error) {

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
