package ecommerce

import (
	"fmt"
	"log"
	"testing"

	"github.com/spurtcms/auth"
)

// To Test Order list
func TestOrdersList(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     1,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Orders", auth.CRUD)

	ecommerce := EcommerceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
		Auth:             Auth,
	})

	if permisison {

		order, count, err := ecommerce.OrdersList(10, 0, Filter{})

		if err != nil {

			panic(err)
		}

		fmt.Println(order, count)
	} else {

		log.Println("permissions enabled not initialised")

	}

}

// To Test

func TestOrderInfo(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     1,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Orders", auth.CRUD)

	ecommerce := EcommerceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
		Auth:             Auth,
	})

	if permisison {

		id := "SP03052024"

		orderstatus, product, address, count, status, err := ecommerce.OrderInfo(id)

		if err != nil {

			panic(err)
		}

		fmt.Println(orderstatus, product, address, count, status)
	} else {

		log.Println("permissions enabled not initialised")

	}

}

func TestUpdateOrderStatus(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     1,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Orders", auth.CRUD)

	ecommerce := EcommerceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
		Auth:             Auth,
	})

	if permisison {

		err := ecommerce.UpdateOrderStatus(1, 1)

		if err != nil {

			panic(err)
		}

	} else {

		log.Println("permissions enabled not initialised")

	}

}

func TestDeleteOrder(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     1,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Orders", auth.CRUD)

	ecommerce := EcommerceSetup(Config{
		DB:               db,
		AuthEnable:       false,
		PermissionEnable: false,
		Auth:             Auth,
	})

	if permisison {

		err := ecommerce.DeleteOrder(1, 1)

		if err != nil {

			panic(err)
		}

	} else {

		log.Println("permissions enabled not initialised")

	}

}

func TestMultiSelectOrdersDelete(t *testing.T) {

	db, _ := DBSetup()

	config := auth.Config{
		UserId:     1,
		ExpiryTime: 2,
		ExpiryFlg:  true,
		SecretKey:  "Secret123",
		DB:         db,
		RoleId:     1,
	}

	Auth := auth.AuthSetup(config)

	token, _ := Auth.CreateToken()

	Auth.VerifyToken(token, SecretKey)

	permisison, _ := Auth.IsGranted("Orders", auth.CRUD)

	ecommerce := EcommerceSetup(Config{
		DB:               db,
		AuthEnable:       true,
		PermissionEnable: true,
		Auth:             Auth,
	})

	var orderid []int

	if permisison {

		err := ecommerce.MultiSelectOrdersDelete(orderid, 1)

		if err != nil {

			panic(err)
		}

	} else {

		log.Println("permissions enabled not initialised")

	}

}
